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
  // fmt.Printf("Write piece %d at %d with length %d\n", pc.pieceInfo.Index(), off, len(p))
  file, err := os.OpenFile(pc.filename, os.O_WRONLY | os.O_CREATE , 0666)
  if err != nil {
    // panic(err)
    return 0 , err
  }
  n, err = file.WriteAt(p, pc.pieceInfo.Offset() + off)
  file.Close()
  return n, err
}

func (pc *Piece) ReadAt(p []byte, off int64) (n int, err error) {
  file, err := os.OpenFile(pc.filename, os.O_RDONLY | os.O_CREATE, 0666)
  if err != nil {
    // panic(err)
    return 0, err
  }
  n, err = file.ReadAt(p, pc.pieceInfo.Offset() + off)
  file.Close()
  return n, err
}

func (pc *Piece) MarkComplete() (err error) {
  // fmt.Printf("Mark piece %d as complete\n", pc.pieceInfo.Index())
  pc.completion.Complete = true
  return nil
}

func (pc *Piece) MarkNotComplete() (err error) {
  pc.completion.Complete = false
  return nil
}

func (pc *Piece) Completion() anacrolixStorage.Completion {
  // fmt.Printf("Status piece %d %v\n", pc.pieceInfo.Index(), *pc.completion)
  return *pc.completion
}
