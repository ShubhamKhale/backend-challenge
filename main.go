package main

import (
	"fmt"
	"task/api"
	"task/game"
	"task/mockuser"
)

func main() {

	// starting api and game engine
	engine := game.NewGameEngine()
	server := api.APIServer{Game: engine}

	// starting server in goroutine
	go func() {
		fmt.Println("API server running at :8080")
		server.Start(":8080")
	}()

	// starting mock users
	mockuser.SimulateUsers(1000, "http://localhost:8080")

	uid, cor, incor := engine.Metrics()
	fmt.Printf("Winner's User Id: %d | Correct Answer Count: %d | Incorrect Answer Count: %d\n", uid, cor, incor)
	engine.Close()
}
