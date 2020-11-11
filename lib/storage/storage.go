package storage

import (
	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
)

type Storage struct {
	torrents map[string]*Torrent
}

func (s *Storage) OpenTorrent(info *anacrolixMetainfo.Info, infohash anacrolixMetainfo.Hash) (anacrolixStorage.TorrentImpl, error) {
	torrent := &Torrent{
		pieces: make([]*PieceBuffer, info.NumPieces()),
	}
	s.torrents[infohash.AsString()] = torrent
	return torrent, nil
}

type Torrent struct {
	pieces []*PieceBuffer
}

func (t *Torrent) Piece(p anacrolixMetainfo.Piece) anacrolixStorage.PieceImpl {
	piece := &PieceBuffer{
		buff:   make([]byte, p.Length()),
		status: anacrolixStorage.Completion{},
	}
	t.pieces[p.Index()] = piece
	return piece
}

func (t *Torrent) Close() (err error) {
	t.pieces = nil
	return
}

type PieceBuffer struct {
	buff   []byte
	status anacrolixStorage.Completion
}

func (pb *PieceBuffer) WriteAt(p []byte, off int64) (n int, err error) {
	n = len(p)

	p_len := int64(len(p))
	new_buff := append(pb.buff[off:], p...)
	new_buff = append(new_buff, pb.buff[off+p_len:]...)
	pb.buff = new_buff
	return
}

func (pb *PieceBuffer) ReadAt(p []byte, off int64) (n int, err error) {
	n = len(p)
	p = pb.buff[off : off+int64(n)]
	if pb.status.Complete {
		pb.buff = make([]byte, len(pb.buff))
		pb.MarkNotComplete()
	}
	return
}

func (pb *PieceBuffer) MarkComplete() (err error) {
	pb.status.Complete = true
	pb.status.Ok = true
	return
}

func (pb *PieceBuffer) MarkNotComplete() (err error) {
	pb.status.Complete = false
	pb.status.Ok = false
	return
}

func (pb *PieceBuffer) Completion() anacrolixStorage.Completion {
	return pb.status
}
