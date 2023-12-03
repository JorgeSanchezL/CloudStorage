package filesystem

import (
	"os"
	"time"
)

type FileData struct {
	Size             int64
	LastModification time.Time
}

type DirectoryData struct {
	Files            map[string]FileData
	Subdirectories   map[string]DirectoryData
	Size             int64
	LastModification time.Time
}

type FileSystem struct {
	Root DirectoryData
}

func ReadFileSystem(path string) FileSystem {
	return FileSystem{Root: ReadDirectory(path)}
}

func ReadDirectory(path string) DirectoryData {
	files := make(map[string]FileData, 0)
	subdirectories := make(map[string]DirectoryData, 0)
	size := int64(0)
	var lastModification time.Time

	entry, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, e := range entry {
		info, err := e.Info()
		if err != nil {
			panic(err)
		}
		if !info.IsDir() {
			files[e.Name()] = FileData{
				Size:             info.Size(),
				LastModification: info.ModTime(),
			}
		} else {
			subdirectories[e.Name()] = ReadDirectory(path + "/" + e.Name())
		}

		size += info.Size()

		if info.ModTime().After(lastModification) {
			lastModification = info.ModTime()
		}
	}
	return DirectoryData{Files: files, Subdirectories: subdirectories, Size: size, LastModification: lastModification}
}

func DeleteFile(path string) error {
	// The Remove function deletes the named file or directory.
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
