package service

import (
	"context"

	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/global"
	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
