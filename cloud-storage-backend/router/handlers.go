package router

import (
	"cloud-storage-backend/filesystem"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func handleGetFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Query("path")
		splittedPath := strings.Split(path, "/")
		currentDirectory := FileSystem.Root
		for index, p := range splittedPath {
			if index == len(splittedPath)-1 {
				c.JSON(http.StatusOK, currentDirectory.Files[p].Size)
				return
			}
			currentDirectory = currentDirectory.Subdirectories[p]
		}
	}
}

func handlePostFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"status": "error", "errorMessage": err.Error()})
			return
		}
		files := form.File["files"]
		paths := form.Value["paths"]
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
}

func handleDeleteFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Query("path")
		splittedPath := strings.Split(path, "/")
		currentDirectory := FileSystem.Root
		for index, p := range splittedPath {
			if index == len(splittedPath)-1 {
				err := filesystem.DeleteFile(filepath.Join("root", path))
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "error", "errorMessage": err.Error()})
					return
				}
				delete(currentDirectory.Files, p)
				return
			}
			currentDirectory = currentDirectory.Subdirectories[p]
		}
	}
}

func handlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	}
}

func defaultHandleRecovery(c *gin.Context, err interface{}) {
	log.Logger.Error().Msgf("Recovered from panic: %v", err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "error", "errorMessage": http.StatusText(http.StatusInternalServerError)})
}
