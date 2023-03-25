package server

import (
	"go-management-auth-school/config"

	"github.com/labstack/echo/v4"
)


func InitApp(router *echo.Echo, conf config.Config, unitTest bool) {

	config.MasterDB = config.SetupMasterDB(conf)
	// setup slave db
	config.SlaveDB = config.SetupSlaveDB(conf)



	v1 := router.Group("/v1")
	// v2 := router.Group("/v2")

	// v1 api group
	apiUserV1 := v1.Group("/apiUser")
	apiAdminV1 := v1.Group("/apiAdmin")
	apiAuthV1 := v1.Group("/apiAuth")
	apiStaticv1 := v1.Group("/apiStatic")


	apiUserV1.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	apiAdminV1.GET("/test", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	apiAuthV1.GET("/test", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	apiStaticv1.GET("/test", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})


	// v2 api group
	// apiUserV2 := v2.Group("/apiUser")
	// apiAdminV2 := v2.Group("/apiAdmin")
	// apiAuthV2 := v2.Group("/apiAuth")
	// apiStaticv2 := v2.Group("/apiStatic")


}
