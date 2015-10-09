package main

import (
    "net/http"
)

/*StartRestAPI: Start new server and wait for queries
*****************************************************/
func StartRestAPI() {
    fs := http.FileServer(http.Dir("/home/adrian/workspace/thermostat/web/static"))
    http.Handle("/", fs)

    http.HandleFunc("/getinfo", RestGetInfo)        //REST method: Get system information
    http.HandleFunc("/getusage", RestGetUsage)      //REST method: Get usage entries
    http.HandleFunc("/getconfig", RestGetConfig)    //REST method: Get system configuration
    http.HandleFunc("/setconfig", RestSetConfig)    //REST method: Set system configuration
    http.HandleFunc("/getevents", RestGetEvents)    //REST method: Get list of events
    http.HandleFunc("/setevents", RestSetEvents)    //REST method: Set list of events

    //Start the HTTP server
    err := http.ListenAndServe(":8080", nil)
    checkErr(err)
}
