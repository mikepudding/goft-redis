package main

import (
	"goft-redis/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	GinRun()
}

func GinRun() {
	r := gin.New()
	r.Handle("GET", "/news/:id", func(c *gin.Context) {
		//1.从Pool对象池，获取新闻缓存对象
		newcache := lib.NewsCache()
		defer lib.ReleaseNewsCache(newcache)

		//2.获取缓存参数，设置DBGetter
		newsID := c.Param("id")
		newcache.DBGetter = lib.NewsDBGetter(newsID)

		//3.取出缓存参数（如果没有。则通过上面的代码从数据库中获取）
		c.Header("Content-type", "application/json")
		c.String(200, newcache.GetCache("news"+newsID).(string))

		/*	使用 GOB序列化输出
			newsModel := lib.NewNewsModel()
			newsCache.GetCacheForObject("news"+newsID,newsModel)
			context.JSON(200,newsModel)
		*/
	})
	r.Run(":8081")
}
