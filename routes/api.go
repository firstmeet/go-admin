package routes

import (
	"github.com/gin-gonic/gin"
	"weserver/app/Http/Controllers/Admin"
	"weserver/app/Http/Controllers/Article"
	"weserver/app/Http/Controllers/Auth"
	"weserver/app/Http/Controllers/Category"
	"weserver/app/Http/Controllers/Goods"
	"weserver/app/Http/Controllers/GoodsType"
	"weserver/app/Http/Controllers/Spec"
	"weserver/app/Http/Controllers/Upload"
	"weserver/app/Http/Controllers/User"
	Element "weserver/app/Http/Controllers/Element"
)

func Route() {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	r.Static("/resource", "resource")

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(corsMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/upload_image", Upload.UploadImage)
	auth := r.Group("/user")
	auth.Use(Auth.WxAuth())
	{
		auth.POST("/", User.Index)
		auth.GET("/info", User.Info)
		auth.PUT("/:id", User.Update)
		auth.POST("/show", User.Show)
	}
	admin := r.Group("/admin")
	admin.POST("/login", Auth.AdminLogin)
	api:=r.Group("/api")
	api.POST("/login",Auth.Login)
	api.Use(Auth.WxAuth())
	{
         api.GET("/element",Element.Index)
         api.GET("/category",Category.Index)
	}
	admin.Use(Auth.WxAuth())
	{
		admin.GET("/info", Admin.Info)
		admin.PUT("/:id", Admin.Update)
		article := r.Group("/article")
		{
			article.GET("", Article.Index)
			article.PUT("/:id", Article.Update)
			article.POST("", Article.Create)
		}
		element:=r.Group("/element")
		{
			element.GET("",Element.Index)
			element.POST("",Element.Create)
			element.PUT("/:id",Element.Update)
		}
		category:=r.Group("/category")
		{
			category.GET("",Category.Index)
			category.POST("",Category.Create)
			category.PUT("/:id",Category.Update)
			category.DELETE("/:id",Category.Delete)
		}
		goods:=r.Group("/goods")
		{
			goods.GET("",Goods.Index)
			goods.POST("",Goods.Create)
			goods.PUT("/:id",Goods.Update)
			goods.DELETE("/:id",Goods.Delete)
		}
		goods_type:=r.Group("/goods_type")
		{
			goods_type.GET("",GoodsType.Index)
			goods_type.POST("",GoodsType.Create)
			goods_type.PUT("/:id",GoodsType.Update)
			goods_type.DELETE("/:id",GoodsType.Delete)
		}
		spec:=r.Group("/spec")
		{
			spec.GET("",Spec.Index)
			spec.POST("",Spec.Create)
			spec.PUT("/:id",Spec.Update)
			spec.DELETE("/:id",Spec.Delete)
			spec.GET("/find",Spec.FindByType)
		}
	}

	r.POST("/register", Auth.Register)
	//r.GET("/user",User.Index)
	r.Run(":8090") //
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}
