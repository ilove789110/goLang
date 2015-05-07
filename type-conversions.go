package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)

	var i int = 42
	var j float64 = float64(i)

	fmt.Println(float64(i) + j) //任何类型的隐式类型转换都会造成编译报错

}
