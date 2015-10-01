package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello world")

    go StartRestAPI()

    for {
        err := AddHist(GetInfo())
        checkErr(err)

        time.Sleep(2 * time.Second)
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
