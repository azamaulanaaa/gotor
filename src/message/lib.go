package message

import (
	"io"
)

func readUntil(r io.Reader, length int64) ([]byte, error) {
    buff := make([]byte, length)
    for {
        _, err := r.Read(buff)
        if err == io.EOF {
            continue
        }
        if err != nil {
            return buff, err
        }

        break
    }

    return buff, nil
}

type limitReader struct {
    limit   int64
    off     int64
    r       io.Reader
}

func newLimitReader(r io.Reader, limit int64) io.Reader {
    return &limitReader{
        limit: limit,
        off: 0,
        r: io.LimitReader(r, limit),
    }
}

func (r *limitReader) Read(b []byte) (int, error) {
    var nTotal int
    for {
        n, err := r.r.Read(b)

        nTotal = nTotal + n
        r.off = r.off + int64(n)

        if err != nil && err != io.EOF {
            return nTotal, err
        }
       
        if r.off == r.limit {
            return nTotal, io.EOF
        }

        if err == nil {
            return nTotal, nil
        }
    }
}

type readerReadAt struct {
    off     int64
    r       io.ReaderAt
}

func newReaderReadAt(r io.ReaderAt, off int64) io.Reader {
    return &readerReadAt{
        off: off,
        r: r,
    }
}

func (r *readerReadAt) Read(b []byte) (int, error) {
    n, err := r.r.ReadAt(b, r.off)
    r.off = r.off + int64(n)

    return n, err
}
