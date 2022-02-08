package metainfo

import (
	"bytes"
	"errors"
	"io"

	"github.com/azamaulanaaa/gotor/src"
	"github.com/azamaulanaaa/gotor/src/bencode"
)

var (
    ErrorInvalidMetainfo = errors.New("invalid metainfo data")
)

func Decode(r io.Reader) (src.Metainfo, error) {
    var err error

    var rawMetainfo bencode.Dictionary
    {
        var rawData interface{}
        rawData, err = bencode.Decode(r)
        if err != nil {
            return src.Metainfo{}, err
        }

        var ok bool
        rawMetainfo, ok = rawData.(bencode.Dictionary)
        if !ok {
            return src.Metainfo{}, ErrorInvalidMetainfo
        }
    }

    var metainfo src.Metainfo

    if rawAnnounce, ok := rawMetainfo["announce"].(bencode.String); ok {
        metainfo.Announce = string(rawAnnounce)
    }

    if rawInfo, ok := rawMetainfo["info"].(bencode.Dictionary); ok {
        var info src.Info
        
        if rawPieceLength, ok := rawInfo["piece length"].(bencode.Integer); ok {
            info.PieceLength = uint64(rawPieceLength)
        }

        if rawPieces, ok := rawInfo["pieces"].(bencode.String); ok {
            pieces := []src.Hash{}

            piecesBuffer := bytes.NewBuffer([]byte(rawPieces))
            for {
                var buff src.Hash
                _, err := piecesBuffer.Read(buff[:])
                if err == io.EOF {
                    break
                }
                if err != nil {
                    return src.Metainfo{}, ErrorInvalidMetainfo
                }
                pieces = append(pieces, buff)
            }

            info.Pieces = pieces
        }

        if rawName, ok := rawInfo["name"].(bencode.String); ok {
            info.Name = string(rawName)
        }

        if rawLength, ok := rawInfo["length"].(bencode.Integer); ok {
            info.Length = uint64(rawLength)
        }

        if rawFiles, ok := rawInfo["files"].(bencode.List); ok {
            files := []src.File{}
            
            for _, v := range rawFiles {
                if rawFile, ok := v.(bencode.Dictionary); ok { 
                    var file src.File

                    if rawLength, ok := rawFile["length"].(bencode.Integer); ok {
                        file.Length = uint64(rawLength)
                    }

                    if rawPath, ok := rawFile["path"].(bencode.String); ok {
                        file.Path = string(rawPath)
                    }

                    files = append(files, file)
                }
            }

            info.Files = files
        }

        metainfo.Info = info
    }

    return metainfo, nil
}
