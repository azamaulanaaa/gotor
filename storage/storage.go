package storage

import (
  "path/filepath"

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
	return &Torrent{
    completionDB : make([]*anacrolixStorage.Completion, info.NumPieces()),
	  filename: filepath.Join(s.baseDir, infohash.HexString()),
	}, nil
}
