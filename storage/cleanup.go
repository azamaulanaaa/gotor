package storage

import (
  "os"
  "path/filepath"
  "time"
)

func Cleanup(dirname string, olderthan time.Duration) (error) {
  err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }

    if !info.IsDir() && info.ModTime().Add(olderthan).Before(time.Now()) {
      return os.Remove(path)
    }
  
    return nil
  })
  return err
}
