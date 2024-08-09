package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	middleware "github.com/tans1/go-web-server/api/v1/middlewares"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	controllers := New(db)
	
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})
	
	// auth routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", controllers.authController.Register)
		authRoutes.POST("/login", controllers.authController.Login)

		// for token validation
		protected := authRoutes.Group("/")
		protected.Use(middleware.AuthMiddleware(db))
		{
			protected.GET("/validate",  func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "User token is valid",
				})
			})
		}
		
	}

	// protected
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(db))
	{
		// blog routes
		blog := protected.Group("/blog")
		{
			blog.POST("/create", controllers.blogController.Create)
			blog.GET("/:id", controllers.blogController.GetById)

			blog.GET("/user/:id", func(ctx *gin.Context) {
				ctx.JSON(500, gin.H{
					"message": "Get by user Id not implemented",
				})
			})
			blog.PATCH("/:id", func(ctx *gin.Context) {
				ctx.JSON(500, gin.H{
					"message": "Not Implemented",
				})
			})
			blog.DELETE("/:id", func(ctx *gin.Context) {
				ctx.JSON(500, gin.H{
					"message": "Delete not Implemented",
				})
			})
		}

		// comment routes
		comment := protected.Group("/comment")
		{
			// get comment by id along with all the replies
			comment.GET("/:id", controllers.commentController.GetById)
			comment.POST("/blog/:id", controllers.commentController.Create)

			// get comment by blog id along with all the replies
			comment.GET("/blog/:id", func(ctx *gin.Context) {
				ctx.JSON(500, gin.H{
					"message": "Get by blog id is not implemented",
				})
			})
			comment.PATCH("/:id", func(ctx *gin.Context) {
				ctx.JSON(500, gin.H{
					"message": "Not Implemented",
				})
			})
			comment.DELETE("/:id", func(ctx *gin.Context) {
				ctx.JSON(500, gin.H{
					"message": "Delete not Implemented",
				})
			})
		}

		// commentReply routes
		commentReply := protected.Group("/reply")
		{
			commentReply.POST("/comment/:id", controllers.replyController.Create)
			// get reply by id
			commentReply.GET("/:id", controllers.replyController.GetById)
			// get replies by comment Id along with the blog
			commentReply.GET("/comment/:id", controllers.replyController.GetRepliesByCommentId)

			commentReply.PATCH("/:id", func(ctx *gin.Context) {
				ctx.JSON(500, gin.H{
					"message": "Not implemented",
				})
			})

			commentReply.DELETE("/:id", func(ctx *gin.Context) {
				ctx.JSON(500, gin.H{
					"message": "Not implemented",
				})
			})
		}
	}
}
