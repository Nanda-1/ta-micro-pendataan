package main

import (
	"ta_microservice_pendataan/app/controllers"
	"ta_microservice_pendataan/app/routers"
)

func main() {

	packagePendataan := controllers.NewPendataan()
	packageSeeder := controllers.NewDbSeerder()

	r := routers.SetupRouter(*packagePendataan, *packageSeeder)
	err := r.Run(":8060")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
