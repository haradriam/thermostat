package main

type EnvInfo struct {
    Temp    float32
    Hum     float32
}

type TimeInfo struct {
    Day     int
    Month   int
    Year    int
    Hour    int
    Min     int
}

type SysInfo struct {
    Env     EnvInfo
    Time    TimeInfo
    Heating bool
}

type EventEntry struct {
    Id          int
    StartTemp   float32
    Monday      int
    Tuesday     int
    Wednesday   int
    Thursday    int
    Friday      int
    Saturday    int
    Sunday      int
    StartHour   int
    StartMin    int
    EndHour     int
    EndMin      int
    Active      int
}
