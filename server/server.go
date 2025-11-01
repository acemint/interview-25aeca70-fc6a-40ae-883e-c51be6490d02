package server

import (
	"github.com/gin-gonic/gin"
	"gobookcabin/gobookcabin"
)

const (
	RouteApi      = "/api"
	RouteCheck    = "/check"
	RouteGenerate = "/generate"
)

func SetupGinEngine(voucherController *gobookcabin.VoucherController) *gin.Engine {
	router := gin.Default()

	router.Use(ErrorHandlerMiddleware())
	vRouter := router.Group(RouteApi)
	{
		vRouter.POST(RouteCheck, voucherController.Check)
		vRouter.POST(RouteGenerate, voucherController.Generate)
	}
	return router

}
