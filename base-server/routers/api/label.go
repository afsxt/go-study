package api

import (
	"base-server/pkg/app"
	"base-server/pkg/e"
	"base-server/service/label_service"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

//-----------------------------------------------------------------------------

type labelForm struct {
	Id    int    `form:"id" valid:"Required;" json:"id"`
	Text  string `form:"text" json:"text"`
	PY    string `form:"py" json:"py"`
	Audio string `form:"audio" json:"audio"`
	Trc   string `form:"trc" json:"trc"`
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
	wavIp := "http://52.163.245.77:8000/wav/"
	//ret := []labelForm{
	//	labelForm{
	//		Id:    1,
	//		Text:  "明天可能会下雨。",
	//		PY:    "ming2 tian1 ke3 neng2 hui4 xia4 yu3.",
	//		Audio: wavIp + "1.wav",
	//		Trc:   "[00:0.1]<190>明<190>天<190>可<190>能<190>会<190>下<190>雨<190>。<190>",
	//	},
	//	labelForm{
	//		Id:    2,
	//		Text:  "我的钱快要用完了。",
	//		PY:    "wo3 de qian2 kuai4 yao4 yong4 wan2 le.",
	//		Audio: wavIp + "2.wav",
	//		Trc:   "",
	//	},
	//}
	var ret []labelForm
	file, err := os.Open("audio/50PINYIN.txt")
	if err != nil {
		log.Error("open wav txt file failed", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	reader := bufio.NewReader(file)
	var line string
	i := 0
	var tmpLabel labelForm
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		i++

		tmpLabel.Id = i
		tmpLabel.PY = line
		trcFile, err := os.Open(
			fmt.Sprintf("%s%d_TRC_pinyin.txt", "audio/TRC/", tmpLabel.Id))
		if err != nil {
			log.Infoln("open file failed")
		}
		data, err := ioutil.ReadAll(trcFile)
		if err != nil {
			log.Infoln("read file failed")
		}
		tmpLabel.Trc = strings.Replace(string(data), "\n\n", "", 1)
		tmpLabel.Audio = fmt.Sprintf("%s%d.wav", wavIp, tmpLabel.Id)
		ret = append(ret, tmpLabel)
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
