package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "networks"

    fmt.Println(s)

    for i := range s {
        fmt.Printf("%d: %c\n", i, s[i])
    }

    for i,char := range s {
        fmt.Printf("%d: %c\n", i, char)
    }

    for _,char := range strings.ToUpper(s) {
        fmt.Println(string(char))
    }

    // this would be a compile error:
    //s[0] = "y"

    return
}
