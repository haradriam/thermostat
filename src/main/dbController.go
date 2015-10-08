package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

/*DbAddHist: Add new historic record to the database
****************************************************/
func DbAddHist(info SysInfo) error {
    //Open database
    db, err := sql.Open("sqlite3", GetConfig().DBPath)
    checkErr(err)
    defer db.Close()

    //Command to run in the database: insert new record
    stmt, err := db.Prepare("INSERT INTO HISTRECORDS(YEAR, MONTH, DAY, HOUR, MINUTE, TEMP, HUM) VALUES(?,?,?,?,?,?,?)")
    checkErr(err)

    //Execute the command: insert new record
    _, err = stmt.Exec(info.Time.Year, info.Time.Month, info.Time.Day, info.Time.Hour, info.Time.Min, info.Env.Temp, info.Env.Hum)
    checkErr(err)

    return err
}

/*DbAddEvents: Add new event list to the database
*************************************************/
func DbAddEvents(eventList []EventEntry) error {
    //Open database
    db, err := sql.Open("sqlite3", GetConfig().DBPath)
    checkErr(err)
    defer db.Close()

    //Command to run in the database: delete all events
    stmt, err := db.Prepare("DELETE FROM EVENTS")
    checkErr(err)

    //Execute the command: delete all events
    _, err = stmt.Exec()
    checkErr(err)

    //Command to run in the database: insert new event
    stmt, err = db.Prepare("INSERT INTO EVENTS(START_TEMP, PERIODIC, BY_TIME, START_HOUR, START_MIN, END_HOUR, END_MIN, ACTIVE) VALUES(?,?,?,?,?,?,?,?)")
    checkErr(err)

    //Tour the list of events
    for i := range eventList {
        //Execute the command: insert new event
        _, err = stmt.Exec(eventList[i].StartTemp, eventList[i].Periodic, eventList[i].ByTime, eventList[i].StartHour, eventList[i].StartMin, eventList[i].EndHour, eventList[i].EndMin, eventList[i].Active)
        checkErr(err)
    }

    return err
}

/*DbReadEvents: Read event list from the database
*************************************************/
func DbReadEvents() ([]EventEntry, error) {
    //Open database
    db, err := sql.Open("sqlite3", GetConfig().DBPath)
    checkErr(err)
    defer db.Close()

    //Command to run in the database: read all events
    rows, err := db.Query("SELECT * FROM EVENTS")
    checkErr(err)

    var eventList []EventEntry
    var event EventEntry

    //Tour the affected rows
    for rows.Next() {
        //Read new event
        err = rows.Scan(&event.StartTemp, &event.Periodic, &event.ByTime, &event.StartHour, &event.StartMin, &event.EndHour, &event.EndMin, &event.Active)
        checkErr(err)

        //Add event to the list
        eventList = append(eventList, event)
    }

    return eventList, err
}

/*DbAddUsageRecord: Store the status of the heating in the database
*******************************************************************/
func DbAddUsageRecord(usageEntry UsageEntry) {
    //Open database
    db, err := sql.Open("sqlite3", GetConfig().DBPath)
    checkErr(err)
    defer db.Close()

    //Command to run in the database: insert new usage record
    stmt, err := db.Prepare("INSERT INTO USAGE(YEAR, MONTH, DAY, START_HOUR, START_MIN, START_SEC, END_HOUR, END_MIN, END_SEC) VALUES(?,?,?,?,?,?,?,?,?)")
    checkErr(err)

    //Execute the command: insert new usage record
    _, err = stmt.Exec(usageEntry.Year, usageEntry.Month, usageEntry.Day, usageEntry.StartHour, usageEntry.StartMin, usageEntry.StartSec, usageEntry.EndHour, usageEntry.EndMin, usageEntry.EndSec)
}
