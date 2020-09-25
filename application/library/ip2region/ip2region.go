package ip2region

import (
	"fmt"
	"strings"
	"sync"

	"github.com/admpub/ip2region/binding/golang/ip2region"
	"github.com/webx-top/echo"
)

var (
	region   *ip2region.Ip2Region
	dictFile string
	once     sync.Once
)

func init() {
	dictFile = echo.Wd() + echo.FilePathSeparator + `data` + echo.FilePathSeparator + `ip2region` + echo.FilePathSeparator + `ip2region.db`

}

func SetDictFile(f string) {
	dictFile = f
}

func Initialize() (err error) {
	if region == nil {
		region, err = ip2region.New(dictFile)
	}
	return
}

func IPInfo(ip string) (info ip2region.IpInfo, err error) {
	if len(ip) == 0 {
		return
	}
	once.Do(func() {
		err = Initialize()
	})
	if err != nil {
		return
	}
	info, err = region.MemorySearch(ip)
	return
}

func Stringify(info ip2region.IpInfo) string {
	var (
		formats []string
		args    []interface{}
	)
	if info.Country != `` && info.Country != `0` {
		formats = append(formats, `"国家":%q`)
		args = append(args, info.Country)
	}
	if info.Region != `` && info.Region != `0` {
		formats = append(formats, `"地区":%q`)
		args = append(args, info.Region)
	}
	if info.Province != `` && info.Province != `0` {
		formats = append(formats, `"省份":%q`)
		args = append(args, info.Province)
	}
	if info.City != `` && info.City != `0` {
		formats = append(formats, `"城市":%q`)
		args = append(args, info.City)
	}
	if info.ISP != `` && info.ISP != `0` {
		formats = append(formats, `"线路":%q`)
		args = append(args, info.ISP)
	}
	return fmt.Sprintf(`{`+strings.Join(formats, `,`)+`}`, args...)
}
