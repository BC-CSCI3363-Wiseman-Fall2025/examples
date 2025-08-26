package main

import "fmt"

func ptrTest(x *int) {
    *x++
    fmt.Println("At end of function:", *x)
}

func swap(x, y *int) {
    *x, *y = *y, *x
}

func main() {
    x := 1
    ptrTest(&x)
    fmt.Println("After function:", x)

    y := 3
    swap(&x, &y)
    fmt.Println(x, y)
}
