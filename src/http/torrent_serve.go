package http

import (
	"io"
	"net/http"
    "os"
	"path/filepath"
	"strings"
	"time"

	"github.com/azamaulanaaa/gotor/src/torrentlib"
)

type TorrentHttpHandler struct{
    torrentClient *torrentlib.Client
}

func NewTorrentHttpHanndler(client *torrentlib.Client) TorrentHttpHandler {
    return TorrentHttpHandler{
        torrentClient: client,
    }
}

func (torrentServe TorrentHttpHandler) Reader(hash string, path string) (io.ReadSeekCloser, error) {
    torrent, err := torrentServe.torrentClient.AddHash(hash)
    if err != nil {
        return nil, os.ErrNotExist
    }

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

/*
    Handle torrent file requiest with url path /[torrent hash]/[file path]
*/
func (torrentServe TorrentHttpHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    urlPath := r.URL.Path
    hash := strings.Split(urlPath,"/")[1]
    path, _ := filepath.Rel("/" + hash, r.URL.Path)
    basename := filepath.Base(path)
    reader, err := torrentServe.Reader(hash, path)
    if err != nil {
        rw.WriteHeader(http.StatusNotFound)
        return
    }
    defer reader.Close()

    http.ServeContent(rw, r, basename, time.Unix(0,0), reader)
    return
}
