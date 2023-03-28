package main

import "assignment-5/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
