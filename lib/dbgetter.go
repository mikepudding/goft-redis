package lib

import (
	"goft-redis/Cache"
	"log"
)

func NewsDBGetter(id string) Cache.DBGetterFunc {
	return func() interface{} {
		log.Println("get from db")
		newsmodel := NewNewsModel()
		Gorm.Table("mynews").Where("id = ?", id).Find(newsmodel)
		return newsmodel
	}
}
