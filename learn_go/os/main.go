package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	values, err := readFile("test.txt")
	if err != nil {
		fmt.Println("failed to read file.")
		return
	}
	writeFile(values)
}

func readFile(filename string) (string, error) {
	//使用os.Open(filename)打开文件
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Failed to open the input file")
		return "", err
	}
	//关闭资源
	defer file.Close()

	var str string
	//使用bufio.NewReader(file)读取数据
	byteReader := bufio.NewReader(file)
	for {
		line, isPrefix, err1 := byteReader.ReadLine()
		if err1 != nil {
			//需要判断下是否是读到末尾
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("a line too long")
			return "", nil
		}

		//输出数据：
		fmt.Println("二进制:", line)
		//没办法，Go天生支持UTF8，就是这么任性
		fmt.Println("字符串:", string(line))
		str += string(line)
		str += "\n"
	}
	fmt.Println("=======================")
	fmt.Println(str)
	return str, nil
}

func writeFile(value string) {
	writeValue(value, "output.txt")
}

func writeValue(value string, filename string) error {

	fmt.Println("writeValue中的value:", value)
	//没存储一行，文件都需要打开，关闭很耗内存的
	//每次创建都会清空内容的
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("failed to create the file")
		return err
	}
	defer file.Close()
	file.WriteString(value)
	return nil
}
