package storage

import (
  "path/filepath"
  "os"

	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
)

type Storage struct {
  baseDir string
}

func New(baseDir string) *Storage {
	return &Storage{
      baseDir: baseDir,
	}
}

func (s *Storage) OpenTorrent(info *anacrolixMetainfo.Info, infohash anacrolixMetainfo.Hash) (anacrolixStorage.TorrentImpl, error) {
  dirname := filepath.Join(s.baseDir, infohash.HexString())
  err := os.MkdirAll(dirname, 0777)
  if err != nil {
    return nil, err
  }

	return &Torrent{
    completionDB : make([]*anacrolixStorage.Completion, info.NumPieces()),
	  dirname: dirname,
	}, nil
}
