package app

import (
	"base-server/pkg/e"
	"github.com/astaxie/beego/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

//-----------------------------------------------------------------------------

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.InvalidParams
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.InvalidParams
	}

	return http.StatusOK, e.SUCCESS
}
