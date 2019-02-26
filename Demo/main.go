package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Only log requests to our admin dashboard to stdout
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "//////")
	})
	router.HandleFunc("/v2", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "v1")
	})
	// Wrap our server with our gzip handler to gzip compress all responses.
	/*	http.ListenAndServe(":8000", handlers.CustomLoggingHandler(nil, router, func(writer io.Writer, params handlers.LogFormatterParams) {
			fmt.Printf("code=%v,request_url%v,timestamp=%s\n", params.StatusCode, params.URL, params.TimeStamp)
		}))
	*/
	// http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, router))

	router.Use(new(MiddleWare).Handler)
	http.ListenAndServe(":8000", router)
}

type MiddleWare struct {
	rw      http.ResponseWriter
	handler http.Handler
	code    int
	size    int
}

func (m *MiddleWare) Header() http.Header {
	return m.rw.Header()
}
func (m *MiddleWare) Write(b []byte) (int, error) {
	size, err := m.rw.Write(b)
	m.size += size
	return size, err
}
func (m *MiddleWare) WriteHeader(statusCode int) {
	m.code = statusCode
	m.WriteHeader(m.code)
}

func (m *MiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.rw = w
	m.code = http.StatusOK
	m.ServeHTTP(w, r)
}

func (m MiddleWare) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.handler.ServeHTTP(w, r)
		m.ServeHTTP(w, r)
		fmt.Println(m.code)
	})
}
