package storage

import (
	"os"
	"sync"

	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
)

type Storage struct {
	torrent *Torrent
}

func New() *Storage {
	return &Storage{
		torrent: &Torrent{
			pieces: make(map[string]*Piece),
		},
	}
}

func (s *Storage) OpenTorrent(info *anacrolixMetainfo.Info, infohash anacrolixMetainfo.Hash) (anacrolixStorage.TorrentImpl, error) {
	return s.torrent, nil
}

type Torrent struct {
	pieces map[string]*Piece
}

func (t *Torrent) Piece(p anacrolixMetainfo.Piece) anacrolixStorage.PieceImpl {
	hash := p.Hash().HexString()
	if piece, ok := t.pieces[hash]; ok {
		return piece
	}
	piece := &Piece{
		filepath:   hash,
		completion: anacrolixStorage.Completion{},
	}
	t.pieces[hash] = piece
	return piece
}

func (t *Torrent) Close() (err error) {
	for _, v := range t.pieces {
		if err := v.Close(); err != nil {
			return err
		}
	}
	return
}

type Piece struct {
	mux        sync.RWMutex
	file       *os.File
	filepath   string
	active     int
	completion anacrolixStorage.Completion
}

func (pc *Piece) wakeup() (err error) {
	if pc.active == 0 {
		pc.file, err = os.OpenFile(pc.filepath, os.O_RDWR|os.O_CREATE, 666)
		if err != nil {
			return err
		}
	}
	pc.active = pc.active + 1
	return
}

func (pc *Piece) sleep() {
	pc.active = pc.active - 1
	if pc.active == 0 {
		pc.Close()
	}
}

func (pc *Piece) lock() (err error) {
	err = pc.wakeup()
	if err != nil {
		return
	}
	pc.mux.Lock()
	return
}

func (pc *Piece) unlock() {
	pc.sleep()
	pc.mux.Unlock()
}

func (pc *Piece) rlock() (err error) {
	err = pc.wakeup()
	if err != nil {
		return
	}
	pc.mux.RLock()
	return
}

func (pc *Piece) runlock() {
	pc.sleep()
	pc.mux.Unlock()
}

func (pc *Piece) WriteAt(p []byte, off int64) (n int, err error) {
	err = pc.lock()
	if err != nil {
		return
	}
	defer pc.unlock()
	return pc.WriteAt(p, off)
}

func (pc *Piece) ReadAt(p []byte, off int64) (n int, err error) {
	err = pc.rlock()
	if err != nil {
		return
	}
	defer pc.runlock()
	return pc.file.ReadAt(p, off)
}

func (pc *Piece) Close() (err error) {
	return pc.file.Close()
}

func (pc *Piece) MarkComplete() (err error) {
	pc.completion.Complete = true
	pc.completion.Ok = true
	return
}

func (pc *Piece) MarkNotComplete() (err error) {
	pc.completion.Complete = false
	pc.completion.Ok = false
	return
}

func (pc *Piece) Completion() anacrolixStorage.Completion {
	return pc.completion
}
