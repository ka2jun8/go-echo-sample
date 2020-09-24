package main

import (
	"github.com/ka2jun8/go-echo-sample/server"
)

func main() {
	router := server.Router()

	router.Logger.Info("start server [port:1323]")
	// Start server
	router.Logger.Fatal(router.Start(":1323"))
}
