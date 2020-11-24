package router

import (
  "net/http"

	gorillaMux "github.com/gorilla/mux"
  anacrolixTorrent "github.com/anacrolix/torrent"
)

func New(client *anacrolixTorrent.Client) (http.Handler, error) {
	router := gorillaMux.NewRouter()

  graphqlHandler, err := newGQLHandler(client)
  if err != nil {
    return nil, err
  }

	router.Handle("/graphql", graphqlHandler)
	router.HandleFunc("/{hash:[\\w\\d]+}/{filename}", newFileServe(client))
  router.PathPrefix("/").Handler(http.FileServer(http.Dir("./router/html")))
  return router, nil
}
