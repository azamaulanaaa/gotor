package storage

import (
    "io"
    "io/fs"
    "path/filepath"
    "time"

    "github.com/spf13/afero"
)

type CleanUp struct {
    filesystem  afero.Fs
    olderthan   time.Duration
    interval    time.Duration
}

func NewCleanUp(fileSystem afero.Fs, olderthan time.Duration, interval time.Duration) *CleanUp {
    return &CleanUp{
        filesystem: fileSystem,
        olderthan:  olderthan,
        interval:   interval,
    }
}

func (cleanUp *CleanUp) StartService() {
    go func(){
        for {
            time.Sleep(cleanUp.interval)
            afero.Walk(cleanUp.filesystem, "./", cleanUp.walkFunc)
        }
    }()
}

func (cleanUp *CleanUp) walkFunc(name string, info fs.FileInfo, err error) error {
    if err != nil {
        return nil
    }

    if name == "./" {
        return nil
    } else if info.IsDir() == true {
        dirInstance, err := cleanUp.filesystem.Open(name)
        if err != nil {
            return nil
        }

        _, err = dirInstance.Readdirnames(1)
        if err != io.EOF {
            return nil
        }
    } else if filepath.Ext(name) == ".db"{
        return nil
    } else if time.Now().Sub(info.ModTime()) < cleanUp.olderthan {
        return nil
    }

    cleanUp.filesystem.Remove(name)
    return nil
}
