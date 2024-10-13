// empty Interface

package main

import "fmt"

var i interface{}

func main() {
	i = 42
	fmt.Println(i) // 42

	i = "Merhaba"
	fmt.Println(i) // Merhaba

	i = true
	fmt.Println(i) // trure
}
