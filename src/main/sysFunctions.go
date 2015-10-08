package main

import (
    "time"
    "encoding/json"
    "os"
)

var config Config

/*GetSysInfo: Get system information and return it
**************************************************/
func GetSysInfo() SysInfo {
    //Get time information
    timestamp := time.Now()
    time := TimeInfo{
        Day: time.Time.Day(timestamp),
        Month: int(time.Time.Month(timestamp)),
        Year: time.Time.Year(timestamp),
        Hour: time.Time.Hour(timestamp),
        Min: time.Time.Minute(timestamp),
        DayOfWeek: int(time.Time.Weekday(timestamp)),
    }

    //Get environment information
    //TODO: Leer informacion real del sensor
    env := EnvInfo {
        Temp: 20,
        Hum: 25.5,
    }

    //Create new system information struct
    //TODO: Leer informacion real del rele
    info := SysInfo {
        Env: env,
        Time: time,
        Heating: false,
    }

    return info
}

/*GetConfig: Return current system configuration
************************************************/
func GetConfig() Config {
    return config
}

/*SetConfig: Update the system configuration with the new one and store it in the configuration file
****************************************************************************************************/
func SetConfig(newConfig Config) {
    //Create the file
    file, err := os.Create("conf.json")
    checkErr(err)
    defer file.Close()

    //Create new JSON encoder based on the created file
    encoder := json.NewEncoder(file)

    //Write the new configuracion in the file
    err = encoder.Encode(newConfig)
    checkErr(err)

    //Update current system configuration
    config = newConfig
}

/*ConfigFirstRead: Read configuration from the file and apply it
****************************************************************/
func ConfigFirstRead() {
    //Open configuration file
    file, err := os.Open("conf.json")
    checkErr(err)
    defer file.Close()

    //Create new JSON decoder based on opened file
    decoder := json.NewDecoder(file)

    //Decode reader JSON into a system configuration struct
    err = decoder.Decode(&config)
    checkErr(err)
}
