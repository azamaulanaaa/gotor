package resolver

import (
	"path"

	anacrolixTorrent "github.com/anacrolix/torrent"
	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	gql "github.com/graphql-go/graphql"
)

func Files(client *anacrolixTorrent.Client) *gql.Field {
	return &gql.Field{
		Type:        gql.NewList(gqlFileO),
		Description: "List files of a torrent",
		Args: gql.FieldConfigArgument{
			"hash": &gql.ArgumentConfig{
				Type:        gql.NewNonNull(gql.String),
				Description: "Info hash of a torrent",
			},
		},
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			var hash anacrolixMetainfo.Hash
			err := hash.FromHexString(p.Args["hash"].(string))
			if err != nil {
				return nil, err
			}

			torrent, ok := client.Torrent(hash)
			if !ok {
				torrent, _ = client.AddTorrentInfoHash(hash)
			}
			<-torrent.GotInfo()

			files := []gqlFile{}
			torrentFiles := torrent.Files()
			for _, v := range torrentFiles {
				filename := v.DisplayPath()
				basename := path.Base(filename)
				url := "/" + hash.HexString() + "/" + filename
				files = append(files, gqlFile{
					Name: basename,
					URL:  url,
				})
			}

			return files, nil
		},
	}
}

type gqlFile struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var gqlFileO = gql.NewObject(gql.ObjectConfig{
	Name: "file",
	Fields: gql.Fields{
		"name": &gql.Field{
			Type:        gql.String,
			Description: "file name",
		},
		"url": &gql.Field{
			Type:        gql.String,
			Description: "url to stream files",
		},
	},
})
