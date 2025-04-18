package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// test the bufio.scanner
func main() {
	file, err := os.Open("D:/Project/trace.out")
	if err != nil {
		fmt.Println("open file error")
	}
	defer file.Close()
	findit := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // line is the string
		if strings.Contains(line, "hello") {
			//strings是自带Contains的！
			fmt.Println("Hello World")
			fmt.Println("found it")
			findit = true
			break
		}
	}
	if findit == false {
		fmt.Println("not found")
	}
}
