package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Weight float64
	Gender string
}

type Cat struct {
	Name   string
	Age    int
	Weight float64
	Gender string
}

func (p Person) Speak() string {
	return fmt.Sprintf("%s says Hello", p.Name)
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s says Miow", c.Name)
}

type Family interface {
	Speak() string
}


func main() {
	per := Person{Name: "Kri"}
	cat := Cat{Name: "Julie"}
	fam := []Family{per, cat}
	fmt.Println(fam[0].Speak())
	fmt.Println(fam[1].Speak())

}
