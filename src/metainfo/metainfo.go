package metainfo

import (
	"bytes"
	"io"

	"github.com/azamaulanaaa/gotor/src"
	"github.com/azamaulanaaa/gotor/src/bencode"
)

type metainfo bencode.Dictionary
type info bencode.Dictionary
type file bencode.Dictionary

func (self metainfo) Announce() string {
    if rawAnnounce, ok := self["announce"].(bencode.String); ok {
        return string(rawAnnounce)
    }

    return ""
}

func (self metainfo) Info() src.MetainfoInfo {
    if rawInfo, ok := self["info"].(bencode.Dictionary); ok {
        return info(rawInfo)
    }

    return nil
}

func (self info) Name() (string, bool) {
    if rawName, ok := self["name"].(bencode.String); ok {
        return string(rawName), true
    }

    return "", false
}

func (self info) PieceLength() uint32 {
    if rawPieceLength, ok := self["piece length"].(bencode.Integer); ok {
        return uint32(rawPieceLength)
    }

    return 0
}

func (self info) Pieces() []src.Hash {
    if rawPieces, ok := self["pieces"].(bencode.String); ok {
        pieces := []src.Hash{}

        piecesBuffer := bytes.NewBuffer([]byte(rawPieces))
        for {
            var buff src.Hash
            _, err := piecesBuffer.Read(buff[:])
            if err == io.EOF {
                break
            }
            if err != nil {
                continue
            }
            pieces = append(pieces, buff)
        }

        return pieces
    }

    return []src.Hash{}
}

func (self info) Length() (uint64, bool) {
    if rawLength, ok := self["length"].(bencode.Integer); ok {
        return uint64(rawLength), true
    }
    
    return 0, false
}

func (self info) Files() ([]src.MetainfoFile, bool) {
    if rawFiles, ok := self["files"].(bencode.List); ok {
        files := make([]src.MetainfoFile, 0, len(rawFiles))

        for _, v := range rawFiles {
            if rawFile, ok := v.(bencode.Dictionary); ok { 
                files = append(files, file(rawFile))
            }
        }

        return files, true
    }

    return nil, false
}

func (self info) Private() (bool, bool) {
    if rawPrivate, ok := self["private"].(bencode.Integer); ok {
        if rawPrivate == 1 {
            return true, true
        }else if rawPrivate == 0 {
            return false, true
        }
    }

    return false, false
}

func (self file) Length() uint64 {
    if rawLength, ok := self["length"].(bencode.Integer); ok {
        return uint64(rawLength)
    }

    return 0
}

func (self file) Path() string {
    if rawPath, ok := self["path"].(bencode.String); ok {
        return string(rawPath)
    }

    return ""
}
