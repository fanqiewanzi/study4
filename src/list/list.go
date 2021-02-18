package list

type List interface {
	//向list末尾加入元素
	Add(obj ...interface{}) error
	//向指定位置加入元素
	Insert(location int, obj interface{}) error
	//向指定位置修改元素
	Set(location int, obj interface{}) error
	//是否存在某元素
	Contain(obj interface{}) bool
	//是否为空
	IsEmpty() bool
	//查看某一位置上的元素
	Get(location int) (interface{}, error)
	//判断是否相等
	Equals(list List) bool
	//转换为Slice类型
	ToSlice() []interface{}
	//输出当前list的长度
	Size() int
	//迭代器
	Iterator() Iterator
}
