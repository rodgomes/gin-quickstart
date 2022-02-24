package server

import "github.com/rodgomes/go-quickstart/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(":" + config.GetString("server.port"))
}
