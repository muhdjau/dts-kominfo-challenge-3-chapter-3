package main

import (
	"challenge-chapter-3-sesi-3/config"
	"challenge-chapter-3-sesi-3/routers"
)

func main() {
	config.ConnectDB()

	routers.StartServer().Run(":80")
}
