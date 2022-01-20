package storage

import (
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestDatabaseGet(t *testing.T) {
    filename := "test"

    fileSystem := afero.NewOsFs()
    databaseFile, _ := fileSystem.OpenFile(
        filename,
        os.O_CREATE | os.O_RDWR,
        0640,
    )
    defer fileSystem.Remove(filename)

    databaseFile.WriteAt([]byte{byte(3)},0)
    database := NewDatabase(databaseFile)
    defer database.Close()

    got := database.Get(1) 
    if got != true {
        t.Errorf("expect\t: %t\ngot\t: %t", true, false)
    }
}

func TestDatabaseSet(t *testing.T) {
    fileSystem := afero.NewMemMapFs()
    databaseFile, _ := fileSystem.OpenFile(
        "test", 
        os.O_CREATE | os.O_RDWR,
        0640,
    )
    database := NewDatabase(databaseFile)
    defer database.Close()
    database.Set(0, true) 

    buffer := make([]byte, 1)
    databaseFile.ReadAt(buffer, 0) 

    if buffer[0] != byte(1) {
        t.Errorf("expect\t: %t\ngot\t: %t", true, false)
    }
}
