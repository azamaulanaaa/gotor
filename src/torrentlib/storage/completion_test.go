package storage

import (
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestCompletionGet(t *testing.T) {
    fileSystem := afero.NewMemMapFs()
    databaseFile, _ := fileSystem.OpenFile(
        "test", 
        os.O_CREATE | os.O_RDWR,
        0640,
    )
    databaseFile.WriteAt([]byte{byte(3)},0)
    database := NewDatabase(databaseFile)
    defer database.Close()
    database.Set(1, true) 

    completion := NewCompletion(database, 1)
    got := completion.Get() 
    if got != true {
        t.Errorf("expect\t: %t\ngot\t: %t", true, false)
    }
}

func TestCompletionSet(t *testing.T) {
    index := 10

    fileSystem := afero.NewMemMapFs()
    databaseFile, _ := fileSystem.OpenFile(
        "test", 
        os.O_CREATE | os.O_RDWR,
        0640,
    )
    database := NewDatabase(databaseFile)
    defer database.Close()

    completion := NewCompletion(database, index)
    completion.Set(true) 

    got := database.Get(int64(index))
    if got != true {
        t.Errorf("expect\t: %t\ngot\t: %t", true, false)
    }
}

func TestCompletionConsistency(t *testing.T) {
    index := 0

    fileSystem := afero.NewMemMapFs()
    databaseFile, _ := fileSystem.OpenFile(
        "test", 
        os.O_CREATE | os.O_RDWR,
        0640,
    )
    database := NewDatabase(databaseFile)
    defer database.Close()

    completion := NewCompletion(database, index)
    completion.Set(true) 

    got := completion.Get()
    if got != true {
        t.Errorf("expect\t: %t\ngot\t: %t", true, false)
    }
}
