package main

import "strings"

/*
*
字符串工具类
*/
func StringUtil(str string) string {
	print(str)
	return strings.Trim(str, "")
}
