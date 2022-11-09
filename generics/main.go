package main

import "fmt"
// best way!!!
type compareTypes interface {
    ~float64 | int
}

func compare[T compareTypes] (a T, b T) T {
    if a > b {
        return a
    } else {
        return b
    }
}
// -----------------------------------------------------------------------

// simple fast way with two generics
func simple[T any, U any](c T, d U) (T,U) {
    return c,d
}
// -----------------------------------------------------------------------

func main () {
    fmt.Println("hello world")
    fmt.Println(compare(2 ,4))
    fmt.Println(simple("hey" ,10))
}
