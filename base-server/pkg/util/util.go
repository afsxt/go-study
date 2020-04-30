package util

import "base-server/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
