package list

//迭代器接口
type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	HasPrevious() bool
	Previous() (interface{}, error)
	NextIndex() (interface{}, error)
	PreviousIndex() (interface{}, error)
	Remove() error
	Set(interface{}) error
	Add(interface{}) error
}
