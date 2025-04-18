package main

import (
	"bufio"
	"fmt"
	"os"
)

func exeFileA(str string) {
	if file, err := os.Open(str); err == nil {
		defer file.Close()
		newScan := bufio.NewScanner(file)
		for newScan.Scan() {
			item := newScan.Text()
			fmt.Println(item)
		}
	}
}

func main() {
	file, err := os.ReadDir("aaa")
	if err == nil {
		fmt.Println(file)
	}
}
