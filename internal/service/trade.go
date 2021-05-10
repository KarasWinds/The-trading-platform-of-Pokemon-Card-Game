package service

import (
	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/internal/model"
)

type QueryTradeRequest struct {
	CardID int     `form:"card_id"`
	Price  float32 `form:"price" binding:"gte=1,lte=10"`
}

func (svc *Service) QueryCardTrade(param *QueryTradeRequest) ([]*model.Trade, error) {
	return svc.dao.QueryCardTrade(param.CardID)
}
