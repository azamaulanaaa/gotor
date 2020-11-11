package main

import (
	"ohkaca/lib/storage"

	"github.com/anacrolix/torrent"
)

func main() {
	customConfig := torrent.NewDefaultClientConfig()
	customConfig.DefaultStorage = &storage.Storage{}
	c, err := torrent.NewClient(customConfig)
	must(err)
	defer c.Close()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
