package main

import (
	"net/http"
	"os"

	"./controllers"
	"./models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/user/new", controllers.CreateAccount)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	r := e.Group("/api/user/login")

	config := middleware.JWTConfig{
		Claims:     &models.Token{},
		SigningKey: []byte(os.Getenv("token_password")),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.POST("", controllers.Authenticate)
	// e.Use(middleware.JWT([]byte("secret")))
	// router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	// router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	e.Logger.Fatal(e.Start(":" + port))
}
