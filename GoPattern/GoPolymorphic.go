package main

import (
	"fmt"
)

type ani interface {
	Eat()
}

type Animal struct {
	name string
	age  int
}

func NewAnimal(n string, a int) *Animal {
	return &Animal{
		name: n,
		age:  a,
	}
}

func (a *Animal) Sleep() {
	fmt.Println(a.name, " have slept for 1 hour.")
}

func (a *Animal) Eat() {
	fmt.Println(a.name, " is eating sth.")
}

type Dog struct {
	Animal
}

func NewDog(n string, a int) *Dog {
	return &Dog{
		Animal: *NewAnimal(n, a),
	}
}

func (d *Dog) Bark() {
	fmt.Println("bark bark!")
}

func (d *Dog) Eat() {
	d.Animal.Eat()
	fmt.Println("That is bone.")
}

type Cat struct {
	Animal
}

func (c *Cat) Meow() {
	fmt.Println("meow meow!")
}

func (c *Cat) Eat() {
	c.Animal.Eat()
	fmt.Println("That is fish.")
}

func NewCat(n string, a int) *Cat {
	return &Cat{
		Animal: *NewAnimal(n, a),
	}
}

var animal_list = []ani{}

func main() {
	animal_list = append(animal_list, NewCat("sky", 10))
	animal_list = append(animal_list, NewDog("rim", 10))
	for _, val := range animal_list {
		val.Eat()
	}
}
