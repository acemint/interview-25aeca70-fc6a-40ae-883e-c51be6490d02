package gobookcabin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VoucherController struct {
	service *GormVoucherService
}

func NewVoucherController(service *GormVoucherService) *VoucherController {
	return &VoucherController{
		service: service,
	}
}

func (v *VoucherController) Check(c *gin.Context) {
	var requestBody CheckVoucherRequest
	err := shouldBindJsonAndValidate(c, &requestBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	voucher, err := v.service.Check(c, &requestBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, NewCheckVoucherResponse(voucher))

}

func (v *VoucherController) Generate(c *gin.Context) {
	var requestBody GenerateVoucherRequest
	err := shouldBindJsonAndValidate(c, &requestBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	voucher, err := v.service.Generate(c, &requestBody)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, NewGenerateVoucherResponse(voucher))
}
