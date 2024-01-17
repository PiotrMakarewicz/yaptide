from abc import abstractmethod
import logging
from threading import Thread
from flask import Flask
from yaptide.redis.redis import redis_client

class RedisConsumerBase(Thread):
    def __init__(self, app: Flask, consumer_name: str, queue_name: str, batch_size: int):
        Thread.__init__(self)
        self.consumer_name = consumer_name
        self.queue_name = queue_name
        self.app = app
        self.batch_size = batch_size

    @abstractmethod
    def handle_message(self, messages) -> None:
        pass
    def log_message_info(self, message: str) -> None:
        logging.info(f"Consumer {self.consumer_name}: {message}")
    def log_message_error(self, message: str) -> None:
        logging.error(f"Consumer {self.consumer_name}: {message}")
    def run(self) -> None:
        self.log_message_info("starts working...")
        while True:
            message = redis_client.lpop(self.queue_name, count = batch_size)
            if(len(message) != 0):
                execute_handler
            else:
                time.sleep(1)

    def execute_handler(): 
            self.log_message_info(f"Received message: {message}")
            self.handle_message(message)
            self.log_message_info(f"Processed message successfully")