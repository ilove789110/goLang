package stack

import (
	"errors"
)

type Stack []interface{} /*定义新类型*/

func (stack Stack) Len() int { /*方法和函数的区别关键字func和方法名之之间用（）标明作用的类型和值命名；值命名在GoLang中术语叫"接收器"*/
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	if stack.Len() == 0 {
		return false
	}

	return true
}

/* go方法的参数是引用传参，相当于复制一个新的。
在方法内修改，不影响原值。如果需要修改原值，需要使用指针类型例如 *Stack */
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x) /* stack保存的是指向Stack的指针值，这里的*是解引用的操作。通过解引用获取指针所指向的实际的Stack值 */
}

func (stack Stack) Top() (interface{}, error) {
	if stack.Len() == 0 {
		return nil, errors.New("can't Top en empty stack")
	}

	return stack[stack.Len()-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("can't pop en empty stack")
	}

	x := theStack[len(theStack)-1]

	*stack = theStack[:len(theStack)-1] /*theStack[first:end] 省略first,默认0；省略end,默认len().新获取的切片包含原切片中从first到end之间的所有元素，包含first元素，不包含end元素*/

	return x, nil
}
