package urltool

import (
	"net/url"
	"path"
	"github.com/zeromicro/go-zero/core/logx"
	"errors"
)

func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
		if err != nil {
			logx.Errorw("url.Parse failed", logx.LogField{Key: "err", Value: err})
			return "", err
		}
		if len(myUrl.Host) == 0 {
			return "", errors.New("invalid url: missing host")
		}
		baseUrl := path.Base(myUrl.Path)
		return baseUrl, nil
}