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
    stmt, err := db.Prepare("INSERT INTO HISTRECORDS(DATE, TEMP, HUM) VALUES(?,?,?)")
    checkErr(err)

    //Execute the command: insert new record
    _, err = stmt.Exec(info.Time.Date, info.Env.Temp, info.Env.Hum)
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
    stmt, err = db.Prepare("INSERT INTO EVENTS(START_TEMP, PERIODIC, START_TIME, END_TIME, ACTIVE) VALUES(?,?,?,?,?)")
    checkErr(err)

    //Tour the list of events
    for i := range eventList {
        //Execute the command: insert new event
        _, err = stmt.Exec(eventList[i].StartTemp, eventList[i].Periodic, eventList[i].StartTime, eventList[i].EndTime, eventList[i].Active)
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
        err = rows.Scan(&event.StartTemp, &event.Periodic, &event.StartTime, &event.EndTime, &event.Active)
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
    stmt, err := db.Prepare("INSERT INTO USAGE(START_DATE, END_DATE) VALUES(?,?)")
    checkErr(err)

    //Execute the command: insert new usage record
    _, err = stmt.Exec(usageEntry.StartDate, usageEntry.EndDate)
}

/*DbReadUsageRecord: Read usage entries between dates
*****************************************************/
/*func DbReadUsageRecord(query UsageQuery) []UsageEntry {
    //Open database
    db, err := sql.Open("sqlite3", GetConfig().DBPath)
    checkErr(err)
    defer db.Close()

    //Command to run in the database: read usage entries between dates
    first, err := db.Query("SELECT ID FROM USAGE WHERE YEAR >= " +
                            strconv.Itoa(query.StartYear) +
                            "AND MONTH >= " +
                            strconv.Itoa(query.StartMonth) +
                            "AND DAY >= " +
                            strconv.Itoa(query.StartDay) +
                            "LIMIT 1")
    checkErr(err)

    last, err := db.Query("SELECT ID FROM USAGE WHERE YEAR >= " +
                            strconv.Itoa(query.StartYear) +
                            "AND MONTH >= " +
                            strconv.Itoa(query.StartMonth) +
                            "AND DAY >= " +
                            strconv.Itoa(query.StartDay) +
                            "LIMIT 1")
    checkErr(err)

    var initId int = 0
    var endId int = 0
    err = rows.Scan(&initId)
    checkErr(err)

    var usageEntryList []UsageEntry

    return usageEntryList
}*/
