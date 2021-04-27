package main

import (
	"os-micro-bookstore/config"
	"os-micro-bookstore/database"
	"os-micro-bookstore/server"
)

func main() {
	config.Init()
	database.Init()
	server.Init()
}
