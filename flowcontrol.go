package main

import (
	"fmt"
	"runtime"
	"time"
)

func addSum(x, y, z int) int {
	//该语句声明的变量作用域仅在 if 之内。
	//这里注意v只能作用域if的流程语句内
	if v := x + y; v < z {
		//将x+y的和赋值给v，v如果小于z则直接返回v的值
		return v
	} else {
		fmt.Println("v是大于等于z的", v, z)
	}
	return z
}

func end() {
	fmt.Println("循环以及条件分支学习结束", time.Now().String())
}

//流程控制
func main() {
	fmt.Println("循环以及条件分支学习开始", time.Now().String())

	//defer语句
	//defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	defer end()

	//迭代累加求和 1+10
	sum := 0
	for i := 1; i < 11; i++ {
		sum += i
	}
	fmt.Println("1+2+...+10 = ", sum)

	sum1 := 1
	for sum1 < 10 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	//for 是 Go 中的 “while”
	sum2 := 15
	sum3 := 0

	for sum2 > 10 {
		sum3++
		sum2--
	}

	fmt.Println(sum3)

	//无限循环
	//for {
	//	fmt.Println("looping...")
	//}

	//if语句
	a, b, c := 1, 2, 3

	if a < b {
		fmt.Println(a, "小于", b)
	}

	if b < c {
		fmt.Println(b, "小于", c)
	}

	//if 的简短语句
	fmt.Println(
		addSum(3, 4, 10),
		addSum(3, 7, 10),
	)

	//switch分支
	fmt.Println("Go runs on", runtime.GOOS)

	//跟php类似，case的类型可以是字符串，不需要手动break，命中会自动终止
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux ")
	default:
		fmt.Printf("%s.\n", os)
	}

	ref := "nikita"

	switch ref {
	case "nikita":
		fmt.Println("I am", ref)
	default:
		fmt.Println("I am not ", ref)
	}

	//没有条件的 switch 同 switch true 一样。这种形式能将一长串 if-then-else 写得更加清晰。
	switch {
	case a < b:
		fmt.Println("a小于b", a, b)
	}

	//defer 栈
	//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}
