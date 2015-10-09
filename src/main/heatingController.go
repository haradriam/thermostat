package main

var usageEntry UsageEntry
var heatingStatus bool = false

/*StartHeating: Register the init time and start heating
********************************************************/
func StartHeating() {
    //Read system information and setup new usage entry
    sysInfo := GetSysInfo()
    usageEntry.StartDate = sysInfo.Time.Date

    //TODO: Start heating
    heatingStatus = true
}

/*StopHeating: Register the stop time and stop the heating
**********************************************************/
func StopHeating() {
    //Read system information and store the usage entry
    sysInfo := GetSysInfo()
    usageEntry.EndDate = sysInfo.Time.Date

    DbAddUsageRecord(usageEntry)

    //TODO Stop heating
    heatingStatus = false
}

func GetHeatingStatus() bool {
    return heatingStatus
}
