package storage

import (
	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
)

type Torrent struct {
  completionDB []*anacrolixStorage.Completion
  filename string
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


  piece := &Piece{
    filename : t.filename,
    pieceInfo: p,
    completion: completion,
  }
  return piece
}

func (t *Torrent) Close() (err error) {
  return nil
}
