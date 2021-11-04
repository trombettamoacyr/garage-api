package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	muxDispatcher = mux.NewRouter()
)

func (*muxRouter) SERVER(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}

func (*muxRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
