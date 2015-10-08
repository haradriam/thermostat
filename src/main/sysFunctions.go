package main

import (
    "time"
    "encoding/json"
    "os"
)

var config Config

func GetSysInfo() SysInfo{
    timestamp := time.Now()
    time := TimeInfo{
        Day: time.Time.Day(timestamp),
        Month: int(time.Time.Month(timestamp)),
        Year: time.Time.Year(timestamp),
        Hour: time.Time.Hour(timestamp),
        Min: time.Time.Minute(timestamp),
        DayOfWeek: int(time.Time.Weekday(timestamp)),
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

func GetConfig() Config {
    return config
}

func SetConfig(newConfig Config) {
    file, err := os.Create("conf.json")
    checkErr(err)
    defer file.Close()

    encoder := json.NewEncoder(file)

    err = encoder.Encode(newConfig)
    checkErr(err)

    config = newConfig
}

func ConfigFirstRead() {
    file, err := os.Open("conf.json")
    checkErr(err)
    defer file.Close()

    decoder := json.NewDecoder(file)

    err = decoder.Decode(&config)
    checkErr(err)
}
