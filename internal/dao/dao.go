package dao

import (
	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/internal/model"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) CreateOrder(traderID int, cardID int, price float32, orderType string) error {
	order := model.Order{
		CardID:   cardID,
		TraderID: traderID,
		Price:    price,
		Type:     orderType,
		Status:   false,
	}
	return order.Create(d.engine)
}

func (d *Dao) CreateTrade(cardID int, price float32) error {
	trade := model.Trade{
		CardID: cardID,
		Price:  price,
	}
	return trade.Create(d.engine)
}

func (d *Dao) CreateTrader() error {
	trader := model.Trader{}
	return trader.Create(d.engine)
}

func (d *Dao) QueryMinSellOrder(cardID int) (*model.Order, error) {
	order := model.Order{}
	card := model.Card{
		ID: cardID,
	}
	return order.GetMinSellWithCard(d.engine, card)
}

func (d *Dao) QueryMaxBuyOrder(cardID int) (*model.Order, error) {
	order := model.Order{}
	card := model.Card{
		ID: cardID,
	}
	return order.GetMaxBuyWithCard(d.engine, card)
}

func (d *Dao) QueryTraderOrder(traderID int) ([]*model.Order, error) {
	order := model.Order{}
	trader := model.Trader{
		ID: traderID,
	}
	return order.ListWithTrader(d.engine, trader)
}

func (d *Dao) QueryCardTrade(cardID int) ([]*model.Trade, error) {
	trade := model.Trade{}
	card := model.Card{
		ID: cardID,
	}
	return trade.ListWithCard(d.engine, card)
}

func (d *Dao) CompleteOrderAndCreateTrade(orderID int, tradePrice float32, traderID int, cardID int, price float32, orderType string) error {
	order := model.Order{
		ID: orderID,
	}
	newOrder := model.Order{
		CardID:   cardID,
		TraderID: traderID,
		Price:    price,
		Type:     orderType,
		Status:   true,
	}
	trade := model.Trade{
		CardID: cardID,
		Price:  tradePrice,
	}
	if err := order.UpdateStatus(d.engine, true); err != nil {
		return err
	}
	if err := newOrder.Create(d.engine); err != nil {
		order.UpdateStatus(d.engine, false)
		return err
	}
	if err := trade.Create(d.engine); err != nil {
		return err
	}
	return nil
}
