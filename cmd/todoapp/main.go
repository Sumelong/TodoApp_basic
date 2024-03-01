package main

import (
	_ "modernc.org/sqlite"
	"net/http"
)

type apiHandle struct{}

func (a apiHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandle{})

}
