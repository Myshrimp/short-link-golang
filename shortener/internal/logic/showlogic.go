package logic

import (
	"context"
	"database/sql"
	"errors"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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
	u, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: req.ShortUrl, Valid: true})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("404 short url not found")
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
