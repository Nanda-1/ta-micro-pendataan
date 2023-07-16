package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"ta_microservice_pendataan/app/models"
	"ta_microservice_pendataan/db"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PendataanRepo struct {
	Db *gorm.DB
}

func NewPendataan() *PendataanRepo {
	db := db.InitDb()
	db.AutoMigrate(&models.Alat{})
	return &PendataanRepo{Db: db}
}

func (repo *PendataanRepo) CreateAlat(c *gin.Context) {
	res := models.JsonResponse{Success: true}
	req := models.Alat{}
	err := c.BindJSON(&req)
	if err != nil {
		errorMsg := err.Error()
		res.Success = false
		res.Error = &errorMsg
		c.JSON(400, res)
		c.Abort()
		return
	}
	// Custom validation to check if req is empty
	if req.Nama == "" || req.Jumlah == 0 || req.Keterangan == "" {
		errorMsg := "Field Tidak Boleh Kosong"
		res.Success = false
		res.Error = &errorMsg
		c.JSON(400, res)
		c.Abort()
		return
	}

	// ambil divisi dari database
	divisi := models.Divisi{}
	err = repo.Db.First(&divisi, req.DivisiId).Error
	if err != nil {
		errorMsg := "Divisi not found"
		res.Success = false
		res.Error = &errorMsg
		c.JSON(400, res)
		c.Abort()
		return
	}

	req.Divisi = &divisi
	addAlat, err := models.CreateAlat(repo.Db, &req)
	if err != nil {
		// Check if the error is related to the 'created_at' column
		if strings.Contains(err.Error(), "Incorrect datetime value") {
			// Set a valid value for the 'created_at' column
			req.Created_at = time.Now()
			req.Updated_at = time.Now()
			// Retry creating the Alat
			addAlat, err = models.CreateAlat(repo.Db, &req)
		}
		if err != nil {
			errorMsg := err.Error()
			res.Success = false
			res.Error = &errorMsg
			c.JSON(400, res)
			c.Abort()
			return
		}
	}

	// update relasi Alat -> Divisi
	repo.Db.Model(&divisi).Association("Alat").Append(addAlat)

	res.Data = addAlat
	c.JSON(http.StatusOK, res)
}

func (repo *PendataanRepo) Update(c *gin.Context) {
	res := models.JsonResponse{Success: true}
	req := models.Alat{}
	err := c.BindJSON(&req)
	if err != nil {
		errorMsg := err.Error()
		res.Success = false
		res.Error = &errorMsg
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}
	getId := c.Query("id")
	id, err := strconv.Atoi(getId)
	if err != nil {
		errMsg := "ID tidak Boleh kosong"
		res.Success = false
		res.Error = &errMsg
		c.JSON(404, res)
		c.Abort()
		return
	}

	if req.Nama == " " || req.Jumlah <= 0 || req.Divisi.Id == 0 {
		errorMsg := "invalid Input"
		res.Success = false
		res.Error = &errorMsg
		c.JSON(400, res)
	}

	existingAlat, err := models.UpdateById(id)
	if err != nil {
		ErroMsg := "Error updating"
		res.Success = false
		res.Error = &ErroMsg
		c.JSON(http.StatusBadRequest, res)
		return
	}

	existingAlat.Nama = req.Nama
	existingAlat.Jumlah = req.Jumlah
	existingAlat.Keterangan = req.Keterangan
	existingAlat.Divisi.Id = req.Divisi.Id

	if err := repo.Db.Save(&existingAlat).Error; err != nil {
		ErrorMsg := err.Error()
		res.Success = false
		res.Error = &ErrorMsg
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res.Data = existingAlat
	c.JSON(200, res)

}

func (repo *PendataanRepo) GetAlatbyId(c *gin.Context) {
	res := models.JsonResponse{Success: true}

	getId := c.Query("id")
	id, err := strconv.Atoi(getId)
	if err != nil {
		errMsg := "ID tidak Boleh kosong"
		res.Success = false
		res.Error = &errMsg
		c.JSON(404, res)
		c.Abort()
		return
	}

	idUser, _ := models.GetAlatById(repo.Db, id)

	res.Data = idUser

	c.JSON(200, res)

}

func (repo *PendataanRepo) DeleteByID(c *gin.Context) {
	res := models.JsonResponse{Success: true}
	// req := models.Alat{}

	PostId := c.Query("id")
	id, err := strconv.Atoi(PostId)
	if err != nil {
		errMsg := "ID tidak Boleh kosong"
		res.Success = false
		res.Error = &errMsg
		c.JSON(404, res)
		c.Abort()
		return
	}

	result := models.Delete(repo.Db, uint(id))
	// db.Db.Delete(req, "id = ? ", PostId)
	if result != nil {
		errorMsg := "Id Not Found"
		res.Success = false
		res.Error = &errorMsg
		c.JSON(404, res)
		c.Abort()
		return
	}

	res.Data = result

	c.JSON(200, res)
}

func (repo *PendataanRepo) GetAll(c *gin.Context) {
	res := models.JsonResponse{Success: true}

	alat, err := models.GetAllAlat(repo.Db)
	if err != nil {
		errorMsg := "Failed to fetch Alat data."
		res.Success = false
		res.Error = &errorMsg
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Data = alat
	c.JSON(http.StatusOK, res)
}

func (repo *PendataanRepo) GetAllDivisi(c *gin.Context) {
	res := models.JsonResponse{Success: true}

	var divisi []models.Divisi

	err := models.GetAllDivisi(repo.Db, &divisi)
	if err != nil {
		c.AbortWithError(500, err)
	}

	res.Data = &divisi

	c.JSON(200, res)
}

func (repo *PendataanRepo) CountAlatsByDivisiID(c *gin.Context) {
	res := models.JsonResponse{Success: true}

	idStr := c.Query("divisi_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorMsg := "Invalid divisi_id"
		res.Success = false
		res.Error = &errorMsg
		c.JSON(http.StatusBadRequest, res)
		return
	}

	count, err := models.CountAlatsByDivisiID(id)
	if err != nil {
		errorMsg := err.Error()
		res.Success = false
		res.Error = &errorMsg
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Data = count
	c.JSON(http.StatusOK, res)
}
