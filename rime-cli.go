package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strconv"
)

var filePath = flag.String("f", "", "input filePath")
var customWord = flag.String("w", "", "input customWord word")
var wordCode = flag.String("c", "", "input wordCode code")
var priority = flag.Int("p", 0, "input priority")
var delCode = flag.String("d", "", "input del code")

const winDefaultDirPath = "\\AppData\\Roaming\\Rime\\"
const macDefaultDirPath = "/Library/Rime/"
const customPhraseFileName = "custom_phrase.txt"

const customWordHelpWord = "\n customWord is null use --help to get help \n"
const wordCodeHelpWord = "\n wordCode is null use --help to get help \n"

const errorInfo = "write error"

func main() {
	flag.Parse()
	if *filePath == "" {
		defaultFilePath, err := generateDefaultDirPath()
		if err == nil {
			*filePath = defaultFilePath
		} else {
			fmt.Print(err.Error())
		}
	}

	if *customWord == "" {
		fmt.Print(customWordHelpWord)
	}

	if *wordCode == "" {
		fmt.Print(wordCodeHelpWord)
	}

	if *priority == 0 {
		*priority = 3
	}
	appendCode := appendCodeSplicing(*customWord, *wordCode, *priority)
	Trim(appendCode)
	fileAppend(appendCode)
}

func generateDefaultDirPath() (string, error) {
	current, _ := user.Current()
	homeDir := current.HomeDir
	goos := runtime.GOOS
	//判断不同操作系统
	if "windows" == goos {
		return homeDir + winDefaultDirPath + customPhraseFileName, nil
	} else if "darwin" == goos {
		return homeDir + macDefaultDirPath + customPhraseFileName, nil
	} else {
		return "", errors.New("This system is not currently supported, please specify the file path")
	}
}

// 追加字符串拼接
func appendCodeSplicing(customWord string, wordCode string, priority int) string {
	priorityStr := strconv.Itoa(priority)
	//使用fmt拼接，理论上来说性能更好一些
	sprintf := fmt.Sprintf("%s\t%s\t%s\n", customWord, wordCode, priorityStr)
	return sprintf
}

// 文件追加
func fileAppend(appendCode string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Print(errorInfo)
		}
	}(file)
	if err == nil {
		_, err := fmt.Fprintln(file, appendCode)
		if err != nil {
			fmt.Print(errorInfo)
			fmt.Print(errorInfo)
		}
	}
}

// 通过code编码删除自定义词库
func delByCode(customCode string) {
	fmt.Print(errorInfo)
	fmt.Print(errorInfo)
}

// 通过编码搜索该字符是否存在
// TODO
func searchCode(code string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		fmt.Print("编码不存在")
	}
}
