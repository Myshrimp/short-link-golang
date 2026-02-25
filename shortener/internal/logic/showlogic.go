package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// Replace all errors.New("404 short url not found") with Error404
var Error404 = errors.New("404 short url not found")

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// redirect short url to long url
	// 1. get long url from database by short url
	// 1.0 bloom filter to prevent cache penetration
	// if short url is not in bloom filter, return 404 not found
	// a. based on memory, but records will be lost when the service restarts
	// b. based on redis
	exist, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil {
		logx.Errorw("bloom filter exists failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	if !exist {
		return nil, Error404
	}

	fmt.Println("Querying database for short url:", req.ShortUrl)

	u, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: req.ShortUrl, Valid: true})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, Error404
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	resp = &types.ShowResponse{
		LongUrl: u.Lurl.String,
	}
	logx.Infow("Found long url", logx.LogField{Key: "Long url", Value: resp.LongUrl})
	return resp, nil
}
