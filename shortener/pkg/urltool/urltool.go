package urltool

import (
	"net/url"
	"path"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetBasePath(targeturl string) (string, error) {
	myUrl, err := url.Parse(targeturl)
		if err != nil {
			logx.Errorw("url.Parse failed", logx.LogField{Key: "err", Value: err})
			return "", err
		}
		baseUrl := path.Base(myUrl.Path)
		return baseUrl, nil
}