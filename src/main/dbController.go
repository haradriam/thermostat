package main

import (
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type  HistRecordEntry struct {
    Day     int
    Month   int
    Year    int
    Hour    int
    Minute  int
    Temp    float32
    Hum     float32
}

func AddHist(info SysInfo) error {
    var err error = nil

    db, err := sql.Open("sqlite3", "/home/adrian/workspace/thermostat/bin/THERMOSTAT")
    checkErr(err)

    stmt, err := db.Prepare("INSERT INTO HISTRECORDS(DAY, MONTH, YEAR, HOUR, MINUTE, TEMP, HUM) VALUES(?,?,?,?,?,?,?)")
    checkErr(err)

    _, err = stmt.Exec(info.Time.Day, info.Time.Month, info.Time.Year, info.Time.Hour, info.Time.Min, info.Env.Temp, info.Env.Hum)
    checkErr(err)
    fmt.Println(err)

    return err
}
