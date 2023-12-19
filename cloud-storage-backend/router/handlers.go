package router

import (
	"cloud-storage-backend/filesystem"
	"net/http"
	"path/filepath"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func handleGetFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Query("path")
		filepath := filepath.Join("root", path)
		c.File(filepath)
	}
}

func handlePostFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"status": "error", "errorMessage": err.Error()})
			return
		}
		filesystem.SaveFiles(c, form)
	}
}

func handleDeleteFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Query("path")
		err := filesystem.DeleteFileOrEmptyDirectory(path)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "error", "errorMessage": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	}
}

func handleGetDirectory() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Query("path")
		directory := filesystem.GetDirectoryInfo(path)
		c.JSON(http.StatusOK, directory)
	}
}

func handleDeleteDirectory() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Query("path")
		err := filesystem.DeleteDirectory(path)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "error", "errorMessage": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]string{"status": "ok"})
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
