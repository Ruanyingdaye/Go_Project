package main

import (
	"fmt"
	"strconv"
)

func main() {
	sss := "abcdefg"
	for i := 0; i < len(sss); i++ {
		fmt.Println(sss[i])
	}

	fmt.Println("******************************************************")

	//结论 遍历string得到的是rune

	for _, item := range sss {
		fmt.Println(item)
		// 这就是说明，遍历过来的是rune数据，需要进行转换才可以
		// 输出%c，表示可以以字符的方式输出
		rv, ok := strconv.ParseInt(string(item), 10, 64)
		if ok == nil {
			fmt.Println(rv)
		}
		fmt.Printf("%c", item)

		//要整理的话可以用
		//1. bytes[]，整合到一起再变成string
		//2. 可以直接用strings.Split转化
	}

	fmt.Println("******************************************************")

	byteSSS := []byte(sss)
	for _, item := range byteSSS {
		fmt.Println(item)
	}

	byteSSS[3] = 125
	news := string(byteSSS)
	fmt.Println(news)
}
