package main

import (
	"fmt"
)

type Dog struct {
	name	string
}

func (d Dog) Speak() string {
	  return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
    return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
    return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}

type Animal interface {
    Speak() string
}

func main() {
    animals := []Animal{Dog{"dog"}, Cat{}, Llama{}, JavaProgrammer{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
