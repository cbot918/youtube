## Intro
Golang的物件導向是使用「組合」, 具有輕巧, 易重複使用, 寫起來也比較方便的特性

## Struct
struct是golang中很常使用到的資料結構
以下為語法示範
```go

type Student struct {
  Name string
  Number string
  FavoriteFood string
}

s1 := &Student {
  Name: "Aaron",
  Number: "12345",
  FavoriteFood: "chocolate",
}

fmt.Println(s1.Name) // Aaron
fmt.Println(s1.Number) // 12345
fmt.Println(s1.FavoriteFood) // chocolate

```

## Object Oriented
- 此範例中, 我們先使用創造了Person struct
- 接著使用一個 NewPerson的 constructor來建立實體
- 然後使用 func (p *Person) 的語法, 將 PrintPerson() 以及 SetPerson()這兩個函式關聯到Person

目前為止
- Person的data member: Name, From
- Person的constructor: NewPerson()
- Person的function member: PrintPerson(), SetPerson()

以下薇範例程式碼
```go
package main

import "fmt"

type Person struct {
	Name string
	From string
}

func NewPerson(name string, from string) *Person {
	return &Person{
		Name: name,
		From: from,
	}
}

func (p *Person) PrintPerson() {
	fmt.Println("name: ", p.Name)
	fmt.Println("from: ", p.From)
}

func (p *Person) SetPerson(name string, from string) {
	p.Name = name
	p.From = from
}

func main() {

  // person1
	p1 := NewPerson("Aaron", "France")

	p1.PrintPerson()
	p1.SetPerson("Aaron_edited", "France_edited")
	p1.PrintPerson()

  // person2
	p2 := NewPerson("Bartner", "Australia")

	p2.PrintPerson()
	p2.SetPerson("Bartner_edited", "Australia_edited")
	p2.PrintPerson()

}

```