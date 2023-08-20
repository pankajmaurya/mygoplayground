package main

import (
    "fmt"

    "github.com/gohack"
)

func main() {
    // Get a greeting message and print it.
    message := gohack.Hello("Gladys")
    fmt.Println(message)
}
