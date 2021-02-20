package iniparser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

//ini配置文件的数据结构
//filename为文件位置
//section则是一个以section名指向元素的map
type IniConfig struct {
	filename string
	section  map[string]*IniSection
}
type IniSection struct {
	elem map[string]string
}

func Parse(cf *IniConfig) error {
	//打开文件
	file, err := os.Open(cf.filename)
	if err != nil {
		return err
	}
	//关闭文件
	defer file.Close()
	buf := bufio.NewReader(file)
	section := ""
	for {
		//读取一行缓存
		line, _, err := buf.ReadLine()
		//读到文件结尾，退出循环
		if err == io.EOF {
			break
		}
		//空行或注释直接跳过循环
		if bytes.Equal(line, []byte("")) || bytes.HasPrefix(line, []byte(";")) {
			continue
		}
		// HasPrefix测试字符串是否以给定字符开头
		// HasSuffix测试字符串是否以给定字符结尾
		if bytes.HasPrefix(line, []byte("[")) && bytes.HasSuffix(line, []byte("]")) {
			//提取出section的整个字符串
			//line[i:k]从line中提取第i个字符到k个字符组成新的slice
			section = string(line[1 : len(line)-1])
			//判断map中是否已经存在该key值的section,不存在就创建新的key
			if _, ok := cf.section[section]; !ok {
				cf.section[section] = &IniSection{make(map[string]string)}
			}
		} else {
			//bytes.SplitN(s, sep []byte, n int)方法会返回以sep为基础将s字符串进行分割成n个子字符串
			str := bytes.SplitN(line, []byte("="), 2)
			//若是新的section就创建一个新的Map指向指针
			if _, ok := cf.section[section]; !ok {
				cf.section[section] = &IniSection{make(map[string]string)}
			}
			//将值赋给map
			cf.section[section].elem[string(str[0])] = string(str[1])
		}
	}
	return nil
}

func IniTest() {
	cf := &IniConfig{"D:\\GoProject\\study4\\src\\iniparser\\demo.ini", make(map[string]*IniSection)}
	Parse(cf)
	for j, elem := range cf.section {
		fmt.Println(j)
		for i, ele := range elem.elem {
			fmt.Println(i, ele)
		}
	}
}
