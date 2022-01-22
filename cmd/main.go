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

    torrentClientConfig.FileSystem = afero.NewMemMapFs()
    if config.TorrentClient.MemoryStorage == false {
        fileSystem := afero.NewOsFs()
        fileSystem.MkdirAll(config.TorrentClient.StorageDir, 0640)
        torrentClientConfig.FileSystem = afero.NewBasePathFs(
            fileSystem,
            config.TorrentClient.StorageDir,
        )
        
    }

    client, err := torrentlib.NewClient(torrentClientConfig)
    if err != nil {
        log.Fatal(err)
    }

    httpServer := http.NewHttpServer(config.HttpServe.Port, client)
    log.Fatal(httpServer.ListenAndServe())
}

