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
	}
}
