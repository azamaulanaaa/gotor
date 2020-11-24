package storage

import (
  "path/filepath"

	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
)

type Torrent struct {
  completionDB []*anacrolixStorage.Completion
  dirname string
}

func (t *Torrent) Piece(p anacrolixMetainfo.Piece) anacrolixStorage.PieceImpl {
  completion := t.completionDB[p.Index()]
  if completion == nil {
    t.completionDB[p.Index()] = &anacrolixStorage.Completion{
      Complete: false,
      Ok : true, 
    }
    completion = t.completionDB[p.Index()]
  }

  filename := filepath.Join(t.dirname, p.Hash().HexString())

  piece := &Piece{
    filename : filename,
    pieceInfo: p,
    completion: completion,
  }
  return piece
}

func (t *Torrent) Close() (err error) {
  return nil
}
