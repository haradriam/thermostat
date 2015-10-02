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

    db.Close()

    return err
}

func AddEvents(events []EventEntry) error {
    var err error = nil

    db, err := sql.Open("sqlite3", "/home/adrian/workspace/thermostat/bin/THERMOSTAT")
    checkErr(err)

    stmt, err := db.Prepare("DELETE FROM EVENTS")
    checkErr(err)

    _, err = stmt.Exec()
    checkErr(err)

    stmt, err = db.Prepare("INSERT INTO EVENTS(ID, START_TEMP, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY, START_HOUR, START_MIN, END_HOUR, END_MIN, ACTIVE) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
    checkErr(err)

    for i := range events {
        _, err = stmt.Exec(events[i].Id, events[i].StartTemp, events[i].Monday, events[i].Tuesday, events[i].Wednesday, events[i].Thursday, events[i].Friday, events[i].Saturday, events[i].Sunday, events[i].StartHour, events[i].StartMin, events[i].EndHour, events[i].EndMin, events[i].Active)
        checkErr(err)
    }

    db.Close()

    return err
}

func ReadEvents() ([]EventEntry, error) {
    db, err := sql.Open("sqlite3", "/home/adrian/workspace/thermostat/bin/THERMOSTAT")
    checkErr(err)

    rows, err := db.Query("SELECT * FROM EVENTS")
    checkErr(err)

    var eventList []EventEntry
    var event EventEntry

    for rows.Next() {

        err = rows.Scan(&event.Id, &event.StartTemp, &event.Monday, &event.Tuesday, &event.Wednesday, &event.Thursday, &event.Friday, &event.Saturday, &event.Sunday, &event.StartHour, &event.StartMin, &event.EndHour, &event.EndMin, &event.Active)
        checkErr(err)

        eventList = append(eventList, event)
    }

    db.Close()

    return eventList, err
}
