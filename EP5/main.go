// Nil Interfaces
package main

import "fmt"

type Shape interface {
}
type Dikdörtgen struct {
	height float64
	weidth float64
}

func main() {

	// var s Shape
	// fmt.Printf("%T,%v\n", s, s)

	// if s == nil {

	// 	fmt.Println("s nil'dir")
	// }
	var s Shape = (*Dikdörtgen)(nil)
	fmt.Printf("%T, %v\n", s, s) // *main.Dikdörtgen, <nil>
	if s == nil {
		fmt.Println("s nil'dir")
	} else {
		fmt.Println("s nil değildir")
	}
}
