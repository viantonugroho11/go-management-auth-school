package server

import (
	"context"
	"fmt"
	"go-management-auth-school/config"
	"log"
	"net/http"

	// echo
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"os"
	"github.com/joho/godotenv"
)

type RestServer struct{}

func New() *RestServer {
	return &RestServer{}
}

func (s *RestServer) Start() {
	conf := config.New()
	ctx := context.Background()


	// railway or heroku
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	router := echo.New()

	InitApp(router, conf, false)

	router.Use(middleware.Logger(), middleware.Recover())

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},

		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
	}))

	// Setup http server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router,
	}

	fmt.Println("Listening on port", os.Getenv("PORT"))
	log.Print(ctx, "Listening on port %s", os.Getenv("PORT"))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(ctx, "%v", err)
	}

}
