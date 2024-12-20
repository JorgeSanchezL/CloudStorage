package router

import (
	"bytes"
	"cloud-storage-backend/filesystem"
	"errors"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewGinRouter() *gin.Engine {

	err := filesystem.InitializeFileSystem("fs")
	if errors.Is(err, filesystem.ErrCannotCreateRootDirectory) {
		log.Logger.Fatal().Err(err).Msg("Cannot create root directory")
	} else if err != nil && os.IsNotExist(err) {
		log.Logger.Warn().Msg("Root directory does not exist. Creating a new one...")
	} else if err != nil {
		log.Logger.Fatal().Err(err).Msg("Unexpected error")
	}

	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	logConf := gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			buf := new(bytes.Buffer)
			ginLog := zerolog.New(buf).With().Timestamp().Logger()
			latency := fmt.Sprint(param.Latency)
			ginLog.Trace().
				Str("method", param.Method).
				Str("path", param.Path).
				Str("request_proto", param.Request.Proto).
				Int("status_code", param.StatusCode).
				Str("latency", latency).
				Str("user_agent", param.Request.UserAgent()).
				Msg("New Request")
			return buf.String()
		},
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost"} // Cambia esto al origen de tu frontend
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}

	engine.Use(cors.New(config))

	engine.Use(
		gin.LoggerWithConfig(logConf),
		gin.CustomRecoveryWithWriter(log.Logger, defaultHandleRecovery),
	)

	router := engine.Group("/api")

	addXGroup(router)
	addFileGroup(router)
	addDirectoryGroup(router)

	return engine
}

func addXGroup(router *gin.RouterGroup) {
	xGroup := router.Group("/health")

	xGroup.GET("/ping", handlePing())
}

func addFileGroup(router *gin.RouterGroup) {
	fileGroup := router.Group("/file")

	fileGroup.GET("/info", handleGetFileInformation())
	fileGroup.GET("", handleGetFile())
	fileGroup.POST("", handlePostFile())
	fileGroup.DELETE("", handleDeleteFile())
}

func addDirectoryGroup(router *gin.RouterGroup) {
	directoryGroup := router.Group("/directory")

	directoryGroup.GET("", handleGetDirectory())
	directoryGroup.POST("", handlePostDirectory())
	directoryGroup.DELETE("", handleDeleteDirectory())
}
