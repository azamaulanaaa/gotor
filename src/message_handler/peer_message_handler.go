package messagehandler

import (
	"fmt"
	"net"
	"time"

	"github.com/azamaulanaaa/gotor/src/peer"
)

type peerMessageHandler struct {
    conn                net.Conn
    amMessageHandler    MessageHandler
    stop                chan bool
}

type PeerMessageHandlerConfig struct {
    Timeout time.Duration
}

func NewPeerMessageHandler(peer peer.Peer, amMessageHandler MessageHandler, config PeerMessageHandlerConfig) (MessageHandler, error) {
    messageHandler := peerMessageHandler{
        amMessageHandler: amMessageHandler,
    }
    {
        host := fmt.Sprintf("%s:%d", peer.IP().String(), peer.Port())
        netconn, err := net.DialTimeout("tcp", host, config.Timeout)
        if err != nil {
            return nil, err
        }
        messageHandler.conn = netconn
    }

    go messageHandler.listening()

    return &messageHandler, nil
}

func (messageHandler *peerMessageHandler) listening() {
    select {
    case <-messageHandler.stop:
        return
    default:
        handshake, err := decodeHandshake(messageHandler.conn)
        if err != nil {
            return
        }

        err = messageHandler.amMessageHandler.SendHandshake(handshake)
        if err != nil {
            return
        }

        for {
            message, err := decodeMessage(messageHandler.conn)
            if err != nil {
                continue
            }

            messageHandler.amMessageHandler.SendMessage(message)
        }
    }
}

func (messageHandler *peerMessageHandler) SendHandshake(handshake Handshake) error {
    var err error

    rawData, err := encodeHandshake(handshake)
    if err != nil {
        return err
    }

    _, err = messageHandler.conn.Write(rawData)
    if err != nil {
        return err
    }

    return nil
}

func (messageHandler *peerMessageHandler) SendMessage(message interface{}) error {
    var err error

    rawData, err := encodeMessage(message)
    if err != nil {
        return err
    }

    _, err = messageHandler.conn.Write(rawData)
    if err != nil {
        return err
    }

    return nil
}

func (messageHandler *peerMessageHandler) Close() error {
    return messageHandler.conn.Close()
}
