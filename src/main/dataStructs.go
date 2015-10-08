package main

/*System configuration
**********************/
type Config struct {
    DBPath      string      //Path to the system database
    PerRecords  int         //Period to write new record in the database
    PerChecks   int         //Period to check the conditions to enable/disable the heating
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
    Day         int         //Day of the month
    Month       int         //Month of the year
    Year        int         //Year number
    Hour        int         //Hour
    Min         int         //Minute
    DayOfWeek   int         //Day of the week (1 ... 7)
}

/*Event entry
*************/
type EventEntry struct {
    StartTemp   float32     //Temperature to start the heating
    Periodic    string      //Days of the week to repeat this event (0000000)
    ByTime      int         //Indicates whether the event is time sensitive
    StartHour   int         //Minimum hour to take this event into account
    StartMin    int         //Minimum minute to take this event into accotunt
    EndHour     int         //Maximum hour to take this event into account
    EndMin      int         //Maximum minute to take this event into account
    Active      int         //Indicates if this event is enabled or not
}
