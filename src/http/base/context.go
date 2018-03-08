package base

import "net/http"

type Context struct {
	Request  *http.Request
	response *http.Response
	http.ResponseWriter
}
