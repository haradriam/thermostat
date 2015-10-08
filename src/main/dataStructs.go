package main

type Config struct {
    DBPath      string
    PerRecords  int
    PerChecks   int
}

type EnvInfo struct {
    Temp        float32
    Hum         float32
}

type TimeInfo struct {
    Day         int
    Month       int
    Year        int
    Hour        int
    Min         int
    DayOfWeek   int
}

type SysInfo struct {
    Env         EnvInfo
    Time        TimeInfo
    Heating     bool
}

type EventEntry struct {
    StartTemp   float32
    Periodic    string
    ByTime      int
    StartHour   int
    StartMin    int
    EndHour     int
    EndMin      int
    Active      int
}
