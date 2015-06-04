package goLang

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
		item, err := hayStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
