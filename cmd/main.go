package main

import (
	"app/cmd/handlers"
	"app/internal/vehicle/loader"
	"app/internal/vehicle/repository"
	"app/internal/vehicle/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env
	godotenv.Load(".env")
	file_path := os.Getenv("FILE_PATH_VEHICLES_JSON")

	// dependencies
	ldVh := loader.NewLoaderVehicleJSON(file_path)
	dbVh, err := ldVh.Load() // map[key, value] -> list of vehicles from json file
	if err != nil {
		panic(err)
	}

	rpVh := repository.NewRepositoryVehicleInMemory(dbVh, file_path)
	svVh := service.NewServiceVehicleDefault(rpVh)
	ctVh := handlers.NewControllerVehicle(svVh)

	// server
	rt := gin.New()
	// -> middlewares
	rt.Use(gin.Recovery())
	rt.Use(gin.Logger())
	// -> handlers
	// 1 group -> path base 
	api  := rt.Group("/api/v1")
	// 2 group ->
	grVh := api.Group("/vehicles")
	// endpoints 
	grVh.GET("", ctVh.GetAll())
	grVh.POST("/batch", ctVh.SaveVehicles())

	// run
	if err := rt.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}
}