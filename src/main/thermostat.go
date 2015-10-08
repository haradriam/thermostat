package main

import (
    "time"
)

/*main: Main system function
****************************/
func main() {
    //Read the startup configuration
    ConfigFirstRead()

    //Run REST API on a different thread
    go StartRestAPI()

    //Run record writer on a different thread
    go writeHist()

    //Infinity loop
    for {
        //Read event list
        eventList, err :=  DbReadEvents()
        checkErr(err)

        //Read system information
        sysInfo := GetSysInfo()

        //Check each event to find if it meets the conditions to start the heating
        //TODO: Add maximum temperature condition
        for i := range eventList {
            checkCondition(eventList[i], sysInfo)
        }

        //Wait for the next check
        time.Sleep(time.Duration(GetConfig().PerChecks) * time.Second)
    }
}

/*writeHist: Record writer
**************************/
func writeHist() {
    //Infinite loop
    for {
        //Read system information and store it in the database
        err := DbAddHist(GetSysInfo())
        checkErr(err)

        //Wait for the next record
        time.Sleep(time.Duration(GetConfig().PerRecords) * time.Second)
    }
}

/*checkCondition: Checks if an event meets the conditions to start the heating
*****************************************************************************/
func checkCondition(event EventEntry, sysInfo SysInfo) bool {
    //Check if the event is active or not
    if event.Active == 0 { return false }

    //Check if current temperature if higher than the event's one
    if (event.StartTemp != 0) && (sysInfo.Env.Temp > event.StartTemp) { return false }

    //Check if the event is periodic and, if so, if it is applicable to the current day of week
    if (event.Periodic != "0000000") {
        if event.Periodic[sysInfo.Time.DayOfWeek - 1] == 0 { return false }
    }

    //Check if the event is time sensitive
    if (event.ByTime == 1) {
        //Check if the event time configuration is limited to the current day
        var night bool = false
        if event.StartHour == event.EndHour {
            if event.StartMin >= event.EndMin { night = true }
        } else {
            if event.StartHour > event.EndHour { night = true }
        }

        //Check if the start time is higher than current time
        enaStart := checkHour(event.StartHour, event.StartMin, sysInfo.Time.Hour, sysInfo.Time.Min)
        //Check if the stop time is higher than current time
        enaStop := checkHour(event.EndHour, event.EndMin, sysInfo.Time.Hour, sysInfo.Time.Min)

        //Decide if the time condition to enable to heating is met
        if night == false {
            if enaStart == enaStop { return false }
        } else {
            if enaStart != enaStop { return false }
        }
    }

    return true
}

/*checkHour: Check if configured time is highed than the current one
********************************************************************/
func checkHour(ConfHour int, ConfMin int, SysHour int, SysMin int) bool {
    if ConfHour < SysHour {
        return false
    }

    if ConfHour == SysHour {
        if ConfMin < SysMin {
            return false
        }
    }

    return true
}

/*checkErr: Common error checking function
******************************************/
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
