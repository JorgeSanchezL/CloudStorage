package filesystem

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
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

var (
	fs FileSystem
)

func InitializeFileSystem(path string) {
	data, err := ReadDirectory(path)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Error reading root directory")
		if os.IsNotExist(err) {
			err = os.Mkdir(path, 0755)
			if err != nil {
				panic(fmt.Sprintf("Error creating root directory: %s", err.Error()))
			}
			InitializeFileSystem(path)
			return
		}
	}
	fs = FileSystem{Root: *data} // check panic, create root if not exists
}

func ReadDirectory(path string) (*DirectoryData, error) {
	files := make(map[string]FileData, 0)
	subdirectories := make(map[string]DirectoryData, 0)
	size := int64(0)
	var lastModification time.Time

	entry, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, e := range entry {
		info, err := e.Info()
		if err != nil {
			log.Logger.Warn().Err(err).Msg("Error reading entry info")
		}
		if !info.IsDir() {
			files[e.Name()] = FileData{
				Size:             info.Size(),
				LastModification: info.ModTime(),
			}
		} else {
			data, _ := ReadDirectory(path + "/" + e.Name())
			subdirectories[e.Name()] = *data
		}

		size += info.Size()

		if info.ModTime().After(lastModification) {
			lastModification = info.ModTime()
		}
	}
	return &DirectoryData{Files: files, Subdirectories: subdirectories, Size: size, LastModification: lastModification}, nil
}

func GetDirectoryInfo(path string) DirectoryData {
	currentDirectory := fs.Root
	if path == "" {
		return currentDirectory
	}
	splittedPath := strings.Split(path, "/")
	for _, p := range splittedPath {
		currentDirectory = currentDirectory.Subdirectories[p]
	}
	return currentDirectory
}

func SaveFiles(c *gin.Context, formData *multipart.Form) {
	files := formData.File["files"]
	paths := formData.Value["paths"]
	filesInfo := make([]string, len(files))
	for file := range files {
		filesInfo = append(filesInfo, files[file].Filename)
	}

	for _, path := range paths {
		for _, file := range files {
			filePath := filepath.Join("root", path, file.Filename)
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "error", "errorMessage": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func DeleteFileOrEmptyDirectory(path string) error {
	splittedPath := strings.Split(path, "/")
	currentDirectory := fs.Root
	for index, p := range splittedPath {
		if index == len(splittedPath)-1 {
			err := deleteFileOrEmptyDirectory(filepath.Join("root", path))
			if err != nil {
				return err
			}
			delete(currentDirectory.Files, p)
			return nil
		}
		currentDirectory = currentDirectory.Subdirectories[p]
	}
	return fmt.Errorf("file %s not found", path)
}

func deleteFileOrEmptyDirectory(path string) error {
	// The Remove function deletes the named file or directory.
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDirectory(path string) error {
	directoryData := GetDirectoryInfo(path)

	// Delete all files in directory
	for file := range directoryData.Files {
		err := DeleteFileOrEmptyDirectory(filepath.Join(path, file))
		if err != nil {
			return err
		}
	}

	// Delete all subdirectories in directory
	for subdirectory := range directoryData.Subdirectories {
		err := DeleteDirectory(filepath.Join(path, subdirectory))
		if err != nil {
			return err
		}
	}

	// If it is the root, we have finished
	if path != "" {
		return nil
	}

	// If it is not the root, delete the directory itself
	err := DeleteFileOrEmptyDirectory(path)
	if err != nil {
		return err
	}

	return nil
}
