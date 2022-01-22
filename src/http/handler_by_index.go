package http

import (
	"net/http"
    "path/filepath"
    "strings"
    "strconv"
    "time"

	"github.com/azamaulanaaa/gotor/src/torrentlib"
)

type TorrentHandlerByIndex struct {
    torrentClient   *torrentlib.Client
}

func NewTorrentHandlerByIndex(torrentClient *torrentlib.Client) http.Handler {
    return TorrentHandlerByIndex{
        torrentClient: torrentClient,
    }
}

func (torrentHandlerByIndex TorrentHandlerByIndex) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    var file *torrentlib.File
    {
        urlPath := r.URL.Path
        hash := strings.Split(urlPath,"/")[1]
        var index uint16
        {
            indexStr, _ := filepath.Rel("/" + hash, urlPath)
            indexInt, err := strconv.ParseInt(indexStr, 10, 16)
            if err != nil {
                rw.WriteHeader(http.StatusNotFound)
                return
            }
            index = uint16(indexInt)
        }

        {
            torrent, err := torrentHandlerByIndex.torrentClient.AddHash(hash)
            if err != nil {
                rw.WriteHeader(http.StatusNotFound)
                return
            }
            
            files := torrent.Files()
            if len(files) <= int(index) {
                rw.WriteHeader(http.StatusNotFound)
                return
            }

            file = files[index]
        }
    }

    basename := filepath.Base(file.Path())
    reader := file.Reader()
    defer reader.Close()

    http.ServeContent(rw, r, basename, time.Unix(0,0), reader)
    return
}
