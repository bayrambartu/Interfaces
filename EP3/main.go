// Type Assertion- Type switch
package main

import "fmt"

var i interface{} = 13

func main() {
	// s, ok := i.(string)
	// if ok {
	// 	fmt.Println("String degeri :", s)
	// } else {
	// 	fmt.Println("String degil")
	// }

	switch v := i.(type) {
	case string:
		fmt.Println("String degeri..:", v)
	case int:
		fmt.Println("Integer degeri..:", v)
	default:
		fmt.Println("Bilinmeyen type..")
	}

}
