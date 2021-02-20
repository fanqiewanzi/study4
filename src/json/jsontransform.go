package json

import (
	"encoding/json"
	"fmt"
)

//结构体中的元素必须大写不然转换不了
type Person struct {
	Name   string
	Age    int
	Weight float32
}

func JsonTest() {
	//json字符串
	str := "{\"name\":\"bob\",\"age\":32,\"weight\":54.3}"
	//定义一个用于存储转化后的字符串的对象
	man := Person{}
	//进行转换
	err := json.Unmarshal([]byte(str), &man)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	fmt.Println(man)
}
