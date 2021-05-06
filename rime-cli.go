package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
)

var filePath = flag.String("f", "", "input filePath")
var customWord = flag.String("w", "", "input customWord word")
var wordCode = flag.String("c", "", "input wordCode code")
var priority = flag.Int("p", 0, "input priority priority")

func main() {
	flag.Parse()
	if *filePath == "" {
		current, _ := user.Current()
		homeDir := current.HomeDir
		*filePath = homeDir + "/Library/Rime/custom_phrase.txt"
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
	priorityStr := strconv.Itoa(*priority)
	appendCode := *customWord + "	" + *wordCode + "	" + priorityStr + "\n"
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err == nil {
		_, err := fmt.Fprintln(file, appendCode)
		if err != nil {
			fmt.Printf("写入错误")
		}
	}
}
