package metainfo

import (
	"bytes"
	"io"

	"github.com/azamaulanaaa/gotor/src/bencode"
    "github.com/azamaulanaaa/gotor/src/hash"
)

type metainfo_impl bencode.Dictionary
type info_impl bencode.Dictionary
type file_impl bencode.Dictionary

func (metainfo metainfo_impl) Announce() string {
    if rawAnnounce, ok := metainfo["announce"].(bencode.String); ok {
        return string(rawAnnounce)
    }

    return ""
}

func (metainfo metainfo_impl) Info() Info {
    if rawInfo, ok := metainfo["info"].(bencode.Dictionary); ok {
        return info_impl(rawInfo)
    }

    return nil
}

func (info info_impl) Name() (string, bool) {
    if rawName, ok := info["name"].(bencode.String); ok {
        return string(rawName), true
    }

    return "", false
}

func (info info_impl) PieceLength() uint64 {
    if rawPieceLength, ok := info["piece length"].(bencode.Integer); ok {
        return uint64(rawPieceLength)
    }

    return 0
}

func (info info_impl) Pieces() []hash.Hash {
    if rawPieces, ok := info["pieces"].(bencode.String); ok {
        pieces := []hash.Hash{}

        piecesBuffer := bytes.NewBuffer([]byte(rawPieces))
        for {
            var buff hash.Hash
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

    return []hash.Hash{}
}

func (info info_impl) Length() (uint64, bool) {
    if rawLength, ok := info["length"].(bencode.Integer); ok {
        return uint64(rawLength), true
    }
    
    return 0, false
}

func (info info_impl) Files() ([]File, bool) {
    if rawFiles, ok := info["files"].(bencode.List); ok {
        files := make([]File, 0, len(rawFiles))

        for _, v := range rawFiles {
            if rawFile, ok := v.(bencode.Dictionary); ok { 
                files = append(files, file_impl(rawFile))
            }
        }

        return files, true
    }

    return nil, false
}

func (info info_impl) Private() (bool, bool) {
    if rawPrivate, ok := info["private"].(bencode.Integer); ok {
        if rawPrivate == 1 {
            return true, true
        }else if rawPrivate == 0 {
            return false, true
        }
    }

    return false, false
}

func (file file_impl) Length() uint64 {
    if rawLength, ok := file["length"].(bencode.Integer); ok {
        return uint64(rawLength)
    }

    return 0
}

func (file file_impl) Path() string {
    if rawPath, ok := file["path"].(bencode.String); ok {
        return string(rawPath)
    }

    return ""
}
