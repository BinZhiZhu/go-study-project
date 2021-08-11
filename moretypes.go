package main

import (
	"fmt"
	"time"
	"strings"
)

func ending() {
	fmt.Println("更多类型章节学习结束...", time.Now().String())
}

//昵称
type NicknameObj struct {
	name string
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

func catchError()  {
	if err := recover(); err != nil {
		fmt.Println("捕获到了panic产生的异常：", err) // 这里的err其实就是panic传入的内容
	}
}

func adder() func(int) int  {
	sum :=1
	return func(i int) int {
		fmt.Println("i===",i,sum)
		return i+sum;
	}
}

//Go中引入的Exception处理：defer, panic, recover
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

	//nil 切片
	var strArr []string;
	fmt.Println(strArr,len(strArr),cap(strArr))
	if strArr == nil{
		fmt.Println("strArr is nil")
	}

	//make创建切片
	//创建长度=容量都为5的切片
	emptySlice := make([]int,5)
	fmt.Println(emptySlice)
	//创建长度=2，容量为5的切片
	slice1 :=make([]string,2,5)
	fmt.Println(slice1,len(slice1),cap(slice1))
	slice1[1] = "Nikita"

	// 必须要先声明defer，否则不能捕获到panic异常,也就是说要先注册函数，后面有异常了，才可以调用
	defer catchError()

	//报错，因为容量最多为5，不能超出
	//slice1[5] = "Nikita1"
	//fmt.Println(slice1)

	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)

	//向切片追加元素:func append(s []T, vs ...T) []T
	//空切片
	slice2 :=make([]int,0)
	fmt.Println(slice2)
	slice2 =append(slice2,1,2)
	fmt.Println(slice2)

	//range使用
	nicknames :=[] string{"Nikita","Peter","pony"}
	fmt.Println(nicknames)
	//循环遍历nicknames
	for i,v :=range nicknames{
		fmt.Println(i,v)
	}
	numberData :=make([]int,0);
	numberData = append(numberData, 1,2,3,4,5)
	for i,v :=range numberData{
		fmt.Println(i,v)
	}
	//忽略value，只遍历key
	for key :=range numberData{
		fmt.Println(key)
	}
	//忽略key，只遍历value
	for _,value :=range numberData{
		fmt.Println("value->",value)
	}


	// 映射map: 其实就是类型定义一个数组，数组key=>对象
	var cpNickname map[string]NicknameObj
	cpNickname =make(map[string]NicknameObj)
	fmt.Println(cpNickname)
	cpNickname["name"] = NicknameObj{name: "哈哈"}
	fmt.Println(cpNickname,cpNickname["name"])

	//定义一个映射文法，people对应的信息
	students := map[string]memberInfo{
		"Nikita":memberInfo{
			id: 1,
			nickname: "Nikita",
			avatar: "",
		},
		"Peter":{id: 2,nickname: "Peter",avatar: ""},
	}
	fmt.Println(students)

	//删除学生Peter
	delete(students,"Peter")
	fmt.Println(students)

	//检查peter是否存在
	v,hasPeter := students["Peter"]
	fmt.Println(v,hasPeter,students["Peter"] == memberInfo{})
	c :=strings.Fields("I am learning Go!")
	fmt.Println(c)
	for i,v :=range c{
		fmt.Println(i,v)
	}

	//函数闭包
	pos := adder()

	pos(0)

	var aName []string
	aName = append(aName,"sss")
	strings.Join(aName,"")
	fmt.Println(aName)
	str := "12"
	str += "test-string"
	fmt.Println(str)

	str = strings.Join([]string{str, "-test-xxx"}, "")
	fmt.Println(str)

	//https://go-zh.org/pkg/strings/#Fields
	//抛出异常:Go中引入的Exception处理：defer, panic, recover
	panic("抛出异常xxx")//直接终止程序
	fmt.Println("error")
}
