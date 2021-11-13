package db

import (
	"github.com/go-redis/redis"
	"log"
)

var Db *redis.Client

func ConnectToRedis() error {
	Db = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	db, err := Db.Ping().Result()
	if err != nil {
		log.Println("cannot connect")
	} else {
		log.Println("coonected", db)
	}

	return nil
}
