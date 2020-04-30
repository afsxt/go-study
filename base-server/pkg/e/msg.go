package e

//-----------------------------------------------------------------------------

var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	ErrorAddLabelFail:          "新增标注失败",
	ErrorUploadSaveFileFail:    "上传文件失败",
	InvalidParams:              "请求参数错误",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "用户名或密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
