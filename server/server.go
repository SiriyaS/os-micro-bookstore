package server

import (
	"os-micro-bookstore/config"
)

func Init() {
	conf := config.GetConfig()
	router := NewRouter()
	// fmt.Printf("\n[Swagger UI] API Docs at http://localhost:%s/swagger/index.html\n", conf.GetString("server.port"))
	router.Run(":" + conf.GetString("server.port"))
}
