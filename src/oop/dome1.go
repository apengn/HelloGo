package main

import "fmt"

type Animal struct {
	Name string
}

type Dog struct {
	Animal
}

type Sleep interface {
	Sleep()
}

func (a *Animal) Eat() {

}
func (a *Animal) Sleep() {
	fmt.Println(a.Name, "sleep")
}
func (a *Animal) Run() {
	fmt.Println(a.Name, "running")
}

func main() {

	//var sleep Sleep
	//
	//d:=&Dog{Animal{"cat"}}
	//
	//sleep=d
	//
	//sleep.Sleep()
	d := new(Dog)
	d.Name = "dog"
	a := new(Animal)
	a.Name = "cat"
	s := []Sleep{d, a}

	for _, sl := range s {
		sl.Sleep()
	}

}
