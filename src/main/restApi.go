package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/getinfo", GetInfo)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
