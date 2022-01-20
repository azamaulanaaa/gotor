package storage

import (
	"fmt"
	"io"
	"os"

	"sync"

	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
	"github.com/spf13/afero"
)

type Piece struct {
	torrentPiece anacrolixMetainfo.Piece
    fileSystem afero.Fs
	lock *sync.RWMutex
    completion Completion
}

func (piece Piece) WriteTo(w io.Writer) (int64, error) {
	piece.lock.RLock()
	defer piece.lock.RUnlock()

    fileInstance, err := piece.fileSystem.Open(piece.path())
    defer fileInstance.Close()
    if err != nil {
        return 0, err
    }

    n, err := io.Copy(w, fileInstance)

    return n, err
}

func (piece Piece) ReadAt(b []byte, off int64) (int, error) {
	piece.lock.RLock()
	defer piece.lock.RUnlock()

    fileInstance, err := piece.fileSystem.Open(piece.path())
    defer fileInstance.Close()
    if err != nil {
        return 0, err
    }

    return fileInstance.ReadAt(b, off)
}

func (piece Piece) WriteAt(b []byte, off int64) (n int, err error) {
	piece.lock.RLock()
    defer piece.lock.RUnlock()

	fileInstance, err := piece.fileSystem.OpenFile(
        piece.path(),
        os.O_CREATE | os.O_WRONLY,
        0640,
    )
    defer fileInstance.Close()
    if err != nil {
        return
    }

    return fileInstance.WriteAt(b, off)
}

func (piece Piece) MarkComplete() error {
    piece.completion.Set(true)
    return nil
}

func (piece Piece) MarkNotComplete() error {
	piece.lock.Lock()
	defer piece.lock.Unlock()

    piece.completion.Set(false)
    return piece.fileSystem.Remove(piece.path())
}

func (piece Piece) Completion() anacrolixStorage.Completion {
    completion := anacrolixStorage.Completion{
        Complete: piece.completion.Get(),
		Ok:       true,
	}
    

    return completion
}

func (piece Piece) path() string {
    path := fmt.Sprintf("%d", piece.torrentPiece.Index())
    return path
}
