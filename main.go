package main

import (
	// "go-management-auth-school/config"
	"go-management-auth-school/server"
	// "log"
)


func main() {
	// conf := config.New()

	// if conf.Env == "local" || conf.Env == "" {
	// 	// config githook for development only
	// 	if err := exec.Command("git", "config", "core.hooksPath", ".githooks").Run(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// swagger.SwaggerInfo(conf)
	srv := server.New()
	srv.Start()
}