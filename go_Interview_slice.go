package main

import (
	"fmt"
	"strings"
)

//reverse string

// func reverseString(s string) {
// 	// i, j := 0, len(s)-1
// 	// for i < j {
// 	// 	// s[i], s[j] = s[j], s[i]
// 	// 	s[i] = 'a'
// 	// }

// 	//golang的某个字符串是不能直接修改某个值的
// 	//可以修改为[]byte

// }

// func main() {
// 	//golang的某个字符串是不能直接修改某个值的
// 	//可以修改为[]byte， 但是这么修改会涉及数制转换的问题！！

// 	ss := "adfsfsd"
// 	aa := []byte(ss)
// 	aa[0] = 122
// 	ss = string(aa)
// 	fmt.Println(ss)

/*************************/

func check(s string) bool {
	ss := strings.Split(s, "")
	tmp := []string{}
	for _, item := range ss {
		if item == "(" || item == "{" || item == "[" {
			tmp = append(tmp, item)
		} else {
			if item == ")" {
				if tmp[len(tmp)-1] != "(" {
					return false
				}
			} else if item == "}" {
				if tmp[len(tmp)-1] != "{" {
					return false
				}
			} else if item == "]" {
				if tmp[len(tmp)-1] != "[" {
					return false
				}
			}
			tmp = tmp[:len(tmp)-1]
		}
		fmt.Println("aa", item, "bb", tmp)
	}
	if len(tmp) != 0 {
		return false
	}
	return true
}
func main() {
	ss1 := "()[]{}"
	ss2 := "{()][]}"
	fmt.Println(check(ss1))
	fmt.Println(check(ss2))
}
