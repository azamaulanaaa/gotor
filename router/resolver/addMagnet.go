package resolver

import (
	anacrolixTorrent "github.com/anacrolix/torrent"
	gql "github.com/graphql-go/graphql"
)

func AddMagnet(client *anacrolixTorrent.Client) *gql.Field {
	return &gql.Field{
		Type:        gql.String,
		Description: "Add torrent using Magnet URI and return info hash",
		Args: gql.FieldConfigArgument{
			"uri": &gql.ArgumentConfig{
				Type:        gql.NewNonNull(gql.String),
				Description: "Magnet URI",
			},
		},
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			torrent, err := client.AddMagnet(p.Args["uri"].(string))
			if err != nil {
				return nil, err
			}

			<-torrent.GotInfo()
			return torrent.InfoHash().HexString(), nil
		},
	}
}
