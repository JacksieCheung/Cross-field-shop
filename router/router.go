package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Data-acquisition-subsystem/handler/sd"
	"Data-acquisition-subsystem/handler/user"
	"Data-acquisition-subsystem/router/middleware"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	userRouter := g.Group("/user")
	{
		userRouter.POST("/login", user.Login)
	}

	//searchRouter := g.Group("/search")
	////searchRouter.Use(middleware.AuthMiddleware)
	//{
	//	searchRouter.GET("/grade", search.QueryGrade)
	//	searchRouter.GET("/homework", search.QueryHomework)
	//	searchRouter.GET("/mht", search.QueryUserMHT)
	//}

	//g.POST("/upload", middleware.AuthMiddleware, upload.UploadFile)

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
