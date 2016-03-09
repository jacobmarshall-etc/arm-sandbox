package main

import (
	"net/http"
	"os"
	"strconv"
	"log"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
	httpPort string = os.Getenv("PORT")
	requestCount uint16 = 0
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello World (Total requests: " + strconv.Itoa(int(requestCount)) + ")"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	url, err := router.Get("home").URL()
	if err == nil {
		// Redirect the user back to the homepage
		http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
	}
}

func main() {
	router = mux.NewRouter()

	router.
		HandleFunc("/", HomeHandler).
		Name("home")

	// Register a 'not found handler' for all other requests
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	http.Handle("/", router)

	// Listen the server on the specified port, log fatal if error
	if err := http.ListenAndServe(":" + httpPort, nil); err != nil {
		log.Fatal(err)
	}
}