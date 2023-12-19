package router

import (
	"bytes"
	"cloud-storage-backend/filesystem"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewGinRouter() *gin.Engine {

	filesystem.InitializeFileSystem("root")

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

	engine.Use(
		gin.LoggerWithConfig(logConf),
		gin.CustomRecoveryWithWriter(log.Logger, defaultHandleRecovery),
	)

	router := engine.Group("/api")

	router.GET("/", handleGetFile())

	router.POST("/", handlePostFile())

	router.DELETE("/", handleDeleteFile())

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

	fileGroup.GET("/", handleGetFile())
	fileGroup.POST("/", handlePostFile())
	fileGroup.DELETE("/", handleDeleteFile())
}

func addDirectoryGroup(router *gin.RouterGroup) {
	directoryGroup := router.Group("/directory")

	directoryGroup.GET("/", handleGetDirectory())
	directoryGroup.DELETE("/", handleDeleteDirectory())
}
