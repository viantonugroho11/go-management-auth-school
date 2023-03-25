package server

import (
	"context"
	"fmt"
	"go-management-auth-school/config"
	"net/http"
	"log"

	// echo
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



type RestServer struct{}

func New() *RestServer {
	return &RestServer{}
}

func (s *RestServer) Start() {
	conf := config.New()
	ctx := context.Background()

	router := echo.New()

	InitApp(router, conf, false)

	router.Use(middleware.Logger(), middleware.Recover())
	
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},

	}))

	// Setup http server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: router,
	}

	log.Print(ctx, "Listening on port %d", conf.Port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(ctx, "%v", err)
	}
	
}