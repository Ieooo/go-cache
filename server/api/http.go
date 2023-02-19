package api

import (
	"cache/pkg/log"
	"cache/server/core"
	"net/http"
	"path"
)

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (h HttpServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	log.Debugf("%v %v %v", path.Base(req.URL.Path), params.Get("k"), params.Get("v"))
	switch req.URL.Path {
	case core.BasePath + "/set":
		core.Set(params.Get("k"), params.Get("v"))
	case core.BasePath + "/get":
		value := get(params.Get("k"))
		rw.Write([]byte(value))
	case core.BasePath + "/del":
		core.Del(params.Get("k"))
	default:
		rw.Write([]byte("bad request"))
	}

}

func get(key string) string {
	val, err := core.Get(key)
	if err != nil {
		return err.Error()
	}
	return val
}
