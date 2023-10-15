package router

import (
	"Cross-field-shop/handler/commodities"
	"Cross-field-shop/handler/consignee"
	"Cross-field-shop/handler/history"
	"Cross-field-shop/handler/purchase"
	"Cross-field-shop/pkg/constvar"
	"net/http"

	"github.com/gin-gonic/gin"

	"Cross-field-shop/handler/sd"
	"Cross-field-shop/handler/user"
	"Cross-field-shop/router/middleware"
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

	normalRequired := middleware.AuthMiddleware(constvar.AuthLevelNormal)
	//adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)
	//superAdminRequired := middleware.AuthMiddleware(constvar.AuthLevelSuperAdmin)

	userRouter := g.Group("/user")
	{
		userRouter.POST("/login", user.Login)
	}

	historyRouter := g.Group("/history")
	historyRouter.Use(normalRequired)
	{
		historyRouter.GET("", history.List)
		historyRouter.POST("", history.Post)
	}

	purchaseRouter := g.Group("/purchase")
	purchaseRouter.Use(normalRequired)
	{
		purchaseRouter.GET("", purchase.ListCart)
		purchaseRouter.POST("", purchase.Post)
		purchaseRouter.DELETE("/:id", purchase.DeleteCart)
		purchaseRouter.DELETE("/:id", purchase.UpdateCart)
	}

	consigneeRouter := g.Group("/consignee")
	consigneeRouter.Use(normalRequired)
	{
		consigneeRouter.GET("", consignee.List)
		consigneeRouter.POST("", consignee.Post)
		consigneeRouter.DELETE("/:id", consignee.DeleteConsignee)
		consigneeRouter.DELETE("/:id", consignee.UpdateConsignee)
	}

	commoditiesRouter := g.Group("/commodities")
	commoditiesRouter.Use(normalRequired)
	{
		commoditiesRouter.GET("", commodities.List)
		//consigneeRouter.POST("", consignee.Post)
		//consigneeRouter.DELETE("/:id", consignee.DeleteConsignee)
		//consigneeRouter.DELETE("/:id", consignee.UpdateConsignee)
	}

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
