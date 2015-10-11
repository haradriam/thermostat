package main

import (
    "time"
    "strings"
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
        eventList :=  DbReadEvents()

        //Read system information
        sysInfo := GetSysInfo()

        //Check if the maximum temperature has been reached
        if (sysInfo.Env.Temp > GetConfig().MaxTemp) && (sysInfo.Heating == true) {
            StopHeating()
        } else {
            //Check each event to find if it meets the conditions to start the heating
            cond := false
            for i := range eventList {
                if cond == false {
                    cond = checkCondition(eventList[i], sysInfo)
                }
            }

            //Check if any event has met the conditions to start the heating
            if cond == false {
                if sysInfo.Heating == true { StopHeating() }
            } else {
                if sysInfo.Heating == false { StartHeating() }
            }
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
        DbAddHist(GetSysInfo())

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
        if (event.Periodic[sysInfo.Time.DayOfWeek] - 1 == 0) { return false }
    }

    sep := strings.Split(sysInfo.Time.Date, " ");
    if ((event.StartTime != "") && (event.StartTime < sep[1])) { return false }
    if ((event.EndTime != "") && (event.EndTime > sep[1])) { return false }

    return true
}

/*checkErr: Common error checking function
******************************************/
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
