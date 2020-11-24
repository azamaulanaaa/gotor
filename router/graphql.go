package router

import (
  "net/http"
  "github.com/azamaulanaaa/ohkaca/router/resolver"
  "github.com/azamaulanaaa/ohkaca/router/schema"

  gqlHandler "github.com/graphql-go/handler"
  anacrolixTorrent "github.com/anacrolix/torrent"
)

func newGQLHandler(client *anacrolixTorrent.Client) (http.Handler, error) {
	schemaGen := schema.New()
	schemaGen.AddQuery("files", resolver.Files(client))
	schemaGen.AddMutation("addMagnet", resolver.AddMagnet(client))
	gqlSchema, err := schemaGen.Generate()
  if err != nil {
    return nil, err
  }

  return gqlHandler.New(&gqlHandler.Config{
		Schema:   &gqlSchema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}

