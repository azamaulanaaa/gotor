package storage

import (
  "os"

	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
)

type Piece struct {
  filename string
  pieceInfo anacrolixMetainfo.Piece
  completion *anacrolixStorage.Completion
}
func (pc *Piece) WriteAt(p []byte, off int64) (n int, err error) {
  file, err := os.OpenFile(pc.filename, os.O_WRONLY | os.O_CREATE , 0666)
  if err != nil {
    return 0 , err
  }
  defer file.Close()
  n, err = file.WriteAt(p, off)
  return n, err
}

func (pc *Piece) ReadAt(p []byte, off int64) (n int, err error) {
  file, err := os.OpenFile(pc.filename, os.O_RDONLY | os.O_CREATE, 0666)
  if err != nil {
    return 0, err
  }
  defer file.Close()
  n, err = file.ReadAt(p, off)
  return n, err
}

func (pc *Piece) MarkComplete() (err error) {
  pc.completion.Complete = true
  return nil
}

func (pc *Piece) MarkNotComplete() (err error) {
  pc.completion.Complete = false
  return nil
}

func (pc *Piece) Completion() anacrolixStorage.Completion {
  return *pc.completion
}
