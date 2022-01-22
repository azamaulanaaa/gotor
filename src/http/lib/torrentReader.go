package lib

import (
    "io"
	"os"

	"github.com/azamaulanaaa/gotor/src/torrentlib"
)

func TorrentReaderByPath(torrent *torrentlib.Torrent, path string) (io.ReadSeekCloser, error) {
    files := torrent.Files()
    for _, file := range files {
        if file.Path() != path {
            continue
        }
        
        reader := file.Reader()
        return reader, nil
    }
    
    return nil, os.ErrNotExist
}
