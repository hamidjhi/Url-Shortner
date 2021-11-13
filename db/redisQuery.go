package db

import (
	"linkshorter/models"
	"log"
	"time"
)

func SetRedis(shorted, url string) error {
	ep := Db.Set(shorted, url, time.Hour)
	log.Println(ep.Err())

	return nil
}
func GetRedis(shortLink string) (models.Url, error) {
	ep, err := Db.Get(shortLink).Result()
	if err != nil {
		log.Println("cannot res")
	}

	return models.Url{
		LongUrl:  ep,
		ShortUrl: shortLink,
	}, nil
}
