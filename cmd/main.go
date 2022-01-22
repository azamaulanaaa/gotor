package main

import (
    "log"
	"time"

	"github.com/azamaulanaaa/gotor/src/http"
    "github.com/azamaulanaaa/gotor/src/torrentlib"

    "github.com/spf13/afero"
)

func main() {
    config := LoadConfig()
  
    torrentClientConfig := torrentlib.DefaultClientConfig()
    torrentClientConfig.Lifetime = time.Duration(config.TorrentClient.PieceLifetime) * time.Second
    torrentClientConfig.ClenaUpInterval = time.Duration(config.TorrentClient.CleanUpPeaceInterval) * time.Second
    torrentClientConfig.FileSystem = afero.NewBasePathFs(
        afero.NewOsFs(),
        config.TorrentClient.StorageDir,
    )
    if config.TorrentClient.MemoryStorage {
        torrentClientConfig.FileSystem = afero.NewMemMapFs()
    }

    client, err := torrentlib.NewClient(torrentClientConfig)
    if err != nil {
        log.Fatal(err)
    }

    httpServer := http.NewHttpServer(config.HttpServe.Port, client)
    log.Fatal(httpServer.ListenAndServe())
}

