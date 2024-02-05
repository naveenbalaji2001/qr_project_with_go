package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naveenbalaji2001/qr_project_with_go/config"
	"github.com/naveenbalaji2001/qr_project_with_go/services"
)

func main() {
	r := gin.Default()
	redisClient := config.SetupRedis()
	neo4jDriver, err := config.GetNeo4jDriver()
	if err != nil {
		log.Fatal(err)
	}

	defer redisClient.Close()
	defer neo4jDriver.Close()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.GET("/:id", func(c *gin.Context) {
		services.GetProductByID(c, redisClient, neo4jDriver)
	})

	r.POST("/product/create", func(c *gin.Context) {
		services.CreateProduct(c, redisClient, neo4jDriver)
	})

	fmt.Println("Server started on port 8080")
	log.Fatal(r.Run(":8080"))
}
