package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
)

var GitCommit string

func printVersion() {
	log.Printf("Current build version: %s", GitCommit)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage CowSay!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func cowPage(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}

	key := query.Get("key")
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing parameter: ?key=")
		return
	}
	output := callCmd(key)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cow Output:\n%s\n", string(output))
}

func versionPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the VersionPage! \n\n\n")

	output := callCmd(GitCommit)
	fmt.Fprintf(w, "Version:\n%s\n", string(output))
}

func callCmd(param1 string) []byte {
	cmd := exec.Command("cowsay", param1)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/version/", versionPage)
	http.HandleFunc("/cow/", cowPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	printVersion()
	handleRequests()
}
