package http

import (
    "fmt"
    "net/http"
    "path/filepath"
    "time"

    "github.com/azamaulanaaa/gotor/src/torrentlib"

    gorilla "github.com/gorilla/mux"
)

type HttpServer struct {
    torrentClient   *torrentlib.Client
    port            uint16
    router          *gorilla.Router
}

func NewHttpServer(port uint16, torrentClient *torrentlib.Client) HttpServer {
    httpServer := HttpServer {
        torrentClient:  torrentClient,
        port:           port,
        router:         gorilla.NewRouter(),
    }

    httpServer.router.PathPrefix("/{hash:[\\w\\d]+}/").HandlerFunc(httpServer.torrentServe)

    return httpServer
}

func (httpServer HttpServer) ListenAndServe() error {
    host := fmt.Sprintf(":%d", httpServer.port)
    return http.ListenAndServe(host, httpServer.router)
}

func (httpServer HttpServer) torrentServe(rw http.ResponseWriter, r *http.Request) {
    hash := gorilla.Vars(r)["hash"]
    path, _ := filepath.Rel("/" + hash, r.URL.Path)
    
    torrent, err := httpServer.torrentClient.AddHash(hash)
    if err != nil {
        rw.WriteHeader(http.StatusNotFound)
        return 
    }

    files := torrent.Files()
    for _, file := range files {
        if file.Path() != path {
            continue
        }
        
        reader := file.Reader()
        defer reader.Close()
        basename := filepath.Base(file.Path())
        http.ServeContent(rw, r, basename, time.Unix(0,0), reader)

        return
    }

    rw.WriteHeader(http.StatusNotFound)
    return
}
