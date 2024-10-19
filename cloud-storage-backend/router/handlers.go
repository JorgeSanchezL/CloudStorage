package router

import (
	"cloud-storage-backend/filesystem"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func handleGetFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		path := c.Query("path")
		c.File("./fs/" + path)
	}
}

func handlePostFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
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
		c.Header("Access-Control-Allow-Origin", "*")
		path := c.Query("path")
		err := filesystem.DeleteEntry(path)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "error", "errorMessage": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	}
}

func handleGetDirectory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		path := c.Query("path")
		directory, err := filesystem.ReadDirectory("./fs/" + path)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "error", "errorMessage": err.Error()})
			return
		}
		c.JSON(http.StatusOK, directory)
	}
}

func handlePostDirectory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusNotImplemented, map[string]string{"status": "error", "errorMessage": http.StatusText(http.StatusNotImplemented)})
	}
}

func handleDeleteDirectory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		path := c.Query("path")
		err := filesystem.DeleteEntry(path)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "error", "errorMessage": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	}
}

func handlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	}
}

func defaultHandleRecovery(c *gin.Context, err interface{}) {
	log.Logger.Error().Msgf("Recovered from panic: %v", err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "error", "errorMessage": http.StatusText(http.StatusInternalServerError)})
}
