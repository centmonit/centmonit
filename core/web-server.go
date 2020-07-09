package core

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/julienschmidt/httprouter"
	// "gopkg.in/natefinch/lumberjack.v2"
)

func collectorPostHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// log.Println("/collector POST: ", p)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// log.Printf("Type: %T", b)
	log.Printf("Request body: %s", b)
	// log.Printf("Request header XXX: %s", r.Header.Get("XXX"))
	// log.Printf("Request query pwd: %s", r.URL.Query().Get("pwd"))

	TestParse(string(b[:]))

	fmt.Fprint(rw, "OK")
}

func StartServer(port string) {
	r := httprouter.New()
	r.POST("/collector", collectorPostHandler)

	// log.SetOutput(&lumberjack.Logger{
	// 	Filename:   "./logs/log.txt",
	// 	MaxSize:    1, // MB
	// 	MaxBackups: 5,
	// 	MaxAge:     1, //days
	// 	Compress:   false, // disabled by default
	// })
	log.Fatal(http.ListenAndServe(":" + port, r))
}
