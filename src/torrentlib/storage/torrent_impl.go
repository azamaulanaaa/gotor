package storage

import (
    "sync"

    anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
    anacrolixStorage "github.com/anacrolix/torrent/storage"
    "github.com/spf13/afero"
)

type TorrentImpl struct {
    fileSystem  afero.Fs
    database    *Database
    locks       []sync.RWMutex
    completion  Completion
}

func (torrentImpl *TorrentImpl) Close() error {
    torrentImpl.database.Close()
    return torrentImpl.Close() 
}


func (torrentImpl *TorrentImpl) Piece(torrentPiece anacrolixMetainfo.Piece) anacrolixStorage.PieceImpl {
    return Piece{
        torrentPiece:   torrentPiece,
        fileSystem:     torrentImpl.fileSystem,
        lock:           &torrentImpl.locks[torrentPiece.Index()],
        completion:     NewCompletion(torrentImpl.database, torrentPiece.Index()),
    }
}
