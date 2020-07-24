package core

import (
	"fmt"
	"log"
	"net/http"
	// "github.com/julienschmidt/httprouter"
	"os/exec"
	// "os"
	// "bytes"
	auth "github.com/abbot/go-http-auth"
	"strings"
)

var webAccounts = make(map[string]string)

// find ./html -maxdepth 3 -type f | xargs sed -i 's/AAA/BBB/g'
func ConfigWebServer(apiHost string, apiPort string, authArray []string) {
	cmd0 := "cp ./template/config.json ./html/config.json"
	_, _ = exec.Command("bash", "-c", cmd0).Output()

	cmd1 := fmt.Sprintf("find ./html -maxdepth 1 -type f | xargs sed -i 's/API_HOST/%s/g'", apiHost)
	_, _ = exec.Command("bash", "-c", cmd1).Output()

	cmd2 := fmt.Sprintf("find ./html -maxdepth 1 -type f | xargs sed -i 's/API_PORT/%s/g'", apiPort)
	_, _ = exec.Command("bash", "-c", cmd2).Output()

	for _, user := range authArray {
		tmp := strings.Split(user, ":")
		webAccounts[tmp[0]] = tmp[1]
	}
	log.Printf("INFO\tLoaded web dashboard accounts: %d", len(webAccounts))
}

func __StartWebServer(port string) {
	// r := httprouter.New()
	// r.ServeFiles("/*filepath", http.Dir("./html"))

	log.Fatal(http.ListenAndServe(":" + port, http.FileServer(http.Dir("./html"))))
}

func secret(user, realm string) string {
	// users := map[string]string{
	// 	"john": "$apr1$ZQoIO.79$IaNltzXgnvcypObxUh4um0",
	// }

	if a, ok := webAccounts[user]; ok {
		return a
	}
	return ""
}

func handleFileServer(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		// log.Println("Request URL:", req.URL)
		realHandler(w, req)
	}
}

func StartWebServer(port string) {
	authenticator := auth.NewBasicAuthenticator("", secret)
	http.HandleFunc("/", auth.JustCheck(authenticator, handleFileServer("./html", "/")))
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
