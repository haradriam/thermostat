package main

import (
    "encoding/json"
    "net/http"
)

func RestGetInfo(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(GetInfo())
}
