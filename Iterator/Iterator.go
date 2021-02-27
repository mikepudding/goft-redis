package Iterator

//迭代器,只有返回值是切片的才用到迭代器，遍历切片

type Iterator struct {
	data  []interface{}
	index int
}

//初始化迭代器，初始时传入data，也是需要迭代显示的数据切片
func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

//判断是否有值
func (this *Iterator) HasNext() bool {
	if this.data == nil || len(this.data) == 0 {
		return false
	}
	return this.index < len(this.data)
}

//遍历data
func (this *Iterator) Next() (ret interface{}) {
	ret = this.data[this.index]
	this.index = this.index + 1
	return
}
