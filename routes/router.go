package routes

import (
	"github.com/fazaalexander/go-gorm/controller"

	echojwt "github.com/labstack/echo-jwt/v4"

	// "project/vendor/github.com/golang-jwt/jwt"
	// "project/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())

	e.POST("/login", controller.Login)

	// config := middleware.JWTConfig{
	// 	Claims:     &controller.JWTCustomClaims{},
	// 	SigningKey: []byte("secret"),
	// }
	jwtMiddleware := echojwt.JWT([]byte("secret"))
	// jwtMiddleware := middleware.JWTWithConfig(config)

	userController := controller.Controller{}
	userGroup := e.Group("/users")
	userGroup.POST("/", userController.CreateUserController)
	userGroup.Use(jwtMiddleware)
	userGroup.GET("/", userController.GetUsersController)
	userGroup.GET("/:id", userController.GetUserController)
	userGroup.DELETE("/:id", userController.DeleteUserController)
	userGroup.PUT("/:id", userController.UpdateUserController)

	bookGroup := e.Group("/books")
	bookGroup.Use(jwtMiddleware)
	bookGroup.GET("/", controller.GetBooksController)
	bookGroup.GET("/:id", controller.GetBookController)
	bookGroup.POST("/", controller.CreateNewBook)
	bookGroup.DELETE("/:id", controller.DeleteBookController)
	bookGroup.PUT("/:id", controller.UpdateBookController)

	blogGroup := e.Group("/blogs")
	blogGroup.Use(jwtMiddleware)
	blogGroup.GET("/", controller.GetBlogsController)
	blogGroup.GET("/:id", controller.GetBlogController)
	blogGroup.POST("/", controller.CreateBlogController)
	blogGroup.DELETE("/:id", controller.DeleteBlogController)
	blogGroup.PUT("/:id", controller.UpdateBlogController)
	return e
}
