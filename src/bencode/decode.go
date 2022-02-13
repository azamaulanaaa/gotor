package bencode

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type parser func(r io.Reader) (interface{}, io.Reader, error)

var (
    ErrorInvalid = errors.New("data is not valid bencode")
    ErrorEOD = errors.New("bencode end of data")
)

func Decode(r io.Reader) (interface{}, error) {
    var err error
    var data interface{}

    if data, r, err = parseBencodeEOD(r); err == ErrorEOD {
        return nil, ErrorEOD
    }

    for _, theParser := range []parser{
        parseBencodeString,
        parseBencodeInteger,
        parseBencodeList,
        parseBencodeDictionary,
    } {
        data, r, err = theParser(r)
        if err == nil {
            return data, nil
        }
    }

    return nil, ErrorInvalid
}

func parseBencodeString(r io.Reader) (interface{}, io.Reader,  error) {
    var err error

    newR := bufio.NewReader(r)
    peek, err := newR.Peek(1)
    if err != nil {
        return nil, newR, ErrorInvalid
    }

    if peek[0] < 48 || peek[0] > 57 {
        return nil, newR, ErrorInvalid
    }

    lenAsStr, err := newR.ReadString(':') 
    if err != nil {
        return nil, newR, ErrorInvalid
    }

    lenAsStr = lenAsStr[:len(lenAsStr)-1]
    lenStr, err := strconv.ParseInt(lenAsStr, 10, 64)
    if err != nil {
        return nil, newR, ErrorInvalid
    }

    data, err := io.ReadAll(io.LimitReader(newR, lenStr))
    if err != nil {
        return nil, newR, ErrorInvalid
    }

    return String(data), newR, nil
}

func parseBencodeInteger(r io.Reader) (interface{}, io.Reader, error) {
    var err error

    newR := bufio.NewReader(r)
    peek, err := newR.Peek(1)
    if err != nil {
        fmt.Println(err)
        return nil, newR, ErrorInvalid
    }

    if peek[0] != 'i' {
        return nil, newR, ErrorInvalid
    }

    newR.Discard(1)
    dataStr, err := newR.ReadString('e') 
    dataStr = dataStr[:len(dataStr)-1]
    if err != nil {
        return nil, newR, ErrorInvalid
    }

    data, err := strconv.ParseInt(dataStr, 10, 64)
    if err != nil {
        return nil, newR, ErrorInvalid
    }

    return Integer(data), newR, nil
}

func parseBencodeEOD(r io.Reader) (interface{}, io.Reader, error) {
    var err error

    newR := bufio.NewReader(r)
    peek, err := newR.Peek(1)
    if err != nil {
        return nil, newR, ErrorInvalid
    }

    if peek[0] != 'e' {
        return nil, newR, ErrorInvalid
    }

    newR.Discard(1)
    return nil, newR, ErrorEOD
}
 

func parseBencodeList(r io.Reader) (interface{}, io.Reader, error) {
    var err error

    newR := bufio.NewReader(r)
    peek, err := newR.Peek(1)
    if err != nil {
        return nil, newR, ErrorInvalid
    }

    if peek[0] != 'l' {
        return nil, newR, ErrorInvalid
    }


    newR.Discard(1)
    data := List{}
    for {
        rawData, err := Decode(newR) 
        if err == ErrorEOD {
            return data, newR, nil
        }
        if err != nil {
            return nil, newR, err
        }

        data = append(data, rawData) 
    }
}

func parseBencodeDictionary(r io.Reader) (interface{}, io.Reader, error) {
    var err error

    newR := bufio.NewReader(r)
    peek, err := newR.Peek(1)
    if err != nil {
        return nil, newR, ErrorInvalid
    }

    if peek[0] != 'd' {
        return nil, newR, ErrorInvalid
    }

    newR.Discard(1)
    data := Dictionary{}
    for {
        key, err := Decode(newR)
        if err == ErrorEOD {
            return data, newR, nil
        }
        if err != nil {
            return nil, newR, err
        }

        keyStr, ok := key.(String)
        if !ok {
            return nil, newR, ErrorInvalid
        }

        value, err := Decode(newR)
        if err == ErrorEOD {
            return nil, newR, ErrorInvalid
        }
        if err != nil {
            return nil, newR, err
        }

        data[keyStr] = value
    }
}
