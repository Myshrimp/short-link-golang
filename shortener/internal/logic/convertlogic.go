package logic

import (
	"context"
	"database/sql"
	"fmt"
	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/model"
	"shortener/pkg/base62"
	"shortener/pkg/connect"
	"shortener/pkg/md5"
	"shortener/pkg/urltool"

	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Input: longUrl && Output: shortUrl
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1. validate the long url
	// 1.1 long url can not be empty
	// we use validator package to validate the long url, if the long url is empty, return an error
	// 1.2 long url must be a valid url, we use http status code
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("invalid url")
	}
	// 1.3 check if long url already exists in the database
	// 1.3.1 generate md5 hash of the long url
	md5Value := md5.Sum([]byte(req.LongUrl))
	// 1.3.2 check if the hash exists in the database
	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{String: md5Value, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, fmt.Errorf("This link is already converted to %s", u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	// 1.4 input url can't be a short url, if it is a short url, return an error
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urltool.GetBasePath failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	if basePath != "" {
		_, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
		if err != sqlx.ErrNotFound {
			if err == nil {
				return nil, errors.New("This link is already a short url")
			}
			logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err})
			return nil, err
		}
	}

	var shortUrl string

	// 2. generate a unique id for the long url
	// for each convert request, we use REPLACE INTO to insert a new record into the database,
	//  and the id will be auto-incremented by the database, so we can get a unique id for each long url
	seq, err := l.svcCtx.Sequence.Next()
	if err != nil {
		logx.Errorw("Sequence.Next failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	logx.Infow("get sequence", logx.LogField{Key: "seq", Value: seq})

	// 3. convert the long url to a short url and return it
	// 3.1 to improve security, we can shuffle base62str to make it more difficult to guess the short url
	// 3.2 blacklist some short urls that may be offensive or inappropriate,
	// we can use a list of blacklisted words and check if the generated short url contains any of the blacklisted words,
	// if it does, we can generate a new short url until it doesn't contain any blacklisted words
	shortUrl = base62.Int2String(seq)
	if l.svcCtx.ShortUrlBlacklist != nil {
		if _, ok := l.svcCtx.ShortUrlBlacklist[shortUrl]; ok {
			logx.Infow("short url is blacklisted", logx.LogField{Key: "shortUrl", Value: shortUrl})
			return nil, errors.New("generated short url is blacklisted, please try again")
		}
	}
	logx.Infow("generate short url", logx.LogField{Key: "shortUrl", Value: shortUrl})
	// 4. store the mapping of the long url and the short url in the database
	if _, err := l.svcCtx.ShortUrlModel.Insert(l.ctx, &model.ShortUrlMap{
		Lurl: sql.NullString{String: req.LongUrl, Valid: true},
		Surl: sql.NullString{String: shortUrl, Valid: true},
		Md5:  sql.NullString{String: md5Value, Valid: true},
	}); err != nil {
		logx.Errorw("ShortUrlModel.Insert failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}

	// 4.1 add short url to bloom filter
	if err := l.svcCtx.Filter.Add([]byte(shortUrl)); err != nil {
		logx.Errorw("bloom filter add failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}

	// 5. return the short url
	shortUrl = l.svcCtx.Config.Domain + "/" + shortUrl
	resp = &types.ConvertResponse{
		ShortUrl: shortUrl,
	}
	return resp, nil
}
