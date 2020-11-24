package main

import (
  "fmt"
	"net/http"
	"ohkaca/lib/schema"
	"ohkaca/resolver"
	"ohkaca/storage"
	"path"
	"time"

	anacrolixTorrent "github.com/anacrolix/torrent"
	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	gorillaMux "github.com/gorilla/mux"
	gqlHandler "github.com/graphql-go/handler"
)

const host = "localhost:3000"

func main() {
	customConfig := anacrolixTorrent.NewDefaultClientConfig()
	customConfig.DefaultStorage = storage.New("./tmp")
	customConfig.NoUpload = true
	client, err := anacrolixTorrent.NewClient(customConfig)
	must(err)
	defer client.Close()

	schemaGen := schema.New()
	schemaGen.AddQuery("files", resolver.Files(client))
	schemaGen.AddMutation("addMagnet", resolver.AddMagnet(client))
	gqlSchema, err := schemaGen.Generate()
	must(err)

	router := gorillaMux.NewRouter()

	router.Handle("/graphql", gqlHandler.New(&gqlHandler.Config{
		Schema:   &gqlSchema,
		Pretty:   true,
		GraphiQL: true,
	}))

	router.HandleFunc("/{hash:[\\w\\d]+}/{filename}", func(rw http.ResponseWriter, r *http.Request) {
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
        fmt.Println(basename)

        reader := v.NewReader()
        defer reader.Close()
        http.ServeContent(rw, r, basename, time.Unix(0,0), reader)
        return
      }
		}

		rw.WriteHeader(http.StatusBadRequest)
	})

	must(http.ListenAndServe(host, router))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
