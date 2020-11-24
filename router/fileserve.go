package router

import (
  "net/http"
  "path"
  "time"

  gorillaMux "github.com/gorilla/mux"
  anacrolixTorrent "github.com/anacrolix/torrent"
  anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
)

func newFileServe(client *anacrolixTorrent.Client) http.HandlerFunc {
  return func(rw http.ResponseWriter, r *http.Request) {
		params := gorillaMux.Vars(r)

		var hash anacrolixMetainfo.Hash
		if err := hash.FromHexString(params["hash"]); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

    torrent, _ := client.AddTorrentInfoHash(hash)
    <-torrent.GotInfo() 

		files := torrent.Files()
		for _, v := range files {
			if v.DisplayPath() == params["filename"] {
        basename := path.Base(v.DisplayPath())

        reader := v.NewReader()
        defer reader.Close()
        http.ServeContent(rw, r, basename, time.Unix(0,0), reader)
        return
      }
		}

		rw.WriteHeader(http.StatusBadRequest)
	}
}

