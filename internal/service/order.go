package service

import "github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/internal/model"

type CreateOrderRequest struct {
	TraderID  int     `form:"trader_id"`
	CardID    int     `form:"card_id"`
	Price     float32 `form:"price" binding:"min=0 max=10"`
	OrderType string  `form:"order_type"`
	Status    bool
}

type QueryOrderRequest struct {
	TraderID  int     `form:"trader_id"`
	CardID    int     `form:"card_id"`
	Price     float32 `form:"price" binding:"min=0 max=10"`
	OrderType string  `form:"order_type"`
	Status    bool
}

func (svc *Service) CreateOrder(param *CreateOrderRequest) error {
	return svc.dao.CreateOrder(param.TraderID, param.CardID, param.Price, param.OrderType)
}

func (svc *Service) QueryMinSellOrder(param *CreateOrderRequest) (*model.Order, error) {
	return svc.dao.QueryMinSellOrder(param.CardID)
}

func (svc *Service) QueryMaxBuyOrder(param *CreateOrderRequest) (*model.Order, error) {
	return svc.dao.QueryMaxBuyOrder(param.CardID)
}

func (svc *Service) QueryTraderOrder(param *QueryOrderRequest) ([]*model.Order, error) {
	return svc.dao.QueryTraderOrder(param.TraderID)
}

func (svc *Service) CompleteOrderAndCreateTrade(Order *model.Order, param *CreateOrderRequest) error {
	return svc.dao.CompleteOrderAndCreateTrade(Order.ID, Order.Price, param.TraderID, param.CardID, param.Price, param.OrderType)
}
