package material

import (
	"github.com/Palantir/palantir/model/color"
)

// PredefinedMaterialRecord contains data needed on frontend related to PredefinedMaterials.
type PredefinedMaterialRecord struct {
	Value string      `json:"value"`
	Name  string      `json:"name"`
	Color color.Color `json:"color"`
}

// IsotopeRecord contains data needed on frontend related to Isotopes.
type IsotopeRecord struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

var waterColor = color.New(0x00, 0x93, 0xDD, 0xFF)

var predefinedMaterials = []PredefinedMaterialRecord{
	PredefinedMaterialRecord{"hydrogen", "Hydrogen (Z: 1)", color.White},
	PredefinedMaterialRecord{"helium", "Helium (Z: 2)", color.Gray},
	PredefinedMaterialRecord{"lithium", "Lithium (Z: 3)", color.Gray},
	PredefinedMaterialRecord{"beryllium", "Beryllium (Z: 4)", color.Gray},
	PredefinedMaterialRecord{"boron", "Boron (Z: 5)", color.Gray},
	PredefinedMaterialRecord{"carbon", "Carbon (Z: 6), Amorphous (density 2.0 g/cm3)", color.Gray},
	PredefinedMaterialRecord{"graphite", "Graphite (Z: 6) (density 1.7 g/cm3)", color.Gray},
	PredefinedMaterialRecord{"nitrogen", "Nitrogen (Z: 7)", color.Gray},
	PredefinedMaterialRecord{"oxygen", "Oxygen (Z: 8)", color.Gray},
	PredefinedMaterialRecord{"fluorine", "Fluorine (Z: 9)", color.Gray},
	PredefinedMaterialRecord{"neon", "Neon (Z: 10)", color.Gray},
	PredefinedMaterialRecord{"sodium", "Sodium (Z: 11)", color.Gray},
	PredefinedMaterialRecord{"magnesium", "Magnesium (Z: 12)", color.Gray},
	PredefinedMaterialRecord{"aluminum", "Aluminum (Z: 13)", color.Gray},
	PredefinedMaterialRecord{"silicon", "Silicon (Z: 14)", color.Gray},
	PredefinedMaterialRecord{"phosphorus", "Phosphorus (Z: 15)", color.Gray},
	PredefinedMaterialRecord{"sulfur", "Sulfur (Z: 16)", color.Gray},
	PredefinedMaterialRecord{"chlorine", "Chlorine (Z: 17)", color.Gray},
	PredefinedMaterialRecord{"argon", "Argon (Z: 18)", color.Gray},
	PredefinedMaterialRecord{"potassium", "Potassium (Z: 19)", color.Gray},
	PredefinedMaterialRecord{"calcium", "Calcium (Z: 20)", color.Gray},
	PredefinedMaterialRecord{"scandium", "Scandium (Z: 21)", color.Gray},
	PredefinedMaterialRecord{"titanium", "Titanium (Z: 22)", color.Gray},
	PredefinedMaterialRecord{"vanadium", "Vanadium (Z: 23)", color.Gray},
	PredefinedMaterialRecord{"chromium", "Chromium (Z: 24)", color.Gray},
	PredefinedMaterialRecord{"manganese", "Manganese (Z: 25)", color.Gray},
	PredefinedMaterialRecord{"iron", "Iron (Z: 26)", color.Gray},
	PredefinedMaterialRecord{"cobalt", "Cobalt (Z: 27)", color.Gray},
	PredefinedMaterialRecord{"nickel", "Nickel (Z: 28)", color.Gray},
	PredefinedMaterialRecord{"copper", "Copper (Z: 29)", color.Gray},
	PredefinedMaterialRecord{"zinc", "Zinc (Z: 30)", color.Gray},
	PredefinedMaterialRecord{"gallium", "Gallium (Z: 31)", color.Gray},
	PredefinedMaterialRecord{"germanium", "Germanium (Z: 32)", color.Gray},
	PredefinedMaterialRecord{"arsenic", "Arsenic (Z: 33)", color.Gray},
	PredefinedMaterialRecord{"selenium", "Selenium (Z: 34)", color.Gray},
	PredefinedMaterialRecord{"bromine", "Bromine (Z: 35)", color.Gray},
	PredefinedMaterialRecord{"krypton", "Krypton (Z: 36)", color.Gray},
	PredefinedMaterialRecord{"rubidium", "Rubidium (Z: 37)", color.Gray},
	PredefinedMaterialRecord{"strontium", "Strontium (Z: 38)", color.Gray},
	PredefinedMaterialRecord{"yttrium", "Yttrium (Z: 39)", color.Gray},
	PredefinedMaterialRecord{"zirconium", "Zirconium (Z: 40)", color.Gray},
	PredefinedMaterialRecord{"niobium", "Niobium (Z: 41)", color.Gray},
	PredefinedMaterialRecord{"molybdenum", "Molybdenum (Z: 42)", color.Gray},
	PredefinedMaterialRecord{"technetium", "Technetium (Z: 43)", color.Gray},
	PredefinedMaterialRecord{"ruthenium", "Ruthenium (Z: 44)", color.Gray},
	PredefinedMaterialRecord{"rhodium", "Rhodium (Z: 45)", color.Gray},
	PredefinedMaterialRecord{"palladium", "Palladium (Z: 46)", color.Gray},
	PredefinedMaterialRecord{"silver", "Silver (Z: 47)", color.Gray},
	PredefinedMaterialRecord{"cadmium", "Cadmium (Z: 48)", color.Gray},
	PredefinedMaterialRecord{"indium", "Indium (Z: 49)", color.Gray},
	PredefinedMaterialRecord{"tin", "Tin (Z: 50)", color.Gray},
	PredefinedMaterialRecord{"antimony", "Antimony (Z: 51)", color.Gray},
	PredefinedMaterialRecord{"tellurium", "Tellurium (Z: 52)", color.Gray},
	PredefinedMaterialRecord{"iodine", "Iodine (Z: 53)", color.Gray},
	PredefinedMaterialRecord{"xenon", "Xenon (Z: 54)", color.Gray},
	PredefinedMaterialRecord{"cesium", "Cesium (Z: 55)", color.Gray},
	PredefinedMaterialRecord{"barium", "Barium (Z: 56)", color.Gray},
	PredefinedMaterialRecord{"lanthanum", "Lanthanum (Z: 57)", color.Gray},
	PredefinedMaterialRecord{"cerium", "Cerium (Z: 58)", color.Gray},
	PredefinedMaterialRecord{"praseodymium", "Praseodymium (Z: 59)", color.Gray},
	PredefinedMaterialRecord{"neodymium", "Neodymium (Z: 60)", color.Gray},
	PredefinedMaterialRecord{"promethium", "Promethium (Z: 61)", color.Gray},
	PredefinedMaterialRecord{"samarium", "Samarium (Z: 62)", color.Gray},
	PredefinedMaterialRecord{"europium", "Europium (Z: 63)", color.Gray},
	PredefinedMaterialRecord{"gadolinium", "Gadolinium (Z: 64)", color.Gray},
	PredefinedMaterialRecord{"terbium", "Terbium (Z: 65)", color.Gray},
	PredefinedMaterialRecord{"dysprosium", "Dysprosium (Z: 66)", color.Gray},
	PredefinedMaterialRecord{"holmium", "Holmium (Z: 67)", color.Gray},
	PredefinedMaterialRecord{"erbium", "Erbium (Z: 68)", color.Gray},
	PredefinedMaterialRecord{"thulium", "Thulium (Z: 69)", color.Gray},
	PredefinedMaterialRecord{"ytterbium", "Ytterbium (Z: 70)", color.Gray},
	PredefinedMaterialRecord{"lutetium", "Lutetium (Z: 71)", color.Gray},
	PredefinedMaterialRecord{"hafnium", "Hafnium (Z: 72)", color.Gray},
	PredefinedMaterialRecord{"tantalum", "Tantalum (Z: 73)", color.Gray},
	PredefinedMaterialRecord{"tungsten", "Tungsten (Z: 74)", color.Gray},
	PredefinedMaterialRecord{"rhenium", "Rhenium (Z: 75)", color.Gray},
	PredefinedMaterialRecord{"osmium", "Osmium (Z: 76)", color.Gray},
	PredefinedMaterialRecord{"iridium", "Iridium (Z: 77)", color.Gray},
	PredefinedMaterialRecord{"platinum", "Platinum (Z: 78)", color.Gray},
	PredefinedMaterialRecord{"gold", "Gold (Z: 79)", color.Gray},
	PredefinedMaterialRecord{"mercury", "Mercury (Z: 80)", color.Gray},
	PredefinedMaterialRecord{"thallium", "Thallium (Z: 81)", color.Gray},
	PredefinedMaterialRecord{"lead", "Lead (Z: 82)", color.Gray},
	PredefinedMaterialRecord{"bismuth", "Bismuth (Z: 83)", color.Gray},
	PredefinedMaterialRecord{"polonium", "Polonium (Z: 84)", color.Gray},
	PredefinedMaterialRecord{"astatine", "Astatine (Z: 85)", color.Gray},
	PredefinedMaterialRecord{"radon", "Radon (Z: 86)", color.Gray},
	PredefinedMaterialRecord{"francium", "Francium (Z: 87)", color.Gray},
	PredefinedMaterialRecord{"radium", "Radium (Z: 88)", color.Gray},
	PredefinedMaterialRecord{"actinium", "Actinium (Z: 89)", color.Gray},
	PredefinedMaterialRecord{"thorium", "Thorium (Z: 90)", color.Gray},
	PredefinedMaterialRecord{"protactinium", "Protactinium (Z: 91)", color.Gray},
	PredefinedMaterialRecord{"uranium", "Uranium (Z: 92)", color.Gray},
	PredefinedMaterialRecord{"neptunium", "Neptunium (Z: 93)", color.Gray},
	PredefinedMaterialRecord{"plutonium", "Plutonium (Z: 94)", color.Gray},
	PredefinedMaterialRecord{"americium", "Americium (Z: 95)", color.Gray},
	PredefinedMaterialRecord{"curium", "Curium (Z: 96)", color.Gray},
	PredefinedMaterialRecord{"berkelium", "Berkelium (Z: 97)", color.Gray},
	PredefinedMaterialRecord{"californium", "Californium (Z: 98)", color.Gray},
	PredefinedMaterialRecord{"a_150", "A-150 Tissue-Equivalent Plastic", color.Gray},
	PredefinedMaterialRecord{"acetone", "Acetone", color.Gray},
	PredefinedMaterialRecord{"acetylene", "Acetylene", color.Gray},
	PredefinedMaterialRecord{"adenine", "Adenine", color.Gray},
	PredefinedMaterialRecord{"adipose", "Adipose Tissue (ICRP)", color.Gray},
	PredefinedMaterialRecord{"air_dry", "Air, Dry (near sea level)", color.Gray},
	PredefinedMaterialRecord{"alanine", "Alanine", color.Gray},
	PredefinedMaterialRecord{"aluminum_oxide", "Aluminum Oxide", color.Gray},
	PredefinedMaterialRecord{"amber", "Amber", color.Gray},
	PredefinedMaterialRecord{"ammonia", "Ammonia", color.Gray},
	PredefinedMaterialRecord{"aniline", "Aniline", color.Gray},
	PredefinedMaterialRecord{"anthracene", "Anthracene", color.Gray},
	PredefinedMaterialRecord{"b_100", "B-100 Bone-Equivalent Plastic", color.Gray},
	PredefinedMaterialRecord{"bakelite", "Bakelite", color.Gray},
	PredefinedMaterialRecord{"barium_fluoride", "Barium Fluoride", color.Gray},
	PredefinedMaterialRecord{"barium_sulfate", "Barium Sulfate", color.Gray},
	PredefinedMaterialRecord{"benzene", "Benzene", color.Gray},
	PredefinedMaterialRecord{"beryllium_oxide", "Beryllium oxide", color.Gray},
	PredefinedMaterialRecord{"bismuth_germanium_oxide", "Bismuth Germanium oxide", color.Gray},
	PredefinedMaterialRecord{"blood", "Blood (ICRP)", color.Gray},
	PredefinedMaterialRecord{"bone_compact", "Bone, Compact (ICRU)", color.Gray},
	PredefinedMaterialRecord{"bone_cortical", "Bone, Cortical (ICRP)", color.Gray},
	PredefinedMaterialRecord{"boron_carbide", "Boron Carbide", color.Gray},
	PredefinedMaterialRecord{"boron_oxide", "Boron Oxide", color.Gray},
	PredefinedMaterialRecord{"brain", "Brain (ICRP)", color.Gray},
	PredefinedMaterialRecord{"butane", "Butane", color.Gray},
	PredefinedMaterialRecord{"n_butyl_alcohol", "N-Butyl Alcohol", color.Gray},
	PredefinedMaterialRecord{"c_552", "C-552 Air-Equivalent Plastic", color.Gray},
	PredefinedMaterialRecord{"cadmium_telluride", "Cadmium Telluride", color.Gray},
	PredefinedMaterialRecord{"cadmium_tungstate", "Cadmium Tungstate", color.Gray},
	PredefinedMaterialRecord{"calcium_carbonate", "Calcium Carbonate", color.Gray},
	PredefinedMaterialRecord{"calcium_fluoride", "Calcium Fluoride", color.Gray},
	PredefinedMaterialRecord{"calcium_oxide", "Calcium Oxide", color.Gray},
	PredefinedMaterialRecord{"calcium_sulfate", "Calcium Sulfate", color.Gray},
	PredefinedMaterialRecord{"calcium_tungstate", "Calcium Tungstate", color.Gray},
	PredefinedMaterialRecord{"carbon_dioxide", "Carbon Dioxide", color.Gray},
	PredefinedMaterialRecord{"carbon_tetrachloride", "Carbon Tetrachloride", color.Gray},
	PredefinedMaterialRecord{"cellulose_acetate_cellophane", "Cellulose Acetate, Cellophane", color.Gray},
	PredefinedMaterialRecord{"cellulose_acetate_butyrate", "Cellulose Acetate Butyrate", color.Gray},
	PredefinedMaterialRecord{"cellulose_nitrate", "Cellulose Nitrate", color.Gray},
	PredefinedMaterialRecord{"ceric_sulfate_dosimeter_solution", "Ceric Sulfate Dosimeter Solution", color.Gray},
	PredefinedMaterialRecord{"cesium_fluoride", "Cesium Fluoride", color.Gray},
	PredefinedMaterialRecord{"cesium_iodide", "Cesium Iodide", color.Gray},
	PredefinedMaterialRecord{"chlorobenzene", "Chlorobenzene", color.Gray},
	PredefinedMaterialRecord{"chloroform", "Chloroform", color.Gray},
	PredefinedMaterialRecord{"concrete_portland", "Concrete, Portland", color.Gray},
	PredefinedMaterialRecord{"cyclohexane", "Cyclohexane", color.Gray},
	PredefinedMaterialRecord{"1_2_ddihlorobenzene", "1,2-Ddihlorobenzene", color.Gray},
	PredefinedMaterialRecord{"dichlorodiethyl_ether", "Dichlorodiethyl Ether", color.Gray},
	PredefinedMaterialRecord{"1_2_dichloroethane", "1,2-Dichloroethane", color.Gray},
	PredefinedMaterialRecord{"diethyl_ether", "Diethyl Ether", color.Gray},
	PredefinedMaterialRecord{"n_n_dimethyl_formamide", "N,N-Dimethyl Formamide", color.Gray},
	PredefinedMaterialRecord{"dimethyl_sulfoxide", "Dimethyl Sulfoxide", color.Gray},
	PredefinedMaterialRecord{"ethane", "Ethane", color.Gray},
	PredefinedMaterialRecord{"ethyl_alcohol", "Ethyl Alcohol", color.Gray},
	PredefinedMaterialRecord{"ethyl_cellulose", "Ethyl Cellulose", color.Gray},
	PredefinedMaterialRecord{"ethylene", "Ethylene", color.Gray},
	PredefinedMaterialRecord{"eye_lens", "Eye Lens (ICRP)", color.Gray},
	PredefinedMaterialRecord{"ferric_oxide", "Ferric Oxide", color.Gray},
	PredefinedMaterialRecord{"ferroboride", "Ferroboride", color.Gray},
	PredefinedMaterialRecord{"ferrous_oxide", "Ferrous Oxide", color.Gray},
	PredefinedMaterialRecord{"ferrous_sulfate_dosimeter_solution", "Ferrous Sulfate Dosimeter Solution", color.Gray},
	PredefinedMaterialRecord{"freon_12", "Freon-12", color.Gray},
	PredefinedMaterialRecord{"freon_12b2", "Freon-12B2", color.Gray},
	PredefinedMaterialRecord{"freon_13", "Freon-13", color.Gray},
	PredefinedMaterialRecord{"freon_13b1", "Freon-13B1", color.Gray},
	PredefinedMaterialRecord{"freon_13i1", "Freon-13I1", color.Gray},
	PredefinedMaterialRecord{"gadolinium_oxysulfide", "Gadolinium Oxysulfide", color.Gray},
	PredefinedMaterialRecord{"gallium_arsenide", "Gallium Arsenide", color.Gray},
	PredefinedMaterialRecord{"gel_in_photographic_emulsion", "Gel in Photographic Emulsion", color.Gray},
	PredefinedMaterialRecord{"glass_lead", "Glass, Lead", color.Gray},
	PredefinedMaterialRecord{"glass_plate", "Glass, Plate", color.Gray},
	PredefinedMaterialRecord{"glass_pyrex", "Glass, Pyrex", color.Gray},
	PredefinedMaterialRecord{"glucose", "Glucose", color.Gray},
	PredefinedMaterialRecord{"glutamine", "Glutamine", color.Gray},
	PredefinedMaterialRecord{"glycerol", "Glycerol", color.Gray},
	PredefinedMaterialRecord{"guanine", "Guanine", color.Gray},
	PredefinedMaterialRecord{"gypsum", "Gypsum, Plaster of Paris", color.Gray},
	PredefinedMaterialRecord{"n_heptane", "N-Heptane", color.Gray},
	PredefinedMaterialRecord{"n_hexane", "N-Hexane", color.Gray},
	PredefinedMaterialRecord{"kapton_polyimide_film", "Kapton Polyimide Film", color.Gray},
	PredefinedMaterialRecord{"lanthanum_oxybromide", "Lanthanum Oxybromide", color.Gray},
	PredefinedMaterialRecord{"lanthanum_oxysulfide", "Lanthanum Oxysulfide", color.Gray},
	PredefinedMaterialRecord{"lead_oxide", "Lead Oxide", color.Gray},
	PredefinedMaterialRecord{"lithium_amide", "Lithium Amide", color.Gray},
	PredefinedMaterialRecord{"lithium_carbonate", "Lithium Carbonate", color.Gray},
	PredefinedMaterialRecord{"lithium_fluoride", "Lithium Fluoride", color.Gray},
	PredefinedMaterialRecord{"lithium_hydride", "Lithium Hydride", color.Gray},
	PredefinedMaterialRecord{"lithium_iodide", "Lithium Iodide", color.Gray},
	PredefinedMaterialRecord{"lithium_oxide", "Lithium Oxide", color.Gray},
	PredefinedMaterialRecord{"lithium_tetraborate", "Lithium Tetraborate", color.Gray},
	PredefinedMaterialRecord{"lung", "Lung (ICRP)", color.Gray},
	PredefinedMaterialRecord{"m3_wax", "M3 Wax", color.Gray},
	PredefinedMaterialRecord{"magnesium_carbonate", "Magnesium Carbonate", color.Gray},
	PredefinedMaterialRecord{"magnesium_fluoride", "Magnesium Fluoride", color.Gray},
	PredefinedMaterialRecord{"magnesium_oxide", "Magnesium Oxide", color.Gray},
	PredefinedMaterialRecord{"magnesium_tetraborate", "Magnesium Tetraborate", color.Gray},
	PredefinedMaterialRecord{"mercuric_iodide", "Mercuric Iodide", color.Gray},
	PredefinedMaterialRecord{"methane", "Methane", color.Gray},
	PredefinedMaterialRecord{"methanol", "Methanol", color.Gray},
	PredefinedMaterialRecord{"mix_d_wax", "Mix D Wax", color.Gray},
	PredefinedMaterialRecord{"ms20_tissue_substitute", "MS20 Tissue Substitute", color.Gray},
	PredefinedMaterialRecord{"muscle_skeletal", "Muscle, Skeletal", color.Gray},
	PredefinedMaterialRecord{"muscle_striated", "Muscle, Striated", color.Gray},
	PredefinedMaterialRecord{"muscle_equivalent_liquid_with_sucrose", "Muscle-Equivalent Liquid, with Sucrose", color.Gray},
	PredefinedMaterialRecord{"muscle_equivalent_liquid_without_sucrose", "Muscle-Equivalent Liquid, without Sucrose", color.Gray},
	PredefinedMaterialRecord{"naphthalene", "Naphthalene", color.Gray},
	PredefinedMaterialRecord{"nitrobenzene", "Nitrobenzene", color.Gray},
	PredefinedMaterialRecord{"nitrousr_oxide", "Nitrous Oxide", color.Gray},
	PredefinedMaterialRecord{"nylon_du_pont_elvamide_8062", "Nylon, Du Pont ELVAmide 8062", color.Gray},
	PredefinedMaterialRecord{"nylon_type_6_and_type_6/6", "Nylon, type 6 and type 6/6", color.Gray},
	PredefinedMaterialRecord{"nylon_type_6/10", "Nylon, type 6/10", color.Gray},
	PredefinedMaterialRecord{"nylon_type_11_rilsan", "Nylon, type 11 (Rilsan)", color.Gray},
	PredefinedMaterialRecord{"octane_liquid", "Octane, Liquid", color.Gray},
	PredefinedMaterialRecord{"paraffin_wax", "Paraffin Wax", color.Gray},
	PredefinedMaterialRecord{"n_pentane", "N-Pentane", color.Gray},
	PredefinedMaterialRecord{"photographic_emulsion", "Photographic Emulsion", color.Gray},
	PredefinedMaterialRecord{"plastic_scintillator", "Plastic Scintillator (Vinyltoluene based)", color.Gray},
	PredefinedMaterialRecord{"plutonium_dioxide", "Plutonium Dioxide", color.Gray},
	PredefinedMaterialRecord{"polyacrylonitrile", "Polyacrylonitrile", color.Gray},
	PredefinedMaterialRecord{"polycarbonate", "Polycarbonate (Makrolon, Lexan)", color.Gray},
	PredefinedMaterialRecord{"polychlorostyrene", "Polychlorostyrene", color.Gray},
	PredefinedMaterialRecord{"polyethylene", "Polyethylene", color.Gray},
	PredefinedMaterialRecord{"polyethylene_terephthalate", "Polyethylene Terephthalate (Mylar)", color.Gray},
	PredefinedMaterialRecord{"polymethyl_methacralate", "Polymethyl Methacralate (Lucite, Perspex)", color.Gray},
	PredefinedMaterialRecord{"polyoxymethylene", "Polyoxymethylene", color.Gray},
	PredefinedMaterialRecord{"polypropylene", "Polypropylene", color.Gray},
	PredefinedMaterialRecord{"polystyrene", "Polystyrene", color.Gray},
	PredefinedMaterialRecord{"polytetrafluoroethylene", "Polytetrafluoroethylene (Teflon)", color.Gray},
	PredefinedMaterialRecord{"polytrifluorochloroethylene", "Polytrifluorochloroethylene", color.Gray},
	PredefinedMaterialRecord{"polyvinyl_acetate", "Polyvinyl Acetate", color.Gray},
	PredefinedMaterialRecord{"polyvinyl_alcohol", "Polyvinyl Alcohol", color.Gray},
	PredefinedMaterialRecord{"polyvinyl_butyral", "Polyvinyl Butyral", color.Gray},
	PredefinedMaterialRecord{"polyvinyl_chloride", "Polyvinyl Chloride", color.Gray},
	PredefinedMaterialRecord{"polyvinylidene_chloride_saran", "Polyvinylidene Chloride, Saran", color.Gray},
	PredefinedMaterialRecord{"polyvinylidene_fluoride", "Polyvinylidene Fluoride", color.Gray},
	PredefinedMaterialRecord{"polyvinyl_pyrrolidone", "Polyvinyl Pyrrolidone", color.Gray},
	PredefinedMaterialRecord{"potassium_iodide", "Potassium Iodide", color.Gray},
	PredefinedMaterialRecord{"potassium_oxide", "Potassium Oxide", color.Gray},
	PredefinedMaterialRecord{"propane", "Propane", color.Gray},
	PredefinedMaterialRecord{"propane_liquid", "Propane, Liquid", color.Gray},
	PredefinedMaterialRecord{"n_propyl_alcohol", "N-Propyl Alcohol", color.Gray},
	PredefinedMaterialRecord{"pyridine", "Pyridine", color.Gray},
	PredefinedMaterialRecord{"rubber_butyl", "Rubber, Butyl", color.Gray},
	PredefinedMaterialRecord{"rubber_natural", "Rubber, Natural", color.Gray},
	PredefinedMaterialRecord{"rubber_neoprene", "Rubber, Neoprene", color.Gray},
	PredefinedMaterialRecord{"silicon_dioxide", "Silicon Dioxide", color.Gray},
	PredefinedMaterialRecord{"silver_bromide", "Silver Bromide", color.Gray},
	PredefinedMaterialRecord{"silver_chloride", "Silver Chloride", color.Gray},
	PredefinedMaterialRecord{"silver_halides_in_photographic_emulsion", "Silver Halides in Photographic Emulsion", color.Gray},
	PredefinedMaterialRecord{"silver_iodide", "Silver Iodide", color.Gray},
	PredefinedMaterialRecord{"skin", "Skin (ICRP)", color.Gray},
	PredefinedMaterialRecord{"sodium_carbonate", "Sodium Carbonate", color.Gray},
	PredefinedMaterialRecord{"sodium_iodide", "Sodium Iodide", color.Gray},
	PredefinedMaterialRecord{"sodium_monoxide", "Sodium Monoxide", color.Gray},
	PredefinedMaterialRecord{"sodium_nitrate", "Sodium Nitrate", color.Gray},
	PredefinedMaterialRecord{"stilbene", "Stilbene", color.Gray},
	PredefinedMaterialRecord{"sucrose", "Sucrose", color.Gray},
	PredefinedMaterialRecord{"terphenyl", "Terphenyl", color.Gray},
	PredefinedMaterialRecord{"testes", "Testes (ICRP)", color.Gray},
	PredefinedMaterialRecord{"tetrachloroethylene", "Tetrachloroethylene", color.Gray},
	PredefinedMaterialRecord{"thallium_chloride", "Thallium Chloride", color.Gray},
	PredefinedMaterialRecord{"tissue_soft_icrp", "Tissue, Soft (ICRP)", color.Gray},
	PredefinedMaterialRecord{"tissue_soft_icru_four_component", "Tissue, Soft (ICRU four-component)", color.Gray},
	PredefinedMaterialRecord{"tissue_equivalent_gas_methane_based", "Tissue-Equivalent GAS (Methane based)", color.Gray},
	PredefinedMaterialRecord{"tissue_equivalent_gas_propane_based", "Tissue-Equivalent GAS (Propane based)", color.Gray},
	PredefinedMaterialRecord{"titanium_dioxide", "Titanium Dioxide", color.Gray},
	PredefinedMaterialRecord{"toluene", "Toluene", color.Gray},
	PredefinedMaterialRecord{"trichloroethylene", "Trichloroethylene", color.Gray},
	PredefinedMaterialRecord{"triethyl_phosphate", "Triethyl Phosphate", color.Gray},
	PredefinedMaterialRecord{"tungsten_hexafluoride", "Tungsten Hexafluoride", color.Gray},
	PredefinedMaterialRecord{"uranium_dicarbide", "Uranium Dicarbide", color.Gray},
	PredefinedMaterialRecord{"uranium_monocarbide", "Uranium Monocarbide", color.Gray},
	PredefinedMaterialRecord{"uranium_oxide", "Uranium Oxide", color.Gray},
	PredefinedMaterialRecord{"urea", "Urea", color.Gray},
	PredefinedMaterialRecord{"valine", "Valine", color.Gray},
	PredefinedMaterialRecord{"viton_fluoroelastomer", "Viton Fluoroelastomer", color.Gray},
	PredefinedMaterialRecord{"water_liquid", "Water, Liquid", waterColor},
	PredefinedMaterialRecord{"water_vapor", "Water Vapor", waterColor},
	PredefinedMaterialRecord{"xylene", "Xylene", color.Gray},
}

