package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // create a scanner on os.Stdin (the keyboard)
    // os.Stdin is a built-in variable in the os module
    scanner := bufio.NewScanner(os.Stdin)

    n := 4
    fmt.Printf("Enter %d names, one per line:\n", n)

    for _ = range n {
        // scan for a line
        scanner.Scan()

        // check if there were any errors
        // nil would mean there were no errors
        if err := scanner.Err(); err != nil {
            fmt.Print("Error! ")
            fmt.Println(err)

            // exit the program if there were any errors
            os.Exit(0)
        }

        // get one line of input
        line := scanner.Text()

        // print the line's length
        fmt.Println(len(line))
    }

    return
}
