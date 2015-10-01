package main

import sc "dht11Controller"

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello world")

    go StartRestAPI()

    a := sc.ReadValues()
    fmt.Println("Hello world ", a.Temp)
    time.Sleep(10 * time.Second)
}
