package main

import (
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

/**
* 字符串判断为空
 */
func isEmpty(str string) bool {
	return len(str) == 0
}

/**
* 字符串判断非空
 */
func isNotEmpty(str string) bool {
	return !isEmpty(str)
}
