package k8s

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("Get")
	return r
}
func home(write http.ResponseWriter, request *http.Request) {
    fmt.Fprintln(write,"Hello you have request progress")
}