var isotopes = []IsotopeRecord{
	IsotopeRecord{"h-1 - hydrogen", "H-1 - Hydrogen"},
	IsotopeRecord{"h-2 - deuterium", "H-2 - Deuterium"},
	IsotopeRecord{"h-3 - tritium", "H-3 - Tritium"},
	IsotopeRecord{"he-3", "He-3"},
	IsotopeRecord{"he-4", "He-4"},
	IsotopeRecord{"li-6", "Li-6"},
	IsotopeRecord{"li-7", "Li-7"},
	IsotopeRecord{"be-9", "Be-9"},
	IsotopeRecord{"b-10", "B-10"},
	IsotopeRecord{"b-11", "B-11"},
	IsotopeRecord{"c-*", "C-*"},
	IsotopeRecord{"n-*", "N-*"},
	IsotopeRecord{"o-*", "O-*"},
	IsotopeRecord{"f-19", "F-19"},
	IsotopeRecord{"na-23", "Na-23"},
	IsotopeRecord{"mg-*", "Mg-*"},
	IsotopeRecord{"al-27", "Al-27"},
	IsotopeRecord{"si-*", "Si-*"},
	IsotopeRecord{"p-31", "P-31"},
	IsotopeRecord{"s-*", "S-*"},
	IsotopeRecord{"cl-*", "Cl-*"},
	IsotopeRecord{"ar-*", "Ar-*"},
	IsotopeRecord{"k-*", "K-*"},
	IsotopeRecord{"ca-*", "Ca-*"},
	IsotopeRecord{"ti-*", "Ti-*"},
	IsotopeRecord{"v-51", "V-51"},
	IsotopeRecord{"cr-*", "Cr-*"},
	IsotopeRecord{"mn-55", "Mn-55"},
	IsotopeRecord{"fe-*", "Fe-*"},
	IsotopeRecord{"co-59", "Co-59"},
	IsotopeRecord{"ni-*", "Ni-*"},
	IsotopeRecord{"cu-*", "Cu-*"},
	IsotopeRecord{"zn-*", "Zn-*"},
	IsotopeRecord{"ga-*", "Ga-*"},
	IsotopeRecord{"ge-*", "Ge-*"},
	IsotopeRecord{"as-75", "As-75"},
	IsotopeRecord{"nb-93", "Nb-93"},
	IsotopeRecord{"mo-*", "Mo-*"},
	IsotopeRecord{"ag-*", "Ag-*"},
	IsotopeRecord{"cd-*", "Cd-*"},
	IsotopeRecord{"sn-*", "Sn-*"},
	IsotopeRecord{"eu-*", "Eu-*"},
	IsotopeRecord{"gd-*", "Gd-*"},
	IsotopeRecord{"er-*", "Er-*"},
	IsotopeRecord{"ta-181", "Ta-181"},
	IsotopeRecord{"w-*", "W-*"},
	IsotopeRecord{"re-*", "Re-*"},
	IsotopeRecord{"au-187", "Au-187"},
	IsotopeRecord{"hg-*", "Hg-*"},
	IsotopeRecord{"pb-*", "Pb-*"},
	IsotopeRecord{"bi-209", "Bi-209"},
	IsotopeRecord{"th-232", "Th-232"},
	IsotopeRecord{"u-235", "U-235"},
	IsotopeRecord{"u-238", "U-238"},
	IsotopeRecord{"pu-239", "Pu-239"},
	IsotopeRecord{"pu-240", "Pu-240"},
}

// PredefinedMaterials return list of all available material.Predefined and Colors assigned to them.
func PredefinedMaterials() []PredefinedMaterialRecord {
	return predefinedMaterials
}

// Isotopes return list of all available isotopes for material.Element.
func Isotopes() []IsotopeRecord {
	return isotopes
}
