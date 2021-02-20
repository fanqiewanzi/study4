package list

import (
	"errors"
	"fmt"
	"sync"
)

//定义两个常量一个用来存默认容量，一个用来存初始下标
const defaultCapacity = 10
const defaultSize = -1

var mut sync.Mutex

//动态数组结构
//size	当前最后一个元素下标位置
//capacity	数组最大容量
//data	数组
type Array struct {
	size     int
	capacity int
	data     []interface{}
}

//ArrayList迭代器
//array为原数组
//cursor为一个移动的下标
//end为最近已经输出元素的下标
type ArrayIterator struct {
	array  *Array
	cursor int
	end    int
}

//创建一个新的动态数组
func NewArray(capacity int) *Array {
	data := make([]interface{}, capacity)
	p := &Array{defaultSize, capacity, data}
	return p
}

//在不输入最大容量时创建
func NewArrayWithoutNoCap() *Array {
	data := make([]interface{}, defaultCapacity)
	return &Array{defaultSize, defaultCapacity, data}
}

//扩展数组
func (array *Array) grow(num int) (bool, []interface{}) {
	data := make([]interface{}, num)
	//将原数组中的值赋给新扩充容量的数组
	for i, elem := range array.data {
		data[i] = elem
	}
	return true, data
}

//检查数组容量是否足够
func (array *Array) check(num int) bool {
	//若存储大小加新数组的大小大于最大容量则进行扩充
	//这里只对数组进行扩充
	if array.size+num+1 > array.capacity {
		//定义新数组最大容量，这里是原数组长度加新元素长度再加一个元素的最大容量
		newCap := array.size + num + 2
		ok, array1 := array.grow(newCap)
		if ok == true {
			array.data = array1
			array.capacity = array.size + num + 2
			return true
		}
	}
	return false
}

//打印数组
func (array Array) Print() {
	for i := 0; i <= array.size; i++ {
		fmt.Print(array.data[i])
		fmt.Print("\t")
	}
}

//向list末尾加入元素，先检查容量是否足够再进行添加
func (array *Array) Add(obj ...interface{}) error {
	mut.Lock()
	defer mut.Unlock()
	array.check(len(obj))
	for _, elem := range obj {
		array.size++
		array.data[array.size] = elem
	}
	return nil
}

//向指定位置加入元素,要判断容量是否足够且目标位置存在再进行添加
func (array *Array) Insert(location int, obj interface{}) error {
	mut.Lock()
	defer mut.Unlock()
	if location <= 0 || location > array.size+1 {
		return errors.New("下标超出")
	}
	array.check(1)
	for i := array.size; i >= location-1; i-- {
		array.data[i+1] = array.data[i]
	}
	array.data[location-1] = obj
	array.size++
	return nil
}

//向指定位置修改元素，要判断目标位置是否存在
func (array *Array) Set(location int, obj interface{}) error {
	mut.Lock()
	defer mut.Unlock()
	if location <= 0 || location > array.size+1 {
		return errors.New("下标超出")
	}
	array.data[location-1] = obj
	return nil
}

//是否存在某元素，进入循环进行遍历存在返回true，否则返回false
func (array Array) Contain(obj interface{}) bool {
	for _, i := range array.data {
		if i == obj {
			return true
		}
	}
	return false
}

//是否为空，判断size如果为-1就说明数组中没有元素返回true,否则返回false
func (array Array) IsEmpty() bool {
	if array.size == -1 {
		return true
	}
	return false
}

//查看某一位置上的元素,判断目标位置是否在范围内
func (array *Array) Get(location int) (interface{}, error) {
	if location <= 0 || location > array.size+1 {
		return nil, errors.New("下标超出")
	}
	return array.data[location-1], nil
}

//判断是否相等
//这里使用迭代器对里面元素进行一一比较
func (array *Array) Equals(list List) bool {
	if array.Size() != list.Size() {
		return false
	}
	it := array.Iterator()
	it1 := list.Iterator()
	for it.HasNext() && it1.HasNext() {
		elem, _ := it.Next()
		elem1, _ := it1.Next()
		if elem != elem1 {
			break
		}
	}
	if it.HasNext() == false && it1.HasNext() == false {
		return true
	}
	return false
}

//转换为Slice类型，data就是slice所以直接赋值过去
func (array Array) ToSlice() []interface{} {
	return array.data
}

//输出当前list的长度
func (array Array) Size() int {
	return array.size + 1
}

//迭代器方法
//判断是否存在下一个元素
func (it *ArrayIterator) HasNext() bool {
	//如果当前下标等于它的大小说明没有下一个元素了
	return it.cursor != it.array.Size()
}

//返回下一个元素
func (it *ArrayIterator) Next() (interface{}, error) {
	//首先获取当前下标的位置
	i := it.cursor
	if i >= it.array.Size() {
		return nil, errors.New("没有这样的索引")
	}
	//下标位置往后移
	it.cursor = it.cursor + 1
	it.end = i
	return it.array.data[it.end], nil
}

//是否有上一个元素
func (it *ArrayIterator) HasPrevious() bool {
	return it.cursor != -1
}

//返回上一个元素
func (it *ArrayIterator) Previous() (interface{}, error) {
	//首先获取当前下标的位置
	i := it.cursor
	if i < 0 {
		return nil, errors.New("没有这样的索引")
	}
	//下标位置往后移
	it.cursor = it.cursor - 1
	it.end = i
	return it.array.data[it.end], nil
}

//下一个下标
func (it *ArrayIterator) NextIndex() (interface{}, error) {
	i := it.cursor
	if i >= it.array.size+1 {
		return nil, errors.New("没有这样的索引")
	}
	return it.cursor + 1, nil
}

//返回上一个下标
func (it *ArrayIterator) PreviousIndex() (interface{}, error) {
	i := it.cursor
	if i < 0 {
		return nil, errors.New("没有这样的索引")
	}
	return it.cursor - 1, nil
}

//移除元素
func (it *ArrayIterator) Remove() error {
	mut.Lock()
	defer mut.Unlock()
	for i := it.end; i < it.array.size; i++ {
		it.array.data[i] = it.array.data[i+1]
	}
	it.array.data[it.array.size] = nil
	it.array.size--
	it.cursor--
	it.end--
	return nil
}

//在当前节点的前一个已经输出的节点赋值
func (it *ArrayIterator) Set(elem interface{}) error {
	mut.Lock()
	defer mut.Unlock()
	it.array.data[it.end] = elem
	return nil
}

//在当前节点的前一个已经输出的节点添加元素
func (it *ArrayIterator) Add(elem interface{}) error {
	mut.Lock()
	defer mut.Unlock()
	if it.end < 0 {
		return errors.New("列表为空")
	}
	it.array.Insert(it.end+1, elem)
	return nil
}

//迭代器初始化
func (array *Array) Iterator() Iterator {
	it := new(ArrayIterator)
	it.array = array
	it.cursor = 0
	it.end = -1
	return it
}
