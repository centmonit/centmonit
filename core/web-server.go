package core

import (
	"fmt"
	"log"
	"net/http"
	// "github.com/julienschmidt/httprouter"
	"os/exec"
	// "os"
	// "bytes"
)

// find ./html -maxdepth 3 -type f | xargs sed -i 's/AAA/BBB/g'
func ConfigWebServer(apiHost string, apiPort string) {
	log.Printf("ConfigWebServer with host %s - port %s", apiHost, apiPort)

	cmd0 := "cp ./template/config.json ./html/config.json"
	_, _ = exec.Command("bash", "-c", cmd0).Output()

	cmd1 := fmt.Sprintf("find ./html -maxdepth 1 -type f | xargs sed -i 's/API_HOST/%s/g'", apiHost)
	_, _ = exec.Command("bash", "-c", cmd1).Output()

	cmd2 := fmt.Sprintf("find ./html -maxdepth 1 -type f | xargs sed -i 's/API_PORT/%s/g'", apiPort)
	_, _ = exec.Command("bash", "-c", cmd2).Output()
}

func StartWebServer(port string) {
	// r := httprouter.New()
	// r.ServeFiles("/*filepath", http.Dir("./html"))

	log.Fatal(http.ListenAndServe(":" + port, http.FileServer(http.Dir("./html"))))
}
