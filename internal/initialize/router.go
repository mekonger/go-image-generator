package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/mekonger/go-image-generator/global"
	"github.com/mekonger/go-image-generator/internal/routers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	if global.ServerMode == "DEV" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	userRouter := routers.UserRouter{}
	mainGroup := r.Group("api/v1")
	{
		userRouter.InitUserRoutes(mainGroup)
	}

	return r
}
