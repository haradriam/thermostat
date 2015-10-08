package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func DbAddHist(info SysInfo) error {
    var err error = nil

    db, err := sql.Open("sqlite3", GetConfig().DBPath)
    checkErr(err)
    defer db.Close()

    stmt, err := db.Prepare("INSERT INTO HISTRECORDS(DAY, MONTH, YEAR, HOUR, MINUTE, TEMP, HUM) VALUES(?,?,?,?,?,?,?)")
    checkErr(err)

    _, err = stmt.Exec(info.Time.Day, info.Time.Month, info.Time.Year, info.Time.Hour, info.Time.Min, info.Env.Temp, info.Env.Hum)
    checkErr(err)

    return err
}

func DbAddEvents(eventList []EventEntry) error {
    var err error = nil

    db, err := sql.Open("sqlite3", GetConfig().DBPath)
    checkErr(err)
    defer db.Close()

    stmt, err := db.Prepare("DELETE FROM EVENTS")
    checkErr(err)

    _, err = stmt.Exec()
    checkErr(err)

    stmt, err = db.Prepare("INSERT INTO EVENTS(START_TEMP, PERIODIC, BY_TIME, START_HOUR, START_MIN, END_HOUR, END_MIN, ACTIVE) VALUES(?,?,?,?,?,?,?,?)")
    checkErr(err)

    for i := range eventList {
        _, err = stmt.Exec(eventList[i].StartTemp, eventList[i].Periodic, eventList[i].ByTime, eventList[i].StartHour, eventList[i].StartMin, eventList[i].EndHour, eventList[i].EndMin, eventList[i].Active)
        checkErr(err)
    }

    return err
}

func DbReadEvents() ([]EventEntry, error) {
    db, err := sql.Open("sqlite3", GetConfig().DBPath)
    checkErr(err)
    defer db.Close()

    rows, err := db.Query("SELECT * FROM EVENTS")
    checkErr(err)

    var eventList []EventEntry
    var event EventEntry

    for rows.Next() {

        err = rows.Scan(&event.StartTemp, &event.Periodic, &event.ByTime, &event.StartHour, &event.StartMin, &event.EndHour, &event.EndMin, &event.Active)
        checkErr(err)

        eventList = append(eventList, event)
    }

    return eventList, err
}
