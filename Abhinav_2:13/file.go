/* //  this is a comment

package main
import ("fmt")

func main() {
  variable2 := 20 // 2nd Method, and only works inside a function
  fmt.Println("Hello World!", variable2, variable1)
  /* this is a comment

  var a, b = 6, "Hello"
  c, d := 7, "World"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)

  fmt.Println(PI)
}

// ThisIsPascalCase
// thisIsCamelCase
// this_is_snake_case

var variable1 int = 10 // 1st Method

const constant1 int = 30 // This is a constant

const PI = 3.14
*/

package main

import (
	"fmt"
)

func main() {

	var i, j = 10, 20
	fmt.Print(i, j, "\n")

	var a, b string = "Hello", "World"
	fmt.Println(a, b)

	var k string = "Hello"
	var l int = 15
	fmt.Printf("k has value: %v and type: %T\n", k, k)
	fmt.Printf("l has value: %v and type: %T\n", l, l)

}
