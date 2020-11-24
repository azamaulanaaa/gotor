package storage

import (
  "path/filepath"
  "os"
  "time"

	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
)

const cleanUpInterval = 5 * time.Minute

type Storage struct {
  baseDir string
}

func New(baseDir string) *Storage {
  go func(){
    for {
      time.Sleep(cleanUpInterval)
      err := Cleanup(baseDir, cleanUpInterval)
      if err != nil {
        panic(err)
      }
    }
  }()

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
  torrent := &Torrent{
    completionDB : make([]*anacrolixStorage.Completion, info.NumPieces()),
	  dirname: dirname,
	}

	return torrent, nil
}
