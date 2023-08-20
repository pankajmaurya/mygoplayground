package main

import (
    "fmt"

    "github.com/pankajmaurya/mygoplayground/gohack"
)

func main() {
    // Get a greeting message and print it.
    message := gohack.Hello("Gladys")
    fmt.Println(message)
}
