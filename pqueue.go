package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// HomeHandler Hello page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Welcome to pqueue")
}

// PopQueue pops value from queue name route
func PopQueue(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Pop from queue %s", vars["name"])
}

// PushQueue pushes value to queue name route
func PushQueue(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Push to queue %s", vars["name"])
}

// main starts http server
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/queueu/{name}", PopQueue).Methods("GET")
    r.HandleFunc("/queue/{name}", PushQueue).Methods("POST")
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":10001", nil))
}
