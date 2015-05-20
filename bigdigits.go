package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 333 ", "3   3", "    3", "   33", "    3", "3   3", " 333 "},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			/*Go中字符串编码为UTF-8，字符'0'对应的是48，字符'1'对应的是49。
			一个字符是与Go的整型类型兼容的整型数，而且Go语言的数值类型会适应上下文，跟Go语言强类型且不支持隐式类型不冲突*/
			if 0 <= digit && digit <= 9 { /*此处0和9被认定为byte类型*/
				line += bigDigits[digit][row] + " " //Go中string类型是不可变的，但是+=和+是支持的，主要是易于使用。实质上是暗地里将字符串替换了
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}
