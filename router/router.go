package router

import (
	"Cross-field-shop/handler/comments"
	"Cross-field-shop/handler/commodities"
	"Cross-field-shop/handler/consignee"
	"Cross-field-shop/handler/history"
	"Cross-field-shop/handler/purchase"
	"Cross-field-shop/handler/tags"
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
	adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)
	//superAdminRequired := middleware.AuthMiddleware(constvar.AuthLevelSuperAdmin)

	userRouter := g.Group("/user")
	{
		userRouter.POST("/login", user.Login)
		userRouter.POST("/register/validate_code", user.ValidateCode)
		userRouter.POST("/register", user.Register)
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
		purchaseRouter.GET("/cart", purchase.ListCart)
		purchaseRouter.POST("", purchase.Post)
		purchaseRouter.DELETE("/cart/:id", purchase.DeleteCart)
		purchaseRouter.DELETE("/cart/:id", purchase.UpdateCart)
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
	{
		commoditiesRouter.GET("", normalRequired, commodities.List)
		commoditiesRouter.POST("", adminRequired, commodities.Post)
		commoditiesRouter.PUT("/:id", adminRequired, commodities.UpdateCommodity)
		commoditiesRouter.DELETE("/:id", adminRequired, commodities.DeleteCommodity)
	}

	commentsRouter := g.Group("/comments")
	commentsRouter.Use(normalRequired)
	{
		commentsRouter.GET("", comments.List)
		consigneeRouter.POST("", comments.Post)
		consigneeRouter.DELETE("/:id", comments.DeleteComment)
		consigneeRouter.PUT("/:id", comments.UpdateComment)
	}

	tagsRouter := g.Group("/tags")
	tagsRouter.Use(normalRequired)
	{
		tagsRouter.GET("/:type", tags.ListTags)
		tagsRouter.POST("", tags.Post)
		tagsRouter.DELETE("/:id", tags.DeleteTags)
		tagsRouter.PUT("/:id", tags.UpdateTags)
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
