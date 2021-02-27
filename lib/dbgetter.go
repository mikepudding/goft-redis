package lib

import (
	"goft-redis/Cache"
	"log"
)

func NewsDBGetter(id string) Cache.DBGetterFunc {
	return func() interface{} {
		log.Println("get from db")
		newsmodel := NewNewsModel()
		if Gorm.Table("mynews").Where("id = ?", id).Find(newsmodel).
			Error != nil || newsmodel.NewsID <= 0 {
			return nil
		}
		return newsmodel
	}
}
