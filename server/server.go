package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var GitCommit string

func printVersion() {
	log.Printf("Current build version: %s", GitCommit)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")

	key := r.URL.Query().Get("key")
	fmt.Println("key =>", key)
	fmt.Println("Endpoint Hit: homePage")
	output := callCmd(key)
	fmt.Fprintf(w, "Output:\n%s\n", string(output))
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
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	printVersion()
	handleRequests()
}
