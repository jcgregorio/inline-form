package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const template = `
<input type="hidden" name="status" value="%s">
<p>&%s;</p>
<p><input type="submit" value="Toggle"></p>
`

func formHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	status := "check"
	if r.FormValue("status") == "check" {
		status = "cross"
	}
	fmt.Fprintf(w, template, status, status)
}

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/form", formHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
