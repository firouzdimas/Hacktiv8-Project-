package main

import (
	"project-2/config"
	"project-2/routers"
)

func main() {
	var PORT = ":8080"

	// connect to DB
	config.ConnectDB()
	// env := new(controllers.Env)
	// var err error
	// env.DB, err = controllers.ConnectDB()
	// if err != nil {
	// 	log.Fatalf("failed to start the server: %v", err)
	// }

	routers.StartServer().Run(PORT)
}
