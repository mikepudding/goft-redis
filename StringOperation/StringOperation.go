package StringOperation

import (
	"context"
	"goft-redis/Result"
	"goft-redis/gedis"
	"time"
)

//专门处理string类型的操作
type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (this *StringOperation) Set(key string, value interface{},
	attrs ...*OperationAttr) *Result.InterfaceResult {
	exp := OperationAttrs(attrs).
		Find(EXPR).
		Unwrap_Or(time.Second * 0).(time.Duration)

	nx := OperationAttrs(attrs).Find(NX).Unwrap_Or(nil)
	if nx != nil {
		return Result.NewInterfaceResut(gedis.Redis().SetNX(this.ctx, key, value, exp).Result())
	}
	xx := OperationAttrs(attrs).Find(XX).Unwrap_Or(nil)
	if xx != nil {
		return Result.NewInterfaceResut(gedis.Redis().SetXX(this.ctx, key, value, exp).Result())
	}
	return Result.NewInterfaceResut(gedis.Redis().Set(this.ctx, key, value,
		exp).Result())

}
func (this *StringOperation) Get(key string) *Result.StringResult {
	return Result.NewStringResult(gedis.Redis().Get(this.ctx, key).Result())
}
func (this *StringOperation) MGet(keys ...string) *Result.SliceResult {
	return Result.NewSliceResult(gedis.Redis().MGet(this.ctx, keys...).Result())
}
