package main

import (
	"fmt"
	"os"
	"strings"
)

/*
*
字符串工具类
*/
func Trim(str string) string {
	print(str)
	return strings.Trim(str, "")
}

func updateCode(code string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		fmt.Print("更新成功")
	}
}

func deleteCode(code string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		fmt.Print("更新成功")
	}
}

func complateCode(code string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		fmt.Print("更新成功")
	}
}

func testCode(code string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		fmt.Print("更新成功")
		// 你好
	}
}

func test5Code(code string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		fmt.Print("更新成功")
		// 你好
	}
}

func test6Code(code string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		fmt.Print("更新成功")
		// 你好
	}
}

func print(code string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		fmt.Print("更新成功")
		// 你好
	}
}

func test(str string) string {
	print(str)
	return strings.Trim(str, "")
}

func query(str string) string {
	print(str)
	return strings.Trim(str, "")
}

func test1(str string) string {
	print(str)
	return strings.Trim(str, "")
}

func test2(str string) string {
	print(str)
	return strings.Trim(str, "")
}

func test3(str string) string {
	print(str)
	return strings.Trim(str, "")
}

func test4(str string) string {
	print(str)
	return strings.Trim(str, "")
}

func test7(str string) string {
	print(str)
	return strings.Trim(str, "")
}

func test8(str string) string {
	print(str)
	return strings.Trim(str, "")
}
