package filesystem

import (
	"errors"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Size             int64
	LastModification time.Time
}

type DirectoryData struct {
	Files            map[string]Entry
	Subdirectories   map[string]Entry
	Size             int64
	LastModification time.Time
}

type FileSystem struct {
	Root DirectoryData
}

var (
	fs FileSystem
)

var (
	ErrCannotCreateRootDirectory = errors.New("cannot create root directory")
)

func InitializeFileSystem(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			return errors.Join(ErrCannotCreateRootDirectory, err)
		}
	}
	return nil
}

func ReadDirectory(path string) (*DirectoryData, error) {
	files := make(map[string]Entry, 0)
	subdirectories := make(map[string]Entry, 0)
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
		entry := Entry{
			Size:             info.Size(),
			LastModification: info.ModTime(),
		}

		if !info.IsDir() {
			files[e.Name()] = entry
		} else {
			subdirectories[e.Name()] = entry
		}

		size += info.Size()

		if info.ModTime().After(lastModification) {
			lastModification = info.ModTime()
		}
	}
	return &DirectoryData{Files: files, Subdirectories: subdirectories, Size: size, LastModification: lastModification}, nil
}

func SaveFiles(c *gin.Context, formData *multipart.Form) {
	files := formData.File["files"]
	paths := formData.Value["paths"]

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
func DeleteEntry(path string) error {
	// The Remove function deletes the named file or directory.
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
