package main

import "fmt"

func main() {

	tm := make(map[int]int)
	tm[1] = 1
	tm[2] = 2
	tm[3] = 3
	tm[4] = 4
	tm[5] = 5
	for _, item := range tm {
		fmt.Println(item)
	}
}
