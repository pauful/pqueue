package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"queues"

	"github.com/gorilla/mux"
)

var queuesManager = queues.NewQueuesManager()

// HomeHandler Hello page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}

// PopQueue pops value from queue name route
func popQueue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	value := queuesManager.Pop(vars["name"])
	w.Write(value)
}

// PushQueue pushes value to queue name route
func pushQueue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	body, _ := ioutil.ReadAll(r.Body)
	queuesManager.Push(vars["name"], body)
	fmt.Fprintf(w, "OK")

}

func lenQueue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", queuesManager.Len(vars["name"]))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/queue/{name}", popQueue).Methods("GET")
	r.HandleFunc("/queue/{name}", pushQueue).Methods("POST")
	r.HandleFunc("/queue/{name}/len", lenQueue).Methods("GET")
	r.Use(loggingMiddleware)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":10001", nil))
}
