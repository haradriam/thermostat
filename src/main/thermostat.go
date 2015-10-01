package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello world")

    go StartRestAPI()
    time.Sleep(10 * time.Second)
}
