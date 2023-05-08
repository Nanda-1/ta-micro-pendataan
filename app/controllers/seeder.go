package controllers

import (
	"ta_microservice_pendataan/app/models"
	seed "ta_microservice_pendataan/app/seeder"
	"ta_microservice_pendataan/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DbSeeder struct {
	Db *gorm.DB
}

func NewDbSeerder() *DbSeeder {
	db := db.InitDb()
	db.AutoMigrate(&models.Divisi{})
	return &DbSeeder{Db: db}
}

func (r *DbSeeder) RunSeeder(c *gin.Context) {
	res := models.JsonResponse{Success: true}

	// seeder.DeleteSeed()

	for _, seed := range seed.All() {
		if err := seed.Run(r.Db); err != nil {
			res.Success = false
			errMsg := err.Error()
			res.Error = &errMsg
			c.JSON(400, res)
		}
	}

	c.JSON(200, res)
}
