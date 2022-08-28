package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")

	key := r.URL.Query().Get("key")
	fmt.Println("key =>", key)
	fmt.Println("Endpoint Hit: homePage")
	output := callCmd(key)
	fmt.Fprintf(w, "Output:\n%s\n", string(output))
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
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
