package connection

import (
	"net"

	"github.com/azamaulanaaa/gotor/src"
)

type Connection interface {
    Listening(callback Callback)
    GetHandshake() <-chan src.Handshake
    SendHandshake(handshake src.Handshake) error
    SendMessage(message src.Message) error
    Close() error
}

type Callback func(message src.Message)

type connection_impl struct {
    conn        net.Conn
    stop        chan bool
}

func NewConnection(conn net.Conn) Connection {
    connection := connection_impl{
        conn: conn,
    }
    return connection
}

func (connection connection_impl) Close() error {
    connection.stop <- true
    return connection.conn.Close()
}

func (connection connection_impl) Listening(callback Callback) {
    go func(){
        for {
            select {
            case <-connection.stop:
                return
            default:
                message, err := parseMessage(connection.conn)
                if err != nil {
                    continue
                }
                
                go callback(message)
            }
        }
    }()
}

func (connection connection_impl) GetHandshake() <-chan src.Handshake {
    chanHandshake := make(chan src.Handshake)
    go func() {
        handshake, err := parseHandshake(connection.conn)
        if err != nil {
            chanHandshake <- src.Handshake{}
            return
        }
        chanHandshake <- handshake
        return
    }()
    return chanHandshake
}

func (connection connection_impl) SendHandshake(handshake src.Handshake) (error) {
    var err error

    pstrln, _ := EncodeBigEndian(uint8(len(handshake.Protocol)))
    var payload []byte
    for _, v := range [][]byte{
        pstrln,
        handshake.Protocol,
        handshake.Reserved[:],
        handshake.InfoHash[:],
        handshake.PeerID[:],
    } {
        payload = append(payload, v...)
    }

    _, err = connection.conn.Write(payload)
    if err != nil {
        return err
    }
    
    return nil
}

func (connection connection_impl) SendMessage(message src.Message) error {
    var err error

    var lenMessage [4]byte
    {
        sliceOfByteLenRawPayload, _ := EncodeBigEndian(uint32(len(message.Payload) + 1))
        copy(lenMessage[:], sliceOfByteLenRawPayload)
    }

    rawPayload := []byte{}
    for _, value := range [][]byte{
        lenMessage[:],
        []byte{byte(message.MessageID)},
        message.Payload,
    }{
        rawPayload = append(rawPayload, value...)
    }
    _, err = connection.conn.Write(rawPayload)
    if err != nil {
        return err
    }

    return nil
}
