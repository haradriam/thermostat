package main

import (
    "time"
)

func main() {
    ConfigFirstRead()

    go StartRestAPI()
    go writeHist()

    for {
        eventList, err :=  DbReadEvents()
        checkErr(err)

        sysInfo := GetSysInfo()
        for i := range eventList{
            checkCondition(eventList[i], sysInfo)
        }

        time.Sleep(time.Duration(GetConfig().PerChecks) * time.Second)
    }
}

func writeHist() {
    for {
        err := DbAddHist(GetSysInfo())
        checkErr(err)

        time.Sleep(time.Duration(GetConfig().PerRecords) * time.Second)
    }
}

func checkCondition(event EventEntry, sysInfo SysInfo) bool {
    if event.Active == 0 { return false }

    if (event.StartTemp != 0) && (event.StartTemp > sysInfo.Env.Temp) { return false }

    if (event.Periodic != "0000000") {
        if event.Periodic[sysInfo.Time.DayOfWeek - 1] == 0 { return false }
    }

    if (event.ByTime == 1) {
        var night bool = false

        if event.StartHour == event.EndHour {
            if event.StartMin >= event.EndMin { night = true }
        } else {
            if event.StartHour > event.EndHour { night = true }
        }

        enaStart := checkHour(event.StartHour, event.StartMin, sysInfo.Time.Hour, sysInfo.Time.Min)
        enaStop := checkHour(event.EndHour, event.EndMin, sysInfo.Time.Hour, sysInfo.Time.Min)

        if night == false {
            if enaStart == enaStop { return false }
        } else {
            if enaStart != enaStop { return false }
        }
    }

    return true
}

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

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
