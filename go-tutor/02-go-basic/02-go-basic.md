## Go Basic
Golang 在學習時, 主要掌握
1. error handle
2. struct 
3. composite object orient
4. json處理
就可以以一個比較容易的方式快速上手
但以下會從基礎開始介紹



## Golang一些常用特性
1. 類型推導
2. 變數及函式名稱開頭大寫表示public, 小寫private


## 標準輸出print
```go
// 印一行
fmt.Println("hello world")

// format print %s是指字串 \n是換行
a := "hello"
b := "world"
fmt.Printf("%s %s\n", a, b)
```

## 資料型別
基本資料型態宣告
```go
// 自動判斷類型
var a = "a" // 一般變數可改值
const b = "b" // 常數變數不可改值

c := "b" // 方便的語法
fmt.Printf("type of %s is: %T",c, c) // %s是指字串 %T是指類型

d := 1
fmt.Printf("type of %s is: %T",d, d)
```


## 函式宣告及呼叫
```go
// 一般函式
func PrintHello(){
  fmt.Println("Hello")
}
PrintHello()


// 帶有參數的函式
func PrintHelloTo(message string) {
  fmt.Println("Hello "+ message)
}
PrintParamTo("Yale")

// 有參數及返回值的函式(盡量用這種, 寫純函式是好習慣)
func GetHelloString(message string) string {
  return "Hello " + message
}
fmt.Println(GetHelloTo("Yale"))

```

## if 及 switch
```go
// 特別注意else的位置, 不然會報錯
flag := 1
if flag==1 {
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

```

## for 和 iterator
```go
// for迴圈
for i:=0;i<100;i++{
  fmt.Println(i)
}

// iterator
people := []string{"yale1", "yale2", "yale3"}
for index,item := range people {
  fmt.Printf("%d: %s", index, people[index])
} 
```

專案  [Github Source Code](https://github.com/cbot918/go-tutor/tree/main/02-go-basic)
