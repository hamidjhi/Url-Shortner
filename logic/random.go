package logic

import (
	"encoding/base64"
	"linkshorter/db"
	"linkshorter/models"
	"log"
)

func UrlChanger(url models.Url) (models.Url,error) {
	var a = url
	Shorted := base64.StdEncoding.EncodeToString([]byte(a.LongUrl))

	err := db.SetRedis(Shorted, url.LongUrl)

	if err != nil {
		log.Println("cannot generated")
	} else {
		log.Println("generated")
	}
	return models.Url{
		LongUrl: url.LongUrl,
		ShortUrl: Shorted,
	},nil
}
