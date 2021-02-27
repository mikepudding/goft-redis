package lib

import (
	"goft-redis/Cache"
	"goft-redis/StringOperation"
	"goft-redis/policy"
	"sync"
	"time"
)

var NewsCachePool *sync.Pool

func init() {
	NewsCachePool = &sync.Pool{
		New: func() interface{} {
			return Cache.NewSimpleCache(StringOperation.NewStringOperation(),
				time.Second*150,      //指定超时时间
				Cache.Serilizer_JSON, //指定序列化方式是json
				policy.NewCrossPolicy("^news\\d{1,5}$", time.Second*30),
			)
		},
	}
}

//返回SimpleCache缓存组件的链接
func NewsCache() *Cache.SimpleCache {
	return NewsCachePool.Get().(*Cache.SimpleCache)
}

//释放链接
func ReleaseNewsCache(cache *Cache.SimpleCache) {
	NewsCachePool.Put(cache)
}
