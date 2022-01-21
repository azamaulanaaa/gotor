package storage

import (
    "os"
    "sync"
    "time"

    anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
    anacrolixStorage "github.com/anacrolix/torrent/storage"
    "github.com/spf13/afero"
)

type Storage struct {
    fileSystem      afero.Fs
    config          StorageConfig
    cleanUp         *CleanUp 
}

type StorageConfig struct {
    Lifetime            time.Duration
    CleanUpInterval     time.Duration
}

func NewStorage(fileSystem afero.Fs, config StorageConfig) anacrolixStorage.ClientImpl {
    storage := &Storage{
        fileSystem:     fileSystem,
        config:         config,
    }

    if config.Lifetime != 0 && config.CleanUpInterval != 0 {
        storage.cleanUp = NewCleanUp(fileSystem, config.Lifetime, config.CleanUpInterval)
        storage.cleanUp.StartService()
    }

    return storage
}

func (storage Storage) OpenTorrent(info *anacrolixMetainfo.Info, infoHash anacrolixMetainfo.Hash) (anacrolixStorage.TorrentImpl, error) {
    database_file, err := storage.fileSystem.OpenFile(
        infoHash.HexString() + ".db",
        os.O_CREATE | os.O_RDWR,
        0640,
    )

    path := "./" + infoHash.HexString()
    fileSystem := afero.NewBasePathFs(storage.fileSystem, path)

    if err != nil {
        return anacrolixStorage.TorrentImpl{}, err
    }

    torrentImpl := TorrentImpl{
        fileSystem: fileSystem,
        database:   NewDatabase(database_file),
        locks:      make([]sync.RWMutex, info.NumPieces()),
	}

    return anacrolixStorage.TorrentImpl{Piece: torrentImpl.Piece, Close: torrentImpl.Close}, nil
}
