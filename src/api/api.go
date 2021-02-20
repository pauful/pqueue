package api

import (
	"github.com/pauful/pqueue/src/collections"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type App struct {
	QueuesMananger *collections.QueuesManager
	Router         *mux.Router
}

// HomeHandler Hello page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}

// PopQueue pops value from queue name route
func (a *App) popQueue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	value := a.QueuesMananger.Pop(vars["name"])
	w.Write(value)
}

// PushQueue pushes value to queue name route
func (a *App) pushQueue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	body, _ := ioutil.ReadAll(r.Body)
	a.QueuesMananger.Push(vars["name"], body)
	fmt.Fprintf(w, "OK")

}

func (a *App) lenQueue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", a.QueuesMananger.Len(vars["name"]))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (a *App) Initialise() {
	a.QueuesMananger = collections.NewQueuesManager()
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", homeHandler)
	a.Router.HandleFunc("/queue/{name}", a.popQueue).Methods("GET")
	a.Router.HandleFunc("/queue/{name}", a.pushQueue).Methods("POST")
	a.Router.HandleFunc("/queue/{name}/len", a.lenQueue).Methods("GET")
	a.Router.Use(loggingMiddleware)
	http.Handle("/", a.Router)
}

func (a *App) StartServer() {
	a.Initialise()
	log.Fatal(http.ListenAndServe(":10001", nil))
}
