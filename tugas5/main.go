package main

import (
	"tugas5/connectiondb"
	project "tugas5/project"
	"tugas5/project/controller"
	"tugas5/project/database/gorm"
	"tugas5/project/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := connectiondb.ConnectGormDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(gin.Logger())

	userface := gorm.NewUserface(db)

	authService := service.NewAuthServices(userface)
	authHandler := controller.NewAuthHandler(authService)

	userService := service.NewUserServices(userface)
	userHandler := controller.NewUserHandler(userService)
	middleware := project.NewMiddleware(userService)

	categoryface := gorm.NewCategoryface(db)
	categoryService := service.NewCategoryServices(categoryface)
	categoryHandler := controller.NewCategoryHandler(categoryService)

	productface := gorm.NewProductface(db)
	productService := service.NewProductServices(productface, categoryface)
	productHandler := controller.NewProductHandler(productService)

	cartface := gorm.NewCartdb(db)
	cartService := service.NewCartServices(cartface, productface, userface)
	cartHandler := controller.NewCartHandler(cartService)

	app := project.NewRouterGin(
		router,
		authHandler,
		userHandler,
		productHandler,
		categoryHandler,
		cartHandler,
		middleware,
	)

	app.Start(":8080")
}
