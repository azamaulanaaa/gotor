package main

import (
	"net/http"
	"github.com/azamaulanaaa/ohkaca/storage"
  "github.com/azamaulanaaa/ohkaca/router"

	anacrolixTorrent "github.com/anacrolix/torrent"
)

const host = ":8000"

func main() {
	customConfig := anacrolixTorrent.NewDefaultClientConfig()
	customConfig.DefaultStorage = storage.New("./tmp")
	customConfig.NoUpload = true
	client, err := anacrolixTorrent.NewClient(customConfig)
	must(err)
	defer client.Close()

  router, err := router.New(client)
  must(err)
	must(http.ListenAndServe(host, router))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
