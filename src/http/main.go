package http

import (
	"net/http"
	"fmt"

	"HelloGo/src/http/wp"
)

func Request(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()



	fmt.Fprintln(rw, "xxxxxxxxxxxxxxx")
}



func StartService() {

	wp.RunService()

}
