package metainfo

import (
	"fmt"

	"github.com/azamaulanaaa/gotor/src"
	"github.com/azamaulanaaa/gotor/src/bencode"
)

func Raw(theMetainfo src.Metainfo) interface{} {
    if ourMetainfo, ok := theMetainfo.(metainfo); ok {
        return bencode.Dictionary(ourMetainfo)
    }

    out := bencode.Dictionary{}

    out["announce"] = bencode.String(theMetainfo.Announce())
    
    info := bencode.Dictionary{}
    {
        theInfo := theMetainfo.Info()
        info["piece length"] = bencode.Integer(theInfo.PieceLength())

        var pieces bencode.String
        {
            thePieces := theInfo.Pieces()
            for _, pieceHash := range thePieces {
                pieces = bencode.String(fmt.Sprintf("%s%s", pieces, string(pieceHash[:])))
            }
            info["pieces"] = pieces
        }

        if theName, ok := theInfo.Name(); ok {
            info["name"] = theName
        }

        if theLength, ok := theInfo.Length(); ok {
            info["length"] = bencode.Integer(theLength)
        }
        
        if theFiles, ok := theInfo.Files(); ok {
            files := bencode.List{}
            {
                for _, theFile := range theFiles {
                    file := bencode.Dictionary{}
                    file["length"] = bencode.Integer(theFile.Length())
                    file["path"] = bencode.String(theFile.Path())
                }
            }
            info["files"] = files
        }

        if thePrivate, ok := theInfo.Private(); ok {
            info["private"] = bencode.Integer(0)
            if thePrivate == true {
                info["private"] = bencode.Integer(1)
            }
        }
    }
    out["info"] = info
    return out 
}
