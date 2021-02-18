package iniparser

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

type IniConfig struct {
	filename string
	section  map[string]map[string]string
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
	for {
		//读取一行缓存
		line, _, err := buf.ReadLine()

		if err == io.EOF {
			//读到文件结尾，退出循环
			break
		}
		if bytes.Equal(line, []byte("")) {
			//空行直接跳过循环
			continue
		}
		// HasPrefix测试字符串是否以给定字符开头
		// HasSuffix测试字符串是否以给定字符结尾
		if bytes.HasPrefix(line, []byte("[")) && bytes.HasSuffix(line, []byte("]")) {
			//提取出section的整个字符串
			//line[i:k]从line中提取第i个字符到k个字符组成新的slice
			section := string(line[1 : len(line)-1])
			//判断map中是否已经存在该key值的section,不存在就创建新的key
			if _, ok := cf.section[section]; !ok {
				cf.section[section] = make(map[string]string)
			}
		} else {

		}
	}
}
