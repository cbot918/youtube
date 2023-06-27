package main

import "fmt"

// struct
type Student struct {
	Name         string
	Number       string
	FavoriteFood string
}

// object oriented
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

	// struct
	s1 := &Student{
		Name:         "Aaron",
		Number:       "12345",
		FavoriteFood: "chocolate",
	}
	fmt.Println(s1.Name)         // Aaron
	fmt.Println(s1.Number)       // 12345
	fmt.Println(s1.FavoriteFood) // chocolate

	// object oriented
	p1 := NewPerson("Aaron", "France")
	p1.PrintPerson()
	p1.SetPerson("Aaron_edit", "France_edit")
	p1.PrintPerson()

	p2 := NewPerson("Bartner", "Australia")
	p2.PrintPerson()
	p2.SetPerson("Bartner_edit", "Australia_edit")
	p2.PrintPerson()

}
