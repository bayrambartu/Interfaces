// Generics and interface,
// package main

// import "fmt"

// func Max[T interface{ int | float64 }](a, b T) T {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
// func main() {
// 	fmt.Println(Max(5, 10))      // 10
// 	fmt.Println(Max(3.14, 2.71)) // 3.14
// }

// Reflection and interfaces
package main

import (
	"fmt"
	"reflect"
)

func PrintTypeAndValue(x interface{}) {
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	fmt.Printf("tipi : %s, Deger : %v\n", t, v)
}

func main() {
	PrintTypeAndValue("Merhaba!")
	PrintTypeAndValue(342)
	PrintTypeAndValue(3.14)

}
