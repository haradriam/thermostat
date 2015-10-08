package main

import (
    "log"
    "net/http"
)

func StartRestAPI() {
    http.HandleFunc("/getinfo", RestGetInfo)
    http.HandleFunc("/getconfig", RestGetConfig)
    http.HandleFunc("/setconfig", RestSetConfig)
    http.HandleFunc("/getevents", RestGetEvents)
    http.HandleFunc("/setevents", RestSetEvents)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
