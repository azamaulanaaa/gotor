package torrentlib

import (
    "io"

    anacrolix "github.com/anacrolix/torrent"
)

type File struct {
    anacrolixFile *anacrolix.File
}

func (file *File) Path() string {
    path := file.anacrolixFile.DisplayPath()
    return path
}

func (file *File) Reader() io.ReadSeekCloser {
    file.Validate()
    return file.anacrolixFile.NewReader()
}

func (file *File) Validate() {
    dataOffset := file.anacrolixFile.Offset()
    dataLength := file.anacrolixFile.Length()
    pieceSize := file.anacrolixFile.Torrent().Info().PieceLength
    startPiece := dataOffset / pieceSize
    endPiece := (dataOffset + dataLength) / pieceSize

    for index := startPiece; index <= endPiece; index ++ {
        piece := file.anacrolixFile.Torrent().Piece(int(index)) 
        if piece.Storage().Completion().Ok == false {
            piece.Storage().MarkNotComplete()
            piece.UpdateCompletion()
        }
    }
}
