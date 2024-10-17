// Basic example application : Zoo
package main

import "fmt"

type Animal interface {
	MakeSound()
}
type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("Meow")
}

type Dog struct {
}

func (d Dog) MakeSound() {
	fmt.Println("Woof")
}

func AnimalSounds(animals []Animal) {
	for _, animal := range animals {
		animal.MakeSound()
	}
}

func main() {
	animals := []Animal{Cat{}, Dog{}}
	AnimalSounds(animals)
}
