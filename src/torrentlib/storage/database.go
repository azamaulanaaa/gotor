package storage

import (
	"sync"

	"github.com/spf13/afero"
)

type Database struct {
    file afero.File
    lock sync.RWMutex
}

func NewDatabase(file afero.File) *Database {
    return &Database{
        file: file,
    }
}

func (db *Database) Get(index int64) bool {
    db.lock.RLock()
    defer db.lock.RUnlock()

    position := index % 8
    section := (index - position) / 8

    stat, err := db.file.Stat()
    if err != nil {
        return false
    }

    if stat.Size() < section {
        return false
    }

    buffer := make([]byte, 1)
    db.file.ReadAt(buffer, section)

    expect := buffer[0] >> position & 1

    return expect == 1
}

func (db *Database) Set(index int64, value bool) {
    db.lock.Lock()
    defer db.lock.Unlock()

    position := index % 8
    section := (index - position) / 8

    buffer := make([]byte, 1)
    db.file.ReadAt(buffer, section)
    buffer[0] = buffer[0] | (1 << position)
    db.file.WriteAt(buffer, section)
    return
} 

func (db *Database) Close() error {
    return db.file.Close()
}
