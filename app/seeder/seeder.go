package seed

import (
	"ta_microservice_pendataan/app/models"
	"ta_microservice_pendataan/db"

	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		// {
		// 	Name: "Delete",
		// 	Run: func(db *gorm.DB) error {
		// 		err := DeleteSeed()
		// 		return err
		// 	},
		// },
		{
			Name: "Create Divisi RC",
			Run: func(db *gorm.DB) error {
				err := CreatePackage(db, "Panjat Tebing")
				return err
			},
		},
		{
			Name: "Create Divisi GH",
			Run: func(db *gorm.DB) error {
				err := CreatePackage(db, "Gunung Hutan")
				return err
			},
		},
		{
			Name: "Create Divisi Diving",
			Run: func(db *gorm.DB) error {
				err := CreatePackage(db, "Selam")
				return err
			},
		},
	}
}

func CreatePackage(db *gorm.DB, name string) error {

	return db.Create(&models.Divisi{
		Name: name,
	}).Error
}

func DeleteSeed() error {
	// 	return app.Db.Delete(models.Packages{}).Error
	// err := app.Db.Exec("DELETE FROM top_up_packages").Error
	err := db.Db.Exec("DROP TABLE alats").Error
	return err
}
