package torrentlib

import (
    "errors"
    "time"

    "github.com/azamaulanaaa/gotor/src/torrentlib/storage"
    anacrolixLog "github.com/anacrolix/log"
    anacrolix "github.com/anacrolix/torrent"
    anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
    "github.com/spf13/afero"
)

type ClientConfig struct {
    Timeout         time.Duration
    FileSystem      afero.Fs
    Lifetime        time.Duration
    ClenaUpInterval time.Duration
}

func DefaultClientConfig() ClientConfig {
    return ClientConfig{
        Timeout         : 30 * time.Second,
        FileSystem      : afero.NewBasePathFs(afero.NewOsFs(), "./download"),
        Lifetime        : 5 * time.Minute,
        ClenaUpInterval : 1 * time.Minute,
    }
}

type Client struct {
    anacrolixClient *anacrolix.Client
    config          ClientConfig
}

func NewClient(config ClientConfig) (*Client, error) {
    anacrolixConfig := anacrolix.NewDefaultClientConfig()
    anacrolixConfig.Seed = false
    anacrolixConfig.NoUpload = true
    anacrolixConfig.Logger = anacrolixLog.Discard
    anacrolixConfig.DefaultStorage = storage.NewStorage(
        config.FileSystem,
        storage.StorageConfig{
            Lifetime: config.Lifetime,
            CleanUpInterval: config.ClenaUpInterval,
        },
    )
    
    anacrolixclient, err := anacrolix.NewClient(anacrolixConfig)
    if err != nil {
        return nil, err
    }


    client := Client{
        anacrolixClient: anacrolixclient,
        config: config,
    }

    return &client, nil
}

func (client *Client) AddHash(hash string) (Torrent, error) {
    var info anacrolixMetainfo.Hash
    if err := info.FromHexString(hash); err != nil {
        return Torrent{}, err
    }

    anacrolixtorrent, _ := client.anacrolixClient.AddTorrentInfoHash(info)
    anacrolixtorrent.DisallowDataDownload()

    select {
    case <-anacrolixtorrent.GotInfo():
    case <-time.After(client.config.Timeout):
        anacrolixtorrent.Drop()
        return Torrent{}, errors.New("Timeout")
    }

    torrent := Torrent{
        anacrolixTorrent: anacrolixtorrent,
    }

    return torrent, nil
}

