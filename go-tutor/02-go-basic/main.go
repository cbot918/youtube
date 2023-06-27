package main

import (
	"fmt"
)

func Print() {
	fmt.Println("hello world")

	// format print %s是指字串 \n是換行
	a := "hello"
	b := "world"
	fmt.Printf("%s %s\n", a, b)
}

func DataType() {
	// 自動判斷類型
	// var a = "a"   // 一般變數可改值
	// const b = "b" // 常數變數不可改值

	c := "b"                                 // 方便的語法
	fmt.Printf("type of %s is: %T \n", c, c) // %s是指字串 %T是指類型

	d := 1
	fmt.Printf("type of %d is: %T \n", d, d)
}

// 一般函式
func PrintHello() {
	fmt.Println("Hello")
}

// 帶有參數的函式
func PrintHelloTo(message string) {
	fmt.Println("Hello " + message)
}

// 有參數及返回值的函式(盡量用這種, 寫純函式是好習慣)
func GetHelloTo(message string) string {
	return "Hello " + message
}

// if 及 switch
func IfAndSwitch() {
	// 特別注意else的位置, 不然會報錯
	flag := 1
	if flag == 1 {
		fmt.Println("true condition")
	} else {
		fmt.Println("false condition")
	}

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("eat apple")
	case "banana":
		fmt.Println("eat banana")
	default:
		fmt.Println("not hungry")
	}
}

func ForAndIterator() {
	// for迴圈
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// iterator
	people := []string{"yale1", "yale2", "yale3"}
	for index, item := range people {
		fmt.Printf("%d: %s: %s \n", index, people[index], item)
	}

}

func main() {
	// print
	Print()

	// datatype
	DataType()

	// Functions
	PrintHello()
	PrintHelloTo("Y")
	fmt.Println(GetHelloTo("Y"))

	// if and switch
	IfAndSwitch()

	// for and iterator
	ForAndIterator()
}
