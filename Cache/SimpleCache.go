package Cache

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"goft-redis/StringOperation"
	"goft-redis/policy"
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
	Policy    policy.CachePolicy               //正则匹配缓存穿透监测
}

func NewSimpleCache(operation *StringOperation.StringOperation,
	expire time.Duration, serilizer string, policy policy.CachePolicy) *SimpleCache {

	policy.SetOperation(operation) //非常关键的一步
	return &SimpleCache{Operation: operation, Expire: expire, Serilizer: serilizer,
		Policy: policy}
}

//写入缓存
func (this *SimpleCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, StringOperation.WithExpire(this.Expire)).Unwrap()
}

//从缓存中读取数据
func (this *SimpleCache) GetCache(key string) (ret interface{}) {
	//1.缓存穿透的ID检查策略
	if this.Policy != nil {
		this.Policy.Before(key)
	}

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
	}

	//如果是空缓存则写入：key，空value
	if ret.(string) == "" && this.Policy != nil {
		this.Policy.IfNil(key, "")
	} else {
		this.SetCache(key, ret) //反之将存在的数据写入redis数据库
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
