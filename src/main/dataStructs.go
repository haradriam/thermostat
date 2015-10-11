package main

/*System configuration
**********************/
type Config struct {
    DBPath      string      //Path to the system database
    PerRecords  int         //Period to write new record in the database
    PerChecks   int         //Period to check the conditions to enable/disable the heating
    MaxTemp     float32     //Maximum temperature
    WebPath     string      //Path to the web resources
}

/*System information
********************/
type SysInfo struct {
    Env         EnvInfo     //Environment information
    Time        TimeInfo    //Time information
    Heating     bool        //Status of the heating
}

/*Environment information 
*************************/
type EnvInfo struct {
    Temp        float32     //Temperature
    Hum         float32     //Humidity
}

/*Time information
******************/
type TimeInfo struct {
    Date        string      //Date
    DayOfWeek   int         //Day of the week (1 ... 7)
}

/*History record
****************/
type HistRec struct {
    Date        string      //Date
    Temp        float32     //Temperature
    Hum         float32     //Humidity
}

/*Event entry
*************/
type EventEntry struct {
    StartTemp   float32     //Temperature to start the heating
    Periodic    string      //Days of the week to repeat this event (0000000)
    StartTime   string      //Minimum time to take this event into accotunt
    EndTime     string      //Maximum time to take this event into account
    Active      int         //Indicates if this event is enabled or not
}

/*Usage entry
*************/
type UsageEntry struct {
    StartDate   string      //Start date
    EndDate     string      //End date
}

/*HistQuery
***********/
type HistQuery struct {
    StartDate   string      //Start date
    EndDate     string      //End date
}
