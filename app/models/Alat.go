package models

import (
	"errors"
	"fmt"
	"time"

	"ta_microservice_pendataan/db"

	"gorm.io/gorm"
)

type Alat struct {
	Id         int       `json:"id" gorm:"primarykey"`
	DivisiId   int       `json:"divisiId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:DivisiId"`
	Nama       string    `json:"nama"`
	Jumlah     int       `json:"jumlah"`
	Keterangan string    `json:"keterangan"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Divisi     *Divisi   `json:"divisi" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Divisi struct {
	Id   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}

func CreateAlat(db *gorm.DB, a *Alat) (*Alat, error) {
	// result := db.Create(a)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// var alat Alat
	if err := db.Preload("Divisi").Create(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func GetAllAlat(db *gorm.DB) ([]Alat, error) {
	var alat []Alat
	if err := db.Preload("Divisi").Order("CASE nama WHEN 'Panjat Tebing' THEN 1 WHEN 'Gunung Hutan' THEN 2 WHEN 'Selam' THEN 3 ELSE 4 END").Find(&alat).Error; err != nil {
		return nil, err
	}
	return alat, nil
}

func GetAlatById(query *gorm.DB, id int) (*Alat, error) {
	p := Alat{}
	query.First(&p, id)
	if p.Id == 0 {
		return nil, errors.New("")
	}
	return &p, nil
}

func Delete(query *gorm.DB, id uint) (err error) {
	d := Alat{}
	query.Delete(d, "id = ?", id)
	if d.Id == 0 {
		return
	}
	return nil
}

func UpdateById(id int) (*Alat, error) {
	var alat Alat
	err := db.Db.Where("id = ?", id).First(&alat).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("alat with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get alat with ID %d: %v", id, err)
	}
	return &alat, nil
}

func GetAllDivisi(query *gorm.DB, Package *[]Divisi) (err error) {
	err = query.Find(Package).Error
	if err != nil {
		return err
	}

	return nil
}

func GeDivisiById(query *gorm.DB, id int) (*Divisi, error) {
	p := Divisi{}
	query.First(&p, id)
	if p.Id == 0 {
		return nil, errors.New("Divisi not found")
	}
	return &p, nil
}

func CountAlatsByDivisiID(divisiID int) (int64, error) {
	var count int64
	if err := db.Db.Table("alat").Where("divisi_id = ?", divisiID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
