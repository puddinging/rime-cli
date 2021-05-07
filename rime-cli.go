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
var priority = flag.Int("p", 0, "input priority priority")

const winDefaultDirPath = "\\AppData\\Roaming\\Rime\\"
const macDefaultDirPath = "/Library/Rime/"
const customPhraseFileName = "custom_phrase.txt"

func main() {
	flag.Parse()
	if *filePath == "" {
		defaultFilePath, err := generateDefaultDirPath()
		if err == nil {
			*filePath = defaultFilePath
		} else {
			fmt.Printf(err.Error())
		}
	}

	if *customWord == "" {
		fmt.Printf("\n customWord is null use --help to get help \n")
	}

	if *wordCode == "" {
		fmt.Printf("\n wordCode is null use --help to get help \n")
	}

	if *priority == 0 {
		*priority = 1
	}
	appendCode := appendCodeSplicing(*customWord, *wordCode, *priority)
	fileAppend(appendCode)
}

func generateDefaultDirPath() (string, error) {
	current, _ := user.Current()
	homeDir := current.HomeDir
	goos := runtime.GOOS
	if "windows" == goos {
		return homeDir + winDefaultDirPath + customPhraseFileName, nil
	} else if "darwin" == goos {
		return homeDir + macDefaultDirPath + customPhraseFileName, nil
	} else {
		return "", errors.New("This system is not currently supported, please specify the file path")
	}
}

// 代追加字符串拼接
func appendCodeSplicing(customWord string, wordCode string, priority int) string {
	priorityStr := strconv.Itoa(priority)
	return customWord + "	" + wordCode + "	" + priorityStr + "\n"
}

// 文件追加
func fileAppend(appendCode string) {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		_, err := fmt.Fprintln(file, appendCode)
		if err != nil {
			fmt.Printf("写入错误")
		}
	}
}
