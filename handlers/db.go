package handlers

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type ConnRedis struct {
	DB redis.Conn
}

func (h *ConnRedis) Connect() {
	var err error
	h.DB, err = redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
}
