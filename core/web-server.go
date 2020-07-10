package core

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func StartWebServer(port string) {
	r := httprouter.New()
	r.ServeFiles("/*filepath", http.Dir("./html"))

	log.Fatal(http.ListenAndServe(":" + port, r))
}
