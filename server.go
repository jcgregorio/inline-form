package main

import (
	"fmt"
	"log"
	"net/http"
)

const template = `
<input type="hidden" name="status" value="%s">
<p>&%s;</p>
<p><input type="submit" value="Toggle"></p>
`

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	formHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		status := "check"
		if r.FormValue("status") == "check" {
			status = "cross"
		}
		fmt.Fprintf(w, template, status, status)
	}

	http.HandleFunc("/form", formHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
