package core

import (
	// "fmt"
	"log"
	"io/ioutil"
	"net/http"
	// "github.com/julienschmidt/httprouter"
	// "gopkg.in/natefinch/lumberjack.v2"
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
)

const MAX_AGENTS = 2
var socketConnections = make([]*websocket.Conn, 0)

type MonitHost struct {
	ID string `json:"id"`
	Poll uint `json:"poll"`
	Hostname string `json:"hostname"`
	Uptime uint `json:"uptime"`
	RAM float32 `json:"ram"`
	CPU float32 `json:"cpu"`
	Services uint `json:"services"`
	GoodServices uint `json:"goodServices"`
	FailServices uint `json:"failServices"`
	SkipServices uint `json:"skipServices"`
	AlertMessage string `json:"alertMessage"`
}

var hostsMap = make(map[string]MonitHost)

type MonitHostService struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Monitor uint `json:"monitor"`
	Status uint `json:"status"`
}

type MonitHostServices struct {
	Data []MonitHostService
}

var hostsServicesMap = make(map[string]MonitHostServices)


func CORS(rw *http.ResponseWriter) {
	(*rw).Header().Set("Access-Control-Allow-Origin", "*")
}

func PreflightCORS(rw *http.ResponseWriter) {
	(*rw).Header().Set("Access-Control-Allow-Origin", "*")
	(*rw).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*rw).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func collectorPostHandler(rw http.ResponseWriter, r *http.Request) {
	// log.Println("/collector POST: ", p)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR\tcollectorPostHandler() error when read request body:", err)
	}

	// log.Printf("Type: %T", b)
	log.Printf("INFO\tAgent request body: %s", b)
	// log.Printf("Request header XXX: %s", r.Header.Get("XXX"))
	// log.Printf("Request query pwd: %s", r.URL.Query().Get("pwd"))

	// log.Println("Map pointer address: ", &hostsMap)
	ParseAgentReport(string(b[:]), &hostsMap)

	// fmt.Fprint(rw, "OK")
	rw.WriteHeader(http.StatusOK)
}

func hostsReportGetHandler(rw http.ResponseWriter, r *http.Request) {
	CORS(&rw)
	rw.Header().Add("Content-Type", "application/json")

	var hostArray = make([]MonitHost, len(hostsMap))

	if len(hostsMap) > 0 {
		var i = 0
		for _, value := range hostsMap {
			// fmt.Fprintf(rw, "key %s - value %s", key, value)
			// log.Printf("\t- API iterate value %+v\n", value)
			hostArray[i] = value
			i++
		}
	}

	json.NewEncoder(rw).Encode(hostArray)
}

func hostReportGetHandler(rw http.ResponseWriter, r *http.Request) {
	CORS(&rw)
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	hostID := vars["host_id"]
	// fmt.Fprintln(rw, "host report here:", hostID)

	var array1 = hostsServicesMap[hostID].Data
	var array2 = make([]MonitHostService, len(array1))

	var i = 0
	for _, value := range array1 {
		array2[i] = value
		i++
	}

	json.NewEncoder(rw).Encode(array2)
}

func hostReportByNameGetHandler(rw http.ResponseWriter, r *http.Request) {
	CORS(&rw)
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	hostName := vars["host_name"]

	var result []string
	for hostID, hostData := range hostsServicesMap {
		if hostsMap[hostID].Hostname == hostName {
			result = make([]string, len(hostData.Data))
			var i = 0
			for _, value := range hostData.Data {
				result[i] = value.Name
				i++
			}
			break
		}
	}
	json.NewEncoder(rw).Encode(result)
}

func channelsGetHandler(rw http.ResponseWriter, r *http.Request) {
	CORS(&rw)
	rw.Header().Add("Content-Type", "application/json")

	dataArray := ChannelGetAll()
	json.NewEncoder(rw).Encode(dataArray)
}

type ChannelCreationAPIResult struct {
	Status string `json:"status"`
	RowID string `json:"id,omitempty"`
}

type AlertRuleCreationAPIResult struct {
	Status string `json:"status"`
	RowID int `json:"id,omitempty"`
}

func channelsPostHandler(rw http.ResponseWriter, r *http.Request) {
	PreflightCORS(&rw)
	if (*r).Method == "OPTIONS" {
		return
	}

	rw.Header().Add("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR\tChannelsPostHandler() error when read request body:", err)
	}

	// log.Printf("INFO\tChannel post request body: %s", b)
	ok, rowID := ChannelSave(b)
	status := "error"
	if ok {
		status = "success"
	}
	json.NewEncoder(rw).Encode(ChannelCreationAPIResult{status, rowID})
}

func channelDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	PreflightCORS(&rw)
	if (*r).Method == "OPTIONS" {
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	chID := vars["channel_id"]

	ok := ChannelRemove(chID)
	status := "error"
	if ok {
		status = "success"
	}
	json.NewEncoder(rw).Encode(ChannelCreationAPIResult{status, ""})
}

func alertRulesGetHandler(rw http.ResponseWriter, r *http.Request) {
	CORS(&rw)
	rw.Header().Add("Content-Type", "application/json")

	dataArray := AlertRuleGetAll()
	json.NewEncoder(rw).Encode(dataArray)
}

func alertRulesPostHandler(rw http.ResponseWriter, r *http.Request) {
	PreflightCORS(&rw)
	if (*r).Method == "OPTIONS" {
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR\tAlertRulesPostHandler() error when read request body:", err)
	}

	log.Printf("INFO\tAlertRulesPostHandler request body: %s", b)
	ok, rowID := AlertRuleSave(b)
	status := "error"
	if ok {
		status = "success"
	}
	json.NewEncoder(rw).Encode(AlertRuleCreationAPIResult{status, rowID})
}

func alertRuleDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	PreflightCORS(&rw)
	if (*r).Method == "OPTIONS" {
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	idString := vars["rule_id"]

	idInt, _ := strconv.Atoi(idString)
	ok := AlertRuleRemove(idInt)
	status := "error"
	if ok {
		status = "success"
	}
	json.NewEncoder(rw).Encode(AlertRuleCreationAPIResult{status, 0})
}

func alertRuleUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	PreflightCORS(&rw)
	if (*r).Method == "OPTIONS" {
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	idString := vars["rule_id"]
	idInt, _ := strconv.Atoi(idString)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR\tAlertRuleUpdateHandler() error when read request body:", err)
	}

	log.Printf("INFO\tAlertRuleUpdateHandler request body: %s", b)

	ok := AlertRuleUpdate(idInt, b)
	status := "error"
	if ok {
		status = "success"
	}
	json.NewEncoder(rw).Encode(AlertRuleCreationAPIResult{status, 0})
}

func eventsPostHandler(rw http.ResponseWriter, r *http.Request) {
	PreflightCORS(&rw)
	if (*r).Method == "OPTIONS" {
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR\tEventsPostHandler() error when read request body:", err)
	}

	log.Printf("INFO\tEventsPostHandler request body: %s", b)
	// PublishEvent(string(b[:]))
	return
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func socketHandler(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket connection
	var err error
	var socket *websocket.Conn
	socket, err = upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
	} else {
		socketConnections = append(socketConnections, socket)
	}

	// helpful log statement to show connections
	log.Println("INFO\tNew socket connected")

	sockerMsg := SocketEventMessage{
		Type: "info",
		Host: "CentMonit",
		Message: "Success connect to Realtime channel.",
	}
	err = socket.WriteMessage(1, []byte(sockerMsg.StringValue()))
	if err != nil {
		log.Println("ERROR\tSocket write back error: ", err)
	}
}

func StartApiServer(port string) {
	r := mux.NewRouter()

	// For monit agent
	r.HandleFunc("/api/collector", collectorPostHandler).Methods("POST")

	// For web dashboard
	r.HandleFunc("/socket", socketHandler)

	r.HandleFunc("/api/hosts/report", hostsReportGetHandler).Methods("GET")
	r.HandleFunc("/api/hosts/{host_id}/report", hostReportGetHandler).Methods("GET")
	r.HandleFunc("/api/hostnames/{host_name}/report", hostReportByNameGetHandler).Methods("GET")

	r.HandleFunc("/api/channels", channelsGetHandler).Methods("GET")
	r.HandleFunc("/api/channels", channelsPostHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/channels/{channel_id}", channelDeleteHandler).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/alert-rules", alertRulesGetHandler).Methods("GET")
	r.HandleFunc("/api/alert-rules", alertRulesPostHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/alert-rules/{rule_id}", alertRuleDeleteHandler).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/alert-rules/{rule_id}", alertRuleUpdateHandler).Methods("PUT", "OPTIONS")

	r.HandleFunc("/api/events", eventsPostHandler).Methods("POST")


	// log.SetOutput(&lumberjack.Logger{
	// 	Filename:   "./logs/log.txt",
	// 	MaxSize:    1, // MB
	// 	MaxBackups: 5,
	// 	MaxAge:     1, //days
	// 	Compress:   false, // disabled by default
	// })
	log.Fatal(http.ListenAndServe(":" + port, r))
}
