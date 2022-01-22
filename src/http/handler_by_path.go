package http

import (
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/azamaulanaaa/gotor/src/http/lib"
	"github.com/azamaulanaaa/gotor/src/torrentlib"
)

type TorrentHandlerByPath struct{
    torrentClient *torrentlib.Client
}

func NewTorrentHandlerByPath(client *torrentlib.Client) http.Handler {
    return TorrentHandlerByPath{
        torrentClient: client,
    }
}

/*
    Handle torrent file requiest with url path /[torrent hash]/[file path]
*/
func (torrentHandlerByPath TorrentHandlerByPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    urlPath := r.URL.Path
    hash := strings.Split(urlPath,"/")[1]
    path, _ := filepath.Rel("/" + hash, r.URL.Path)
    basename := filepath.Base(path)

    torrent, err := torrentHandlerByPath.torrentClient.AddHash(hash)
    if err != nil {
        rw.WriteHeader(http.StatusNotFound)
        return
    }

    reader, err := lib.TorrentReaderByPath(&torrent, path)
    if err != nil {
        rw.WriteHeader(http.StatusNotFound)
        return
    }
    defer reader.Close()

    http.ServeContent(rw, r, basename, time.Unix(0,0), reader)
    return
}
