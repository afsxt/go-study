package api

import (
	"base-server/pkg/app"
	"base-server/pkg/e"
	"base-server/pkg/util"
	"base-server/service/auth_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//-----------------------------------------------------------------------------

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

// GetAuth get auth
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	//username := c.PostForm("username")
	//password := c.PostForm("password")

	//a := auth{Username: username, Password: password}
	var a auth
	err := c.BindJSON(&a)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	ok, _ := valid.Valid(&a)
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	//authService := auth_service.Auth{Username: username, Password: password}
	authService := auth_service.NewAuth(a.Username, a.Password)
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuthCheckTokenFail, nil)
		return
	}
	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ErrorAuth, nil)
		return
	}

	token, err := util.GenerateToken(a.Username, a.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuthToken, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
