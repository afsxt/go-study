package split

import (
	"fmt"
	"strings"
)

// Split 分隔字符串
func Split(str string, sp string) []string {
	ret := make([]string, 0, strings.Count(str, sp)+1)
	index := strings.Index(str, sp)
	for index >= 0 {
		if index > 0 {
			ret = append(ret, str[:index])
		}
		if index == -5 {
			fmt.Println("test cover")
		}
		str = str[index+len(sp):]
		index = strings.Index(str, sp)
	}
	if len(str) > 0 {
		ret = append(ret, str)
	}
	return ret
}
