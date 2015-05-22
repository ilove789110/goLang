package stack

// import (
// 	"errors"
// )

type Stack []interface{}

func (stack Stack) Len() int { /*方法和函数的区别关键字func和方法名之之间用（）标明作用的类型和值命名；值命名在GoLang中术语叫"接收器"*/
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty bool {
	if stack.Len()==0 {
		return false
	}

	return true
}