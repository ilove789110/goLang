package main

import (
	"fmt"
	"math"
)

const (
	factor = 3
)

func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}

	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}

func main() {
	i := 2000
	i *= factor
	j := int16(20)
	i += int(j)
	k := uint8(0)
	k, e := Uint8FromInt(i)

	fmt.Printf("%T", factor)
	fmt.Printf("%v", factor)
	fmt.Println()

	fmt.Printf("%T", i)
	fmt.Printf("%v", i)
	fmt.Println()

	if e != nil {
		fmt.Println(i, j, k)
	}

}
