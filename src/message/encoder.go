package message

import "github.com/azamaulanaaa/gotor/src/big_endian"

func EncodeMessage(rawMessage interface{}) ([]byte, error) {
    switch message := rawMessage.(type) {
    case KeepAlive:
        return encodeKeepAlive()
    case Choke:
        return encodeChoke(message)
    case UnChoke:
        return encodeUnChoke(message)
    case Interested:
        return encodeInterested(message)
    case NotInterested:
        return encodeNotInterested(message)
    case Have:
        return encodeHave(message)
    case Bitfield:
        return encodeBitfield(message)
    case Reqeust:
        return encodeRequest(message)
    case Piece:
        return encodePiece(message)
    case Cancel:
        return encodeCancel(message)
    default:
        return nil, ErrorMessageInvalid
    }
}

func EncodeHandshake(handshake Handshake) ([]byte, error) {
    var err error

    encodedProtocolLength, err := bigendian.Encode(uint8(len(handshake.Protocol)))
    if err != nil {
        return nil, err
    }

    rawData := []byte{}
    for _, theData := range [][]byte{
        encodedProtocolLength,
        []byte(handshake.Protocol),
        handshake.Reserved[:],
        handshake.Infohash[:],
        handshake.PeerID[:],
    } {
        rawData = append(rawData, theData...)
    }

    return rawData, nil
}

func encodeKeepAlive() ([]byte, error) {
    encodedLenMessage, err := bigendian.Encode(uint32(0))
    if err != nil {
        return nil, err
    }

    return encodedLenMessage, nil
}

func finisher(id messageID, data []byte) ([]byte, error) {
    encodedLenMessage, err := bigendian.Encode(uint32(len(data) + 1))
    if err != nil {
        return nil, err
    }

    rawData := []byte{}
    for _, theData := range [][]byte{
        encodedLenMessage,
        []byte{byte(id)},
        data,
    }{
        rawData = append(rawData, theData...)
    }
    
    return rawData, nil
}

func encodeChoke(message Choke) ([]byte, error) {
    return finisher(messageChoke, nil)
}

func encodeUnChoke(message UnChoke) ([]byte, error) {
    return finisher(messageChoke, nil)
}

func encodeInterested(message Interested) ([]byte, error) {
    return finisher(messageInterested, nil)
}

func encodeNotInterested(message NotInterested) ([]byte, error) {
    return finisher(messageNotInterested, nil)
}

func encodeHave(message Have) ([]byte, error) {
    var err error

    data := make([]byte, 0, 4)
    data, err = bigendian.Encode(message.Index)
    if err != nil {
        return nil, err
    }

    return finisher(messageHave, data)
}

func encodeBitfield(message Bitfield) ([]byte, error) {
    data := message.Bitfield.AsBytes()

    return finisher(messageBitfield, data)
}

func encodeRequest(message Reqeust) ([]byte, error) {
    var err error

    if message.Length > MaxLength {
        return nil, ErrorMessageTooLong
    }

    encodedIndex, err := bigendian.Encode(message.Index)
    if err != nil {
        return nil, err
    }

    encodedBegin, err := bigendian.Encode(message.Begin)
    if err != nil {
        return nil, err
    }

    encodedLength, err := bigendian.Encode(message.Length)
    if err != nil {
        return nil,  err
    }

    data := make([]byte, 0, 12)
    for _, theData := range [][]byte{
        encodedIndex,
        encodedBegin,
        encodedLength,
    } {
        data = append(data, theData...)
    }

    return finisher(messageRequest, data)
}

func encodePiece(message Piece) ([]byte, error) {
    var err error

    if len(message.Piece) > int(MaxLength) {
        return nil, ErrorMessageTooLong
    }

    encodedIndex, err := bigendian.Encode(message.Index)
    if err != nil {
        return nil, err
    }

    encodedBegin, err := bigendian.Encode(message.Begin)
    if err != nil {
        return nil, err
    }

    data := make([]byte, 0, 12)
    for _, theData := range [][]byte{
        encodedIndex,
        encodedBegin,
        message.Piece,
    } {
        data = append(data, theData...)
    }

    return finisher(messagePiece, data)
}

func encodeCancel(message Cancel) ([]byte, error) {
    var err error

    if message.Length > MaxLength {
        return nil, ErrorMessageTooLong
    }

    encodedIndex, err := bigendian.Encode(message.Index)
    if err != nil {
        return nil, err
    }

    encodedBegin, err := bigendian.Encode(message.Begin)
    if err != nil {
        return nil, err
    }

    encodedLength, err := bigendian.Encode(message.Length)
    if err != nil {
        return nil, err
    }

    data := make([]byte, 0, 12)
    for _, theData := range [][]byte{
        encodedIndex,
        encodedBegin,
        encodedLength,
    } {
        data = append(data, theData...)
    }

    return finisher(messageCancel, data)
}
