package main

var usageEntry UsageEntry

/*StartHeating: Register the init time and start heating
********************************************************/
func StartHeating() {
    //Read system information and setup new usage entry
    sysInfo := GetSysInfo()
    usageEntry.Year = sysInfo.Time.Year
    usageEntry.Month = sysInfo.Time.Month
    usageEntry.Day = sysInfo.Time.Day
    usageEntry.StartHour = sysInfo.Time.Hour
    usageEntry.StartMin = sysInfo.Time.Min
    usageEntry.StartSec = sysInfo.Time.Sec

    //TODO: Start heating
}

/*StopHeating: Register the stop time and stop the heating
**********************************************************/
func StopHeating() {
    //Read system information and store the usage entry
    sysInfo := GetSysInfo()
    if sysInfo.Time.Day != usageEntry.Day {
        //Heating started and stopped in different days
        usageEntry.EndHour = 23
        usageEntry.EndMin = 59
        usageEntry.EndSec = 59

        DbAddUsageRecord(usageEntry)

        usageEntry.Year = sysInfo.Time.Year
        usageEntry.Month = sysInfo.Time.Month
        usageEntry.Day = sysInfo.Time.Day
        usageEntry.StartHour = 00
        usageEntry.StartMin = 00
        usageEntry.StartSec = 00
    }
    usageEntry.EndHour = sysInfo.Time.Hour
    usageEntry.EndMin = sysInfo.Time.Min
    usageEntry.EndSec = sysInfo.Time.Sec

    DbAddUsageRecord(usageEntry)

    //TODO Stop heating
}
