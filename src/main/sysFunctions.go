package main

import (
    "time"
)

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
