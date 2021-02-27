package main

import (
	"goft-redis/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				context.JSON(400, gin.H{"message": e})
			}
		}()
		context.Next()
	})
	r.Handle("GET", "/news/:id", func(context *gin.Context) {
		// 1、从对象池 获取新闻缓存 对象
		newsCache := lib.NewsCache()
		defer lib.ReleaseNewsCache(newsCache)

		//2、获取参数，设置DBGetter
		newsID := context.Param("id")
		newsCache.DBGetter = lib.NewsDBGetter(newsID) //一旦缓存没有，则需要从哪去取

		// 3、取缓存输出（一旦没有，上面的DBGetter会被调用)
		newsModel := lib.NewNewsModel()
		newsCache.GetCacheForObject("news"+newsID, newsModel)
		context.JSON(200, newsModel)
		// context.Header("Content-type", "application/json")
		// context.String(200, newsCache.GetCache("news"+newsID).(string))
	})

	r.Run(":8081")

}
