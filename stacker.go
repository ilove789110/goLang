package main

import (
	"fmt"
	"goLang/stacker/stack"
)

func main() {
	var hayStack stack.Stack /*type Stack []interface{}*/
	hayStack.Push("hay")
	hayStack.Push(-15)
	hayStack.Push([]string{"pin", "clip", "needle"})
	hayStack.Push(81.52)

	for {

		fmt.Println(hayStack.Len())
		item, err := hayStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
	fmt.Println(hayStack)

	fmt.Println(hayStack.Len())

}
