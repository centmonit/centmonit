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

func collectorPostHandler(rw http.ResponseWriter, r *http.Request) {
	// log.Println("/collector POST: ", p)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
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
	r.HandleFunc("/api/hosts/report", hostsReportGetHandler).Methods("GET")
	r.HandleFunc("/api/hosts/{host_id}/report", hostReportGetHandler).Methods("GET")
	r.HandleFunc("/socket", socketHandler)

	// log.SetOutput(&lumberjack.Logger{
	// 	Filename:   "./logs/log.txt",
	// 	MaxSize:    1, // MB
	// 	MaxBackups: 5,
	// 	MaxAge:     1, //days
	// 	Compress:   false, // disabled by default
	// })
	log.Fatal(http.ListenAndServe(":" + port, r))
}
