package main

import (
    "log"
	"time"

	"gotor/src/http"
    "gotor/src/torrentlib"

    "github.com/spf13/afero"
)

func main() {
    config := LoadConfig()
  
    torrentClientConfig := torrentlib.DefaultClientConfig()
    torrentClientConfig.FileSystem = afero.NewMemMapFs()
    torrentClientConfig.Lifetime = time.Duration(config.TorrentClient.PieceLifetime) * time.Second
    torrentClientConfig.ClenaUpInterval = time.Duration(config.TorrentClient.CleanUpPeaceInterval) * time.Second

    client, err := torrentlib.NewClient(torrentClientConfig)
    if err != nil {
        log.Fatal(err)
    }

    httpServer := http.NewHttpServer(config.HttpServe.Port, client)
    log.Fatal(httpServer.ListenAndServe())
}

