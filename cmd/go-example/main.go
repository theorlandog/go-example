package main

import (
    "fmt"
    "github.com/theorlandog/go-example/pkg/config"
    "github.com/theorlandog/go-example/pkg/utils"
)

// this is a comment

func main() {
    fmt.Println("Hello World")
    // use an imported constant
    fmt.Println("VERSION:", config.VERSION)
    // use an imported function
    fmt.Println("Answer is", utils.AddNumbers(2, 3))
}