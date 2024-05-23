package main

import "strconv"

type Person struct {
	name string
	age  int
}

func (p Person) getName() string {
	return "My Name is " + p.name + " and my age is " + strconv.Itoa(p.age) + "!"
}

func strct() {
	var person = Person{name: "John", age: 20}
	println(person.getName())
}
