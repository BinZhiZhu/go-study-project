package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

/**
类型在变量名 之后
*/
func sum(x int, y int) int {
	return x + y
}

/**
当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
*/
func add(x, y int) int {
	return x + y
}

/**
返回两个字符串
*/
func returnMultipleValue(a, b string) (string, string) {
	return a, b
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//声明变量->类型
var q,w,e,r  bool ;

//声明初始化默认值
var aname,bname string = "aname","bname"

//分组导入变量
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//注意：常量不能用 := 语法声明

func main() {
	//根据当前时间生成种子数咯
	rand.Seed(time.Now().Unix())
	//随机字符串
	fmt.Println("随机数", rand.Intn(10))
	fmt.Println("π", math.Pi)

	//两数相加
	fmt.Println("3+4 = ", sum(3, 4))
	//参数类型相同，可缩写入参类型定义
	fmt.Println("3+4 = ", add(3, 4))

	//函数多值返回: 类似php的数组array['a'=>'','b'=>'']
	 a, b := returnMultipleValue("hello", "world")
	fmt.Println(a, b);

	fmt.Println(split(17))

	fmt.Println(q,w,e,r)

	fmt.Println(aname,bname)

	//短变量声明
	c,python,java,php := true, true, false,"a"

	fmt.Println(c,python,java,php)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
