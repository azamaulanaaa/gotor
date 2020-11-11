package main

import (
	"net/http"
	"ohkaca/lib/storage"

	"github.com/anacrolix/torrent"
	"github.com/gorilla/mux"
)

const host = "localhost:3000"

func main() {
	customConfig := torrent.NewDefaultClientConfig()
	customConfig.DefaultStorage = &storage.Storage{}
	c, err := torrent.NewClient(customConfig)
	must(err)
	defer c.Close()

	router := mux.NewRouter()
	http.ListenAndServe(host, router)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
