package messagehandler

import "github.com/azamaulanaaa/gotor/src/big_endian"

func encodeMessage(rawMessage interface{}) ([]byte, error) {
    switch message := rawMessage.(type) {
    case MessageChoke:
        return encodeChoke(message)
    case MessageUnChoke:
        return encodeUnChoke(message)
    case MessageInterested:
        return encodeInterested(message)
    case MessageNotInterested:
        return encodeNotInterested(message)
    case MessageHave:
        return encodeHave(message)
    case MessageBitfield:
        return encodeBitfield(message)
    case MessageRequest:
        return encodeRequest(message)
    case MessagePiece:
        return encodePiece(message)
    case MessageCancel:
        return encodeCancel(message)
    default:
        return nil, ErrorMessageInvalid
    }
}

func encodeHandshake(handshake Handshake) ([]byte, error) {
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

func encodeChoke(message MessageChoke) ([]byte, error) {
    return finisher(messageChoke, nil)
}

func encodeUnChoke(message MessageUnChoke) ([]byte, error) {
    return finisher(messageChoke, nil)
}

func encodeInterested(message MessageInterested) ([]byte, error) {
    return finisher(messageInterested, nil)
}

func encodeNotInterested(message MessageNotInterested) ([]byte, error) {
    return finisher(messageNotInterested, nil)
}

func encodeHave(message MessageHave) ([]byte, error) {
    var err error

    data := make([]byte, 0, 4)
    data, err = bigendian.Encode(message.Index)
    if err != nil {
        return nil, err
    }

    return finisher(messageHave, data)
}

func encodeBitfield(message MessageBitfield) ([]byte, error) {
    data := message.Bitfield.AsBytes()

    return finisher(messageBitfield, data)
}

func encodeRequest(message MessageRequest) ([]byte, error) {
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

func encodePiece(message MessagePiece) ([]byte, error) {
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

func encodeCancel(message MessageCancel) ([]byte, error) {
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
