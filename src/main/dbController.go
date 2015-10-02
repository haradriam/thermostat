package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func AddHist(info SysInfo) error {
    var err error = nil

    db, err := sql.Open("sqlite3", "/home/adrian/workspace/thermostat/bin/THERMOSTAT")
    checkErr(err)

    stmt, err := db.Prepare("INSERT INTO HISTRECORDS(DAY, MONTH, YEAR, HOUR, MINUTE, TEMP, HUM) VALUES(?,?,?,?,?,?,?)")
    checkErr(err)

    _, err = stmt.Exec(info.Time.Day, info.Time.Month, info.Time.Year, info.Time.Hour, info.Time.Min, info.Env.Temp, info.Env.Hum)
    checkErr(err)

    return err
}

func AddEvent(events []EventEntry) error {
    var err error = nil
    for i := range events {
        db, err := sql.Open("sqlite3", "/home/adrian/workspace/thermostat/bin/THERMOSTAT")
        checkErr(err)

        stmt, err := db.Prepare("INSERT INTO EVENTS(ID, START_TEMP, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY, START_HOUR, START_MIN, END_HOUR, END_MIN, ACTIVE) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
        checkErr(err)

        _, err = stmt.Exec(events[i].Id, events[i].StartTemp, events[i].Monday, events[i].Tuesday, events[i].Wednesday, events[i].Thursday, events[i].Friday, events[i].Saturday, events[i].Sunday, events[i].StartHour, events[i].StartMin, events[i].EndHour, events[i].EndMin, events[i].Active)
        checkErr(err)
    }
    return err
}
