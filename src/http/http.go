package http

import (
	"fmt"
	"net/http"

	"github.com/azamaulanaaa/gotor/src/torrentlib"
	gorilla "github.com/gorilla/mux"
)

type HttpServer struct {
    torrentClient   *torrentlib.Client
    port            uint16
    router          http.Handler
}

func NewHttpServer(port uint16, torrentClient *torrentlib.Client) HttpServer {
    router := gorilla.NewRouter()

    serveByIndex := NewTorrentHandlerByIndex(torrentClient)
    router.PathPrefix("/index/{\\w\\d}/").Handler(http.StripPrefix("/index", serveByIndex))

    serveByPath := NewTorrentHandlerByPath(torrentClient)
    router.PathPrefix("/{\\w\\d}/").Handler(serveByPath)

    return HttpServer {
        torrentClient:  torrentClient,
        port:           port,
        router:         router,
    }
}

func (httpServer HttpServer) ListenAndServe() error {
    host := fmt.Sprintf(":%d", httpServer.port)
    return http.ListenAndServe(host, httpServer.router)
}
