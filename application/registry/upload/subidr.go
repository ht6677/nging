package upload

import "github.com/webx-top/echo"

var Subdir = echo.NewKVData()

func init() {
	Subdir.Add(`default`, `默认`)
	Subdir.Add(`avatar`, `头像`)
}
