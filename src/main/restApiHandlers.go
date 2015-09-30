package main

import (
    "encoding/json"
    "net/http"
    "time"
)

type SysInfo struct {
    Temp    int
    Time    SysTime
}

type SysTime struct {
    Day     int
    Month   int
    Year    int
    Hour    int
    Min     int
    Sec     int
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
    timestamp := time.Now()
    time := SysTime{
        Day: time.Time.Day(timestamp),
        Month: int(time.Time.Month(timestamp)),
        Year: time.Time.Year(timestamp),
        Hour: time.Time.Hour(timestamp),
        Min: time.Time.Minute(timestamp),
        Sec: time.Time.Second(timestamp),
    }
    info := SysInfo {
        Temp: 20,
        Time: time,
    }

    json.NewEncoder(w).Encode(info)
}
