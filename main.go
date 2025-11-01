package main

import (
	"gobookcabin/gobookcabin"
	"gobookcabin/infra"
	"gobookcabin/server"
)

func main() {
	gobookcabin.InitializeGlobals()

	db, err := infra.InitializeGorm()
	if err != nil {
		panic(err)
	}

	voucherService := gobookcabin.NewGormVoucherService(db)
	voucherController := gobookcabin.NewVoucherController(voucherService)

	router := server.SetupGinEngine(voucherController)
	err = router.Run(":" + gobookcabin.AppConfigurationInstance.ServerPort)
	if err != nil {
		panic(err)
	}
}
