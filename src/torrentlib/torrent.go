package torrentlib

import (
    anacrolix "github.com/anacrolix/torrent"
)

type Torrent struct {
    anacrolixTorrent *anacrolix.Torrent
}

func (torrent *Torrent) Files() []*File {
    files := []*File{}

    anacrolixfiles := torrent.anacrolixTorrent.Files()
    for _, anacrolixfile := range anacrolixfiles {
        files = append(files, &File{
            anacrolixFile: anacrolixfile,
        })
    }

    return files
}
