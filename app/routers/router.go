package routers

import (
	"ta_microservice_pendataan/app/controllers"
	"ta_microservice_pendataan/app/middleware"

	"github.com/gin-gonic/gin"
)

type API struct {
	RepoPendataan controllers.PendataanRepo
	RepoSeeder    controllers.DbSeeder
}

func SetupRouter(RepoPendataan controllers.PendataanRepo, RepoSeeder controllers.DbSeeder) *gin.Engine {
	r := gin.New()
	api := API{
		RepoPendataan, RepoSeeder,
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	protectedRouter := r.Group("/api/pendataan")
	protectedRouter.Use(middleware.CORSMiddleware(), middleware.ApiKey(), middleware.ReqJson(), middleware.Jwt())
	protectedRouter.POST("/create", api.RepoPendataan.CreateAlat)
	protectedRouter.POST("/seeder", api.RepoSeeder.RunSeeder)
	protectedRouter.POST("/delete", api.RepoPendataan.DeleteByID)
	protectedRouter.POST("/update", api.RepoPendataan.Update)
	protectedRouter.GET("/get-all", api.RepoPendataan.GetAll)
	protectedRouter.GET("/get", api.RepoPendataan.GetAlatbyId)

	return r
}
