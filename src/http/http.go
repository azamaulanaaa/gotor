package http

import (
    "fmt"
    "net/http"

    "github.com/azamaulanaaa/gotor/src/torrentlib"
)

type HttpServer struct {
    torrentClient   *torrentlib.Client
    port            uint16
    httpHandler     http.Handler
}

func NewHttpServer(port uint16, torrentClient *torrentlib.Client) HttpServer {
    httpServer := HttpServer {
        torrentClient:  torrentClient,
        port:           port,
        httpHandler:    NewTorrentHttpHanndler(torrentClient),
    }

    return httpServer
}

func (httpServer HttpServer) ListenAndServe() error {
    host := fmt.Sprintf(":%d", httpServer.port)
    return http.ListenAndServe(host, httpServer.httpHandler)
}
