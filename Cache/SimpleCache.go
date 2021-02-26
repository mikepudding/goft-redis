package Cache

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"goft-redis/StringOperation"
	"time"
)

type DBGetterFunc func() interface{}

const (
	Serilizer_JSON = "json"
	Serilizer_GOB  = "gob"
)

type SimpleCache struct {
	Operation *StringOperation.StringOperation //stirng类型
	Expire    time.Duration                    //超时时间
	DBGetter  DBGetterFunc                     //缓存命中数据库
	Serilizer string                           //序列化方式
}

func NewSimpleCache(operation *StringOperation.StringOperation,
	expire time.Duration, serilizer string) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire, Serilizer: serilizer}
}

//写入缓存
func (this *SimpleCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, StringOperation.WithExpire(this.Expire)).Unwrap()
}

//从缓存中读取数据
func (this *SimpleCache) GetCache(key string) (ret interface{}) {
	if this.Serilizer == Serilizer_JSON {
		f := func() string {
			obj := this.DBGetter() //拿到DBGetterFunc 获取的 newsmodel
			if obj == nil {
				return ""
			}
			b, err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return string(b)
		}
		ret = this.Operation.Get(key).Unwrap_Or_Else(f)
		//Get获取到数据结果后，别忘了写入缓存，SetCache
		this.SetCache(key, ret)
	} else if this.Serilizer == Serilizer_GOB {
		f := func() string {
			obj := this.DBGetter()
			if obj == nil {
				return ""
			}
			var buf = &bytes.Buffer{}
			enc := gob.NewEncoder(buf)
			if err := enc.Encode(obj); err != nil {
				return ""
			}
			return buf.String()
		}
		ret = this.Operation.Get(key).Unwrap_Or_Else(f)
		this.SetCache(key, ret)
	}

	return
}

func (this *SimpleCache) GetCacheForObject(key string, obj interface{}) interface{} {
	ret := this.GetCache(key)
	if ret == nil {
		return nil
	}
	if this.Serilizer == Serilizer_JSON {
		err := json.Unmarshal([]byte(ret.(string)), obj)
		if err != nil {
			return nil
		}
	} else if this.Serilizer == Serilizer_GOB {
		var buf = &bytes.Buffer{}
		buf.WriteString(ret.(string))
		dec := gob.NewDecoder(buf)
		if dec.Decode(obj) != nil {
			return nil
		}
	}
	return nil
}
