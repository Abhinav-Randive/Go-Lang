package main

import "fmt"

//func main() {

    // for loop

    /* for i := 0; i < 5; i++ {
        if i == 3 {
            // continue
            break
        }
        fmt.Println(i)
    } */

    //nested loops
    /* adj := [2]string{ "big", "tasty"}
    fruits := [3]string{"apple", "banana", "orange"}
    for i := 0; i < len(adj); i++ {
        for j := 0; j < len(fruits); j++ {
            fmt.Println(adj[i], fruits[j])
        }
    } */
    
   /*  fruits := [3]string{"apple", "banana", "orange"}
    for _, val := range fruits {
        fmt.Printf( "%v\n", val)
    } */
    
    //--------------------------------------------------------------------------------

    /* number := 10

    if number > 0 {
        fmt.Println("Number is positive")
    } else if number < 0 {
        fmt.Println("Number is negative")
    } else {
        fmt.Println("Number is zero")
    }

    //--------------------------------------------------------------------------------
    if number%2 == 0 {
        fmt.Println("Number is even")
    } else {
        fmt.Println("Number is odd")
    } 

    //--------------------------------------------------------------------------------
    num1 := 5
    num2 := 10

    if num1 > num2 {
        fmt.Println("Num1 is greater than Num2")
    } else if num2 > num1 {
        fmt.Println("Num2 is greater than Num1")
    } else {
        fmt.Println("Both numbers are equal")
    } 

    //--------------------------------------------------------------------------------
    fruits := [3]string{"apple", "banana", "orange"}
    for _, val := range fruits {
        fmt.Printf( "%v\n", val)
    }

    //--------------------------------------------------------------------------------
    a := 20
    b := 4
    fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
    fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
    fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
    fmt.Printf("Division: %d / %d = %d\n", a, b, a/b) */

    //--------------------------------------------------------------------------------
//}

/* func myMessage() {
        fmt.Println("I just got executed!")
}

func familyName (fname string, age int ) {
        fmt.Println("Hello", age, "years old", fname, "Refsnes")
}

func main() {
        familyName("Liam", 3)
        familyName("Jenny", 14)
        familyName("Anja", 30)

        myMessage()
        myMessage()
} */

func  myFunction(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(myFunction(5, 6))
}
