package project

import (
	"tugas5/project/controller"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	router     *gin.Engine
	auth       *controller.AuthHandler
	user       *controller.UserHandler
	category   *controller.CategoryHandler
	product    *controller.ProductHandler
	cart       *controller.CartHandler
	middleware *Middleware
}

func NewRouterGin(
	router *gin.Engine,
	auth *controller.AuthHandler,
	user *controller.UserHandler,
	product *controller.ProductHandler,
	category *controller.CategoryHandler,
	cart *controller.CartHandler,
	middleware *Middleware,
) *GinRouter {
	return &GinRouter{
		router:     router,
		auth:       auth,
		user:       user,
		product:    product,
		category:   category,
		cart:       cart,
		middleware: middleware,
	}
}

func (r *GinRouter) Start(port string) {
	r.router.Use(r.middleware.Trace)

	auth := r.router.Group("/auth")
	auth.POST("/register", r.auth.Register)
	auth.POST("/login", r.auth.Login)

	user := r.router.Group("/users")
	user.POST("/", r.middleware.Auth, r.middleware.Check(r.user.Create, []string{"admin"}))
	user.GET("/", r.middleware.Auth, r.middleware.Check(r.user.FindAll, []string{"admin"}))
	user.GET("/profile", r.middleware.Auth, r.middleware.Check(r.user.Profile, []string{"admin", "user"}))
	user.PUT("/profile", r.middleware.Auth, r.middleware.Check(r.user.Update, []string{"admin", "user"}))
	user.GET("/email/:email", r.middleware.Auth, r.middleware.Check(r.user.FindByEmail, []string{"admin"}))

	product := r.router.Group("/products")
	product.POST("/", r.middleware.Auth, r.middleware.Check(r.product.CreateProduct, []string{"admin"}))
	product.GET("/", r.middleware.Auth, r.middleware.Check(r.product.GetProducts, []string{"admin", "user"}))
	product.GET("/:productId", r.middleware.Auth, r.middleware.Check(r.product.GetProductById, []string{"admin", "user"}))
	product.PUT("/:productId", r.middleware.Auth, r.middleware.Check(r.product.UpdateProduct, []string{"admin"}))
	product.PATCH("/:productId", r.middleware.Auth, r.middleware.Check(r.product.DeleteProduct, []string{"admin"}))

	category := r.router.Group("/categories")
	category.POST("/", r.middleware.Auth, r.middleware.Check(r.category.CreateCategory, []string{"admin"}))
	category.GET("/", r.middleware.Auth, r.middleware.Check(r.category.GetCategories, []string{"admin", "user"}))
	category.GET("/:categoryId", r.middleware.Auth, r.middleware.Check(r.category.GetCategoryById, []string{"admin", "user"}))
	category.PUT("/:categoryId", r.middleware.Auth, r.middleware.Check(r.category.UpdateCategory, []string{"admin"}))
	category.PATCH("/:categoryId", r.middleware.Auth, r.middleware.Check(r.category.DeleteCategory, []string{"admin"}))

	cart := r.router.Group("/carts")
	cart.GET("/", r.middleware.Auth, r.middleware.Check(r.cart.GetCarts, []string{"admin"}))
	cart.POST("/", r.middleware.Auth, r.middleware.Check(r.cart.CreateCart, []string{"admin", "user"}))
	cart.GET("/session", r.middleware.Auth, r.middleware.Check(r.cart.GetUserCarts, []string{"admin", "user"}))
	cart.GET("/:cartId", r.middleware.Auth, r.middleware.Check(r.cart.GetCartById, []string{"admin"}))
	cart.PUT("/:cartId", r.middleware.Auth, r.middleware.Check(r.cart.UpdateCart, []string{"admin", "user"}))
	cart.PATCH("/:cartId", r.middleware.Auth, r.middleware.Check(r.cart.DeleteCart, []string{"admin", "user"}))

	r.router.Run(port)
}
