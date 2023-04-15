package main

import (
	"log"

	_ "github.com/firouzdimas/Hacktiv8-Project-/docs"

	"github.com/firouzdimas/Hacktiv8-Project-/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const PORT = ":8080"

// @title					MyGram API
// @version					1.0
// @description				This is a MyGram API.
// @host 					localhost:8080
// @BasePath 				/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	route := gin.Default()

	database.StartDB()
	db := database.GetDB()

	route.SetupUserRoute(route, db)
	route.SetupPhotoRoute(route, db)
	route.SetupSocialRoute(route, db)
	route.SetupCommentRoute(route, db)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	route.Run(PORT)
}
