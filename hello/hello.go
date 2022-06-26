package main

import (
    "fmt"

    "example/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Kale")
    fmt.Println(message)
}