package Result

import "goft-redis/Iterator"

type SliceResult struct {
	Result []interface{} //是redis的返回结果
	Error  error         //是redis返回的错误信息
}

//凡是所有的返回值都会返回 StringResult 的两个结构体字段，&{Result，Error}
func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Error: err}
}

//如果有错误了，则所有的err错误都在 Unwrap里判断并返回
//如果执行 Unwrap但是并没有错误，则只返回 Result
func (this *SliceResult) Unwrap() []interface{} {
	if this.Error != nil {
		panic(this.Error)
	}
	return this.Result
}

//带默认值得返回值封装。
//如果发现结果是错误的，不返回报错，只返回默认值
func (this *SliceResult) Unwrap_Or(str string) interface{} {
	if this.Error != nil {
		return str
	}
	return this.Result
}

//调用迭代器的方法
func (this *SliceResult) Iter() *Iterator.Iterator {
	return Iterator.NewIterator(this.Result)
}
