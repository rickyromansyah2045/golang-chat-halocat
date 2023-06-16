package main

import (
	"server/internal/ws"
	"server/router"
)

func main() {

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(wsHandler)
	router.Start(":8989")
}
