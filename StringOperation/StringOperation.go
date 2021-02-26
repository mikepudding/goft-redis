package StringOperation

import (
	"context"
	"fmt"
	"goft-redis/Result"
	"goft-redis/gedis"
	"time"
)

//专门处理 string类型的操作

type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (this *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *Result.InterfaceResult {

	//set创建的超时时间
	expr := OperationAttrs(attrs).Find(EXPR)
	exp := expr.Unwrap_Or(time.Second * 0).(time.Duration) //如果有错默认为0秒。

	//setNX锁的方式创建 string类型的 key:value数据
	nx := OperationAttrs(attrs).Find(NX).Unwrap_Or("nil")
	fmt.Println("到这333")
	if nx != nil {
		return Result.NewInterfaceResut(gedis.Redis().SetNX(this.ctx, key, value, exp).Result())
	}

	//setXX锁的方式创建，如果该key存在才能创建成功
	xx := OperationAttrs(attrs).Find(XX).Unwrap_Or("nil")
	fmt.Println("到这222")
	if xx != nil {
		return Result.NewInterfaceResut(gedis.Redis().SetXX(this.ctx, key, value, exp).Result())
	}

	//默认是 set方法
	return Result.NewInterfaceResut(gedis.Redis().Set(this.ctx, key, value, exp).Result())
}

func (this *StringOperation) Get(key string) *Result.StringResult {
	return Result.NewStringResult(gedis.Redis().Get(this.ctx, key).Result())
}

func (this *StringOperation) Mget(key ...string) *Result.SliceResult {
	return Result.NewSliceResult(gedis.Redis().MGet(this.ctx, key...).Result())
}
