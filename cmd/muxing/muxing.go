package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func HandleNameView(w http.ResponseWriter, r *http.Request) {
	param, _ := mux.Vars(r)["PARAM"]
	fmt.Fprint(w, "Hello, "+param+"!")
}

func HandleBadView(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func HandleDataView(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Fprint(w, "I got message:\n"+string(body))
}

func HandleHeadersView(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")

	aa, _ := strconv.Atoi(a)
	bb, _ := strconv.Atoi(b)

	sum := strconv.Itoa(aa + bb)

	w.Header().Set("a+b", sum)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", HandleNameView).Methods(http.MethodGet)
	router.HandleFunc("/bad", HandleBadView).Methods(http.MethodGet)
	router.HandleFunc("/data", HandleDataView).Methods(http.MethodPost)
	router.HandleFunc("/headers", HandleHeadersView).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
