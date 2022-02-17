package main

import (
	"log"
	"os"

	"github.com/azamaulanaaa/gotor/src/metainfo"
)

func main() {
    var metainfo_path string
    metainfo_path = "./examples/http.torrent"

    file, err := os.Open(metainfo_path)
    if err != nil {
        log.Panicln(err)
    }

    theMetainfo, err := metainfo.Decode(file)
    if err != nil {
        log.Panicln(err)
    }

    log.Println(theMetainfo.Raw())
}
