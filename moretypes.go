package main

import (
	"fmt"
	"time"
)

func ending() {
	fmt.Println("更多类型章节学习结束...", time.Now().String())
}

//指针问题
func pointer()  {
	i,j :=10,20

	p :=&i;
	//引用了i的地址:类似内存地址0xc0000b4038
	fmt.Println(p);
	//* 操作符表示指针指向的底层值：10
	fmt.Println(*p);
	//10
	fmt.Println(i);

	p =&j;
	*p = *p / 2;
	//j指向的是*p
	fmt.Println(j);
	fmt.Println(*p);
}

//定义一个用户信息对象
type memberInfo struct {
	id int
	nickname string
	avatar string
}

func main() {
	fmt.Println("更多类型学习...", time.Now().String())
	defer ending()
	//指针
	pointer();
	//一个结构体（struct）就是一组字段（field）
	memberData :=memberInfo{1,"Nikita","https://xxxx.png"}
	fmt.Println("用户信息",memberData)
	//结构体字段使用点号来访问,类似一个js对象的变量引用
	fmt.Println("用户ID为：",memberData.id)
	fmt.Println("用户昵称为：",memberData.nickname)

	memberOneData :=memberInfo{2,"Pony","https://xxxx.png"}
	memberOneData.nickname = "Jerry"
	memberTwoData :=&memberOneData
	memberTwoData.nickname = "Harry"
	fmt.Println("用户信息",memberOneData)

	//结构体文法通过直接列出字段的值来新分配一个结构体。
	v1 :=memberInfo{3,"Peter","https://xxxx.png"}
	v2 :=memberInfo{id:4}
	v3 :=memberInfo{}
	p1 :=&memberInfo{5,"Lily","https://xxxx"}
	*p1 =memberInfo{id: 6}
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(p1)
	fmt.Println(*p1)

	//数组：类型 [n]T 表示拥有 n 个 T 类型的值的数组。
	var arr [2] string
	arr[0] = "Nikita"
	arr[1] = "awesome"
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[1])

	listData := [4] int {1,2,3,4}
	fmt.Println(listData)
	//数组长度
	fmt.Println(len(listData))

	//切片：类型 []T 表示一个元素类型为 T 的切片。
	numberList :=[6] int {1,2,3,4,5,6}
	fmt.Println(numberList)

	//从numberList中取出3，4，5组成新的数组
	var newNumberList []int = numberList[2:5]
	fmt.Println(newNumberList)
	//复制数组KEY
	newNumberList[0] = 100
	fmt.Println(newNumberList)

	//定义一个数组
	//这是一个数组文法
	letterArr :=[] int{1,2,3,4,5,6}
	fmt.Println(letterArr)

	//定义个书本array-json对象,一个数组里面有N个对象，每个对象包含书本名称name、书本价格prize
	bookData :=[] struct{
		name string
		price float32
	}{
		{"书本1",10.99},
		{"书本2",0.2},
	}
	fmt.Println(bookData)

	//切片的默认行为：切片下界的默认值为 0，上界则是该切片的长度。
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	//下界默认为0
	s = s[:2]
	fmt.Println(s)

	//上界肯定是最大长度
	s = s[1:]
	fmt.Println(s)
}
