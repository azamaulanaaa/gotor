package storage

import (
	"os"
	"sync"

	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
	"github.com/spf13/afero"
)

type Storage struct {
    fileSystem afero.Fs
	config StorageConfig
}

type StorageConfig struct {
}

func NewStorage(fileSystem afero.Fs, config StorageConfig) anacrolixStorage.ClientImpl {
	return &Storage{
        fileSystem:     fileSystem,
		config:         config,
	}
}

func (storage Storage) OpenTorrent(info *anacrolixMetainfo.Info, infoHash anacrolixMetainfo.Hash) (anacrolixStorage.TorrentImpl, error) {
    path := "./" + infoHash.HexString()
    storage.fileSystem.MkdirAll(path, 0640)

    fileSystem := afero.NewBasePathFs(storage.fileSystem, path)

    database_file, err := fileSystem.OpenFile(
        infoHash.HexString() + ".db",
        os.O_CREATE | os.O_RDWR,
        0640,
    )

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
