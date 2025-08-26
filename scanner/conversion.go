package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    n := 4
    fmt.Printf("Enter %d numbers, one per line:\n", n)

    total := 0
    for _ = range n {
        scanner.Scan()
        if err := scanner.Err(); err != nil {
            fmt.Print("Error! ")
            fmt.Println(err)
            os.Exit(0)
        }

        line := scanner.Text()

        num, err := strconv.Atoi(line)
        if err != nil {
            fmt.Print("Error! ")
            fmt.Println(err)
            os.Exit(0)
        }

        total += num
    }

    fmt.Printf("Total=%d\n", total)

    return
}
