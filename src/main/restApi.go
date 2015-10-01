package main

import (
    "log"
    "net/http"
)

func StartRestAPI() {
    http.HandleFunc("/getinfo", RestGetInfo)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
