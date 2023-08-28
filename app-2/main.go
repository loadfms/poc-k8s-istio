package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("application running on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	result := process(param)

	fmt.Fprint(w, result)
}

func process(param string) string {
	// TODO: Implement the logic to process the parameter and return the result.
	return "Processed: " + param
}
