package main

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "strings"
)

func RestGetInfo(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(GetSysInfo())
}

func RestGetEvents(w http.ResponseWriter, r *http.Request) {
    eventList, err := DbReadEvents()
    checkErr(err)

    json.NewEncoder(w).Encode(eventList)
}

func RestSetEvents(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    checkErr(err)

    dec := json.NewDecoder(strings.NewReader(string(body)))
    dec.Token()

    var eventList []EventEntry
    var event EventEntry

    for dec.More() {
        dec.Decode(&event)

        eventList = append(eventList, event)
    }

    DbAddEvents(eventList)
}
