package api

import (
	"cache/server/core"
	"net/http"
)

const basePath = "/cache"

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (h HttpServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	switch req.URL.Path {
	case basePath + "/set":
		core.Set(params.Get("k"), params.Get("v"))
	case basePath + "/get":
		value := core.Get(params.Get("k"))
		rw.Write([]byte(value))
	case basePath + "/del":
		core.Delete(params.Get("k"))
	default:
		rw.Write([]byte("bad request"))
	}

}
