package main

import (
    "time"
)

type EnvInfo struct {
    Temp    float32
    Hum     float32
}

type TimeInfo struct {
    Day     int
    Month   int
    Year    int
    Hour    int
    Min     int
}

type SysInfo struct {
    Env     EnvInfo
    Time    TimeInfo
    Heating bool
}

func GetInfo() SysInfo{
    timestamp := time.Now()
    time := TimeInfo{
        Day: time.Time.Day(timestamp),
        Month: int(time.Time.Month(timestamp)),
        Year: time.Time.Year(timestamp),
        Hour: time.Time.Hour(timestamp),
        Min: time.Time.Minute(timestamp),
    }

    env := EnvInfo {
        Temp: 20,
        Hum: 25.5,
    }

    info := SysInfo {
        Env: env,
        Time: time,
        Heating: false,
    }

    return info
}
