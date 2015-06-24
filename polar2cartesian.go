package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"runtime"
// )

// type polar struct {
// 	radius float64 /*半径*/
// 	θ      float64 /*极坐标的角度*/
// }

// type cartesian struct { /*struct 结构体关键字*/
// 	x float64
// 	y float64
// }

// var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90" + "or %s to quit."

// func init() { /*一个包里包含一个或多个init()函数，那么它们会在main()函数之前被自动执行，而且init（）函数不能被显式调用*/
// 	if runtime.GOOS == "windows" { /*runtime.GOOS常用值darwin(Mac OS X)、freebsd、linux及windows*/
// 		prompt = fmt.Sprintf(prompt, "Ctrl+Z,Enter")
// 	} else {
// 		prompt = fmt.Sprintf(prompt, "Ctrl+D")
// 	}
// }

// func main() {
// 	questions := make(chan polar) /*make(chan type) 定义通道*/
// 	defer close(questions)        /*close(chan) 关闭通道*/
// 	answers := createSolver(questions)
// 	defer close(answers)
// 	interact(questions, answers)
// }

// func createSolver(questions chan polar) chan cartesian {
// 	answers := make(chan cartesian)
// 	go func() { /*go 创建一个独立的异步goroutine来执行函数，当前函数的控制流程会继续向下执行*/
// 		/*匿名函数，该函数有一个无限循环体处于阻塞等待状态，直到它接受到一个问题。但不会阻塞其他goroutine,也不会阻塞创建该goroutin的函数*/
// 		for {
// 			polarCoord := <-questions /*通信操作符<-作为一元操作符获取通道里的元素，通道中的元素先入先出.它作为一个接受器，一直阻塞直到获取一个可以返回的数据*/
// 			θ := polarCoord.θ * math.Pi / 180.0
// 			x := polarCoord.radius * math.Cos(θ)
// 			y := polarCoord.radius * math.Sin(θ)
// 			answers <- cartesian{x, y} /*通信操作符<-作为二元操作符，它的左操作数必须是一个通道，右操作数必须是发往该通道的数据，其类型为通道声明时所能接收的类型*/
// 		}

// 	}()
// 	return answers
// }

// const result = "Polar radius=%.02f θ=%.02f° →Cartesian x=%.02f y=%.02f\n"

// func interact(questions chan polar, answers chan cartesian) {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println(prompt)
// 	for {
// 		fmt.Println("Radius and angle: ")
// 		line, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			break
// 		}
// 		var radius, θ float64
// 		if _, err := fmt.Sscan(line, "%f %f", &radius, &θ); err != nil {
// 			fmt.Println(os.Stderr, "invalid input")
// 			continue
// 		}
// 		questions <- polar{radius, θ} /*createSolver的goruntine开始执行，并将结果cartesian发送会answers通道*/
// 		coord := <-answers            /*一直等到answers通道有可传递的cartsian前一直阻塞*/
// 		fmt.Println(result, radius, θ, coord.x, coord.y)
// 	}
// 	fmt.Println()
// }
