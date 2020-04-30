package api

import (
	"base-server/pkg/app"
	"base-server/pkg/e"
	"base-server/service/label_service"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//-----------------------------------------------------------------------------

type labelForm struct {
	Id    int    `form:"id" valid:"Required;" json:"id"`
	Text  string `form:"text" valid:"Required" json:"text"`
	Audio string `form:"audio" json:"audio"`
}

func GetTotalLabel(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"total": 10,
	})
}

func GetLabelByIdx(c *gin.Context) {
	appG := app.Gin{C: c}
	idx := com.StrTo(c.Param("index")).MustInt()
	log.Infoln(idx)

	file, err := os.Open("/home/sxt/wav/1000hoursrc/CLEAN_1_0997.wav")
	if err != nil {
		log.Infoln("open file failed")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Infoln("read file failed")
	}
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"id":    1,
		"text":  "你好，你叫什么名字",
		"audio": com.Base64Encode(string(data)),
	})
}

func ListLabel(c *gin.Context) {
	appG := app.Gin{C: c}
	//idx := com.StrTo(c.Param("index")).MustInt()

	//file, err := os.Open("/home/sxt/wav/1000hoursrc/CLEAN_1_0997.wav")
	//if err != nil {
	//	log.Infoln("open file failed")
	//}
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	log.Infoln("read file failed")
	//}
	wavIp := "http://192.168.5.24:8000/wav/"
	ret := []labelForm{
		labelForm{
			Id:    1,
			Text:  "11111111",
			Audio: wavIp + "CLEAN_210_0714.wav",
		},
		labelForm{
			Id:    2,
			Text:  "2222222",
			Audio: wavIp + "CLEAN_210_0715.wav",
		},
	}
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"labels": ret,
	})
}

func AddLabel(c *gin.Context) {
	appG := app.Gin{C: c}

	var form labelForm
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	_, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if header == nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	rand.Seed(time.Now().UnixNano())
	wavFile := fmt.Sprintf("%d.wav", rand.Int63())
	if err := c.SaveUploadedFile(header, wavFile); err != nil {
		log.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ErrorUploadSaveFileFail, nil)
		return
	}

	labelService := label_service.Label{
		ID:    form.Id,
		Text:  form.Text,
		Audio: wavFile,
	}
	if err := labelService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAddLabelFail, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
