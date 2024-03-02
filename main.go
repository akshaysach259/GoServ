package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Printf("Starting server at port : 8080 \n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error while listning to server", err)
	}
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err: %v", err)
		return
	}
	fmt.Fprint(w, "Post Request Success")
	firstName := r.FormValue("FirstName")
	lastName := r.FormValue("LastName")

	fmt.Fprintf(w, "FirstName %v and LastName : %v", firstName, lastName)

}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello!!")
}
