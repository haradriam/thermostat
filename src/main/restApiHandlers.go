package main

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "strings"
)

/*RestGetInfo: Return system information as JSON object
*******************************************************/
func RestGetInfo(w http.ResponseWriter, r *http.Request) {
    //Encode system information and send
    json.NewEncoder(w).Encode(GetSysInfo())
}

/*RestGetUsage: Ask for usage entries and encode them into a JSON object
************************************************************************/
func RestGetUsage(w http.ResponseWriter, r *http.Request) {

}

/*RestGetConfig: Return system configuration as JSON object
***********************************************************/
func RestGetConfig(w http.ResponseWriter, r *http.Request) {
    //Encode system configuraton and send
    json.NewEncoder(w).Encode(GetConfig())
}

/*RestSetConfig: Decode JSON object to get new system configuration
*******************************************************************/
func RestSetConfig(w http.ResponseWriter, r *http.Request) {
    //Read body of HTTP request
    body, err := ioutil.ReadAll(r.Body)
    checkErr(err)

    //Create new JSON decoder based on readed HTTP body
    decoder := json.NewDecoder(strings.NewReader(string(body)))

    //Remove [ ] from the JSON string
    decoder.Token()

    //Decode received JSON into a system configuration struct
    var newConfig Config
    decoder.Decode(&newConfig)

    //Save the received new configuration
    SetConfig(newConfig)
}

/*RestGetEvents: Ask for the list of event and encode it into a JSON object
***************************************************************************/
func RestGetEvents(w http.ResponseWriter, r *http.Request) {
    //Request current events
    eventList, err := DbReadEvents()
    checkErr(err)

    //Encode the list of events and send
    json.NewEncoder(w).Encode(eventList)
}

/*RestSetEvents: Decode JSON object to get new list of events
*************************************************************/
func RestSetEvents(w http.ResponseWriter, r *http.Request) {
    //Read body of HTTP request
    body, err := ioutil.ReadAll(r.Body)
    checkErr(err)

    //Create new JSON decoder based on readed HTTP body
    decoder := json.NewDecoder(strings.NewReader(string(body)))

    //Remove [ ] from the JSON string
    decoder.Token()

    //Tour the received event list
    var eventList []EventEntry
    var event EventEntry

    for decoder.More() {
        //Read new event
        decoder.Decode(&event)

        //Append the readed event to the event list.
        eventList = append(eventList, event)
    }

    //Save the received event list
    DbAddEvents(eventList)
}
