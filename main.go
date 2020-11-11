package main

import (
	"github.com/anacrolix/torrent"
)

func main() {
	c, err := torrent.NewClient(nil)
	must(err)
	defer c.Close()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
