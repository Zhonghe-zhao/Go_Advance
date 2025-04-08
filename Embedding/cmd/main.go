package main

import "fmt"

type base struct {
	num int
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
			str: "base",
		},
		str: "contianer",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) //直接访问嵌套结构体的字段

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe()) //方法继承：若类型base有方法那么 co也会继承这个方法
	// 如果外层结构体和嵌入类型有同名字段或方法，外层会 遮蔽（Shadow） 内层：
	fmt.Println(":", co.str)
	//显示调用内层方法
	fmt.Println(":", co.base.str)

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
