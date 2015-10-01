package main

import (
    "log"
    "net/http"
)

func StartRestAPI() {
    http.HandleFunc("/getinfo", GetInfo)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
