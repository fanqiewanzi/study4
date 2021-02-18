package list

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func groutine1(chs chan List) {
	mutex.Lock()
	defer mutex.Unlock()
	elem := <-chs
	it := elem.Iterator()
	elem.Add(1)
	fmt.Print("groutine1\t")
	for it.HasNext() {
		i, _ := it.Next()
		fmt.Print(i)
	}
	fmt.Println()

}

func groutine2(chs chan List) {
	mutex.Lock()
	defer mutex.Unlock()
	elem := <-chs
	it := elem.Iterator()
	elem.Add(2)
	fmt.Print("groutine2\t")
	for it.HasNext() {
		i, _ := it.Next()
		fmt.Print(i)
	}
	fmt.Println()

}

func ListTest() {
	array := NewArrayWithoutNoCap()
	chs := make(chan List, 1)
	for i := 0; i < 10; i++ {
		chs <- array
		go groutine1(chs)
		go groutine2(chs)
	}
	time.Sleep(2 * time.Second)
}
