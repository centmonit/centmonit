package core

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	// "github.com/julienschmidt/httprouter"
	// "gopkg.in/natefinch/lumberjack.v2"
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
)

var hostname string
var socket *websocket.Conn

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
	log.Printf("Request body: %s", b)
	// log.Printf("Request header XXX: %s", r.Header.Get("XXX"))
	// log.Printf("Request query pwd: %s", r.URL.Query().Get("pwd"))

	hostname = TestParse(string(b[:]))
	if socket != nil {
		socket.WriteMessage(1, []byte(fmt.Sprintf("The host is %s", hostname)))
	} else {
		log.Println("socket was nil")
	}

	// fmt.Fprint(rw, "OK")
	rw.WriteHeader(http.StatusOK)
}

func getSharedData(rw http.ResponseWriter, r *http.Request) {
	CORS(&rw)
	fmt.Fprintf(rw, "Here is the hostname: %s", hostname)
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func socketHandler(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket connection
	var err error
    socket, err = upgrader.Upgrade(rw, r, nil)
    if err != nil {
        log.Println(err)
	}

	// helpful log statement to show connections
	log.Println("[Socket] Client Connected")

	err = socket.WriteMessage(1, []byte("Hi Client!!!"))
    if err != nil {
        log.Println(err)
	}

	socket.WriteMessage(1, []byte(fmt.Sprintf("The host is: %s", hostname)))
}

func StartApiServer(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/api/collector", collectorPostHandler).Methods("POST")
	r.HandleFunc("/api/test", getSharedData).Methods("GET")
	r.HandleFunc("/api/hosts/report", getSharedData).Methods("GET")
	r.HandleFunc("/api/hosts/{host_id}/report", getSharedData).Methods("GET")

	// r.ServeFiles("/*filepath", http.Dir("./html"))

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
