package main

import (
    "net/http"
    "html/template"
    "path"
)

/*StartRestAPI: Start new server and wait for queries
*****************************************************/
func StartRestAPI() {
    fs := http.FileServer(http.Dir("/home/adrian/workspace/thermostat/web"))
    http.Handle("/static/", fs)

    http.HandleFunc("/", serveTemplate)

    http.HandleFunc("/getinfo", RestGetInfo)        //REST method: Get system information
    http.HandleFunc("/gethist", RestGetHist)        //REST method: Get history records
    http.HandleFunc("/getusage", RestGetUsage)      //REST method: Get usage entries
    http.HandleFunc("/getconfig", RestGetConfig)    //REST method: Get system configuration
    http.HandleFunc("/setconfig", RestSetConfig)    //REST method: Set system configuration
    http.HandleFunc("/getevents", RestGetEvents)    //REST method: Get list of events
    http.HandleFunc("/setevents", RestSetEvents)    //REST method: Set list of events

    //Start the HTTP server
    err := http.ListenAndServe(":8080", nil)
    checkErr(err)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
    lp := path.Join(GetConfig().WebPath, "static/templates/layout.html")

    var fp string
    if(r.URL.Path == "/") {
        fp = path.Join(GetConfig().WebPath, "index.html")
    } else {
        fp = path.Join(GetConfig().WebPath, r.URL.Path)
    }

    tmpl, _ := template.ParseFiles(lp, fp)
    tmpl.ExecuteTemplate(w, "layout", nil)
}
