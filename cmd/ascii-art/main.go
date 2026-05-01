package main

import (
    "fmt"
    "os"

    "ascii-art/internal/ascii"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run . \"text\"")
        return
    }

    input := os.Args[1]

    result, err := ascii.Generate(input, "banners/standard.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Print(result)
}
