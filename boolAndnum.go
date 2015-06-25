package main

import (
	"fmt"
)

const (
	Cyan    = 1
	Magenta = 2
	Yellow  = 0
	Black   = iota /*iota的特性：遇到const关键字重置为0；在const的第一行出现，iota=0；在第二行出现，iota=1；依此类推*/
	White
	Red
)

type BitFlag int

const (
	Active  BitFlag = 1 << iota
	Send            //隐式设置成BitFlag=1<<iota 1<<1=2
	Receive         //隐式设置成BitFlag=1<<iota 1<<2=2
)

func main() {
	fmt.Println(Cyan)
	fmt.Println(Magenta)
	fmt.Println(Yellow)
	fmt.Println(Black)
	fmt.Println(White)
	fmt.Println(Red)

	fmt.Println(Active)
	fmt.Println(Send)
	fmt.Println(Receive)

	flag := Active | Receive
	fmt.Println(flag)
}
