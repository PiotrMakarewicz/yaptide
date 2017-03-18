package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Palantir/palantir/model/auth"
	"github.com/Palantir/palantir/web/auth/token"
	"github.com/Palantir/palantir/web/server"
)

type loginHandler struct {
	*server.Context
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	acc := auth.Account{}
	err := json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := h.Db.Copy()
	defer db.Close()

	responseMap := make(map[string]string)

	res, err := db.Account.GetIfAuthenticated(&acc)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tokenString, err := token.Generate(res, h.JWTKey)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseMap["token"] = tokenString

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseMap)
	if err != nil {
		log.Println(err.Error())
	}
}
