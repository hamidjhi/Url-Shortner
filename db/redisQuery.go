package db

import (
	"log"
	"time"
)

func SetRedis(shorted, url string)error  {
	ep:=Db.Set( shorted,url,time.Hour)
	log.Println(ep.Err())

	return nil
}
