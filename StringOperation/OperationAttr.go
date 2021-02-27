package StringOperation

import (
	"fmt"
	"goft-redis/Result"
	"time"
)

//Attr 定义的是 set时的其它属性

const (
	EXPR = "expr" //过期时间
	NX   = "nx"   //setNX锁
	XX   = "xx"   //setXX锁
)

type OperationAttr struct {
	Name  string      //根据传入的属性名来判定执行什么
	Value interface{} //具体的属性
}

type OperationAttrs []*OperationAttr

//根据 OperationAttrs 这个切片，创建Find方法
//Find方法是用来从 attrs 的传参中取值的
//会传入常量定义的 EXPR 等数据，根据这些string数据，来获取对应的 Value
func (this OperationAttrs) Find(name string) *Result.InterfaceResult {
	for _, attr := range this {
		if attr.Name == name {
			return Result.NewInterfaceResut(attr.Value, nil)
		}
	}
	return Result.NewInterfaceResut(nil, fmt.Errorf("Operation find error :%s", name))
}

//Expire 设置set时的超时时间属性
//传入 expr，则 Value 就是超时时间的具体时长。
func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{Name: EXPR, Value: t}
}

//SetNX命令
//Value 建议传入一个 空结构，空结构体不占空间
func WithNX() *OperationAttr {
	return &OperationAttr{Name: NX, Value: struct{}{}}
}

//setXX命令
func WithXX() *OperationAttr {
	return &OperationAttr{Name: XX, Value: struct{}{}}
}
