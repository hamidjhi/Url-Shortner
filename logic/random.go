package logic

import (
	"encoding/base64"
	"linkshorter/db"
	"linkshorter/models"
	"log"
)

func UrlChanger(url models.Url) (models.Url, error) {
	var a = url
	Shorted := base64.StdEncoding.EncodeToString([]byte(a.LongUrl))

	err := db.SetRedis(Shorted, url.LongUrl)

	if err != nil {
		log.Println("cannot generated")
	} else {
		log.Println("generated")
	}
	return models.Url{
		LongUrl:  url.LongUrl,
		ShortUrl: Shorted,
	}, nil
}

func UrlReturn(shortLink string) (models.Url, error) {
	long, err := db.GetRedis(shortLink)
	if err != nil {
		log.Println("cannot do that ")
	} else {
		log.Println("find it!")
	}

	return models.Url{
		LongUrl:  long.LongUrl,
		ShortUrl: "",
	}, nil

}
