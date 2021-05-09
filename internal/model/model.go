package model

import (
	"gorm.io/gorm"
)

type Card struct {
	ID   int    `gorm:"primary_key;auto_increase"`
	Name string `json:"name"`
}

type Trader struct {
	ID int `gorm:"primary_key;auto_increase"`
}

type Order struct {
	ID        int     `gorm:"primary_key;auto_increase"`
	Price     float32 `json:"price"`
	OrderType string  `json:"order_type"`
	CardID    int     `json:"card_ID"`
	Card      Card
	TraderID  int `json:"trader_ID"`
	Trader    Trader
	Status    bool `json:"status"`
}

type Trade struct {
	ID     int     `gorm:"primary_key;auto_increase"`
	Price  float32 `json:"price"`
	CardID int     `json:"card_ID"`
	Card   Card
}

func (o Order) Create(db *gorm.DB) error {
	return db.Create(&o).Error
}

func (o Order) UpdateStatus(db *gorm.DB, status bool) error {
	return db.Model(&o).Where("status = ?", !status).Update("status", status).Error
}

func (t Trade) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (tr Trader) Create(db *gorm.DB) error {
	return db.Create(&tr).Error
}

func (o Order) ListWithTrader(db *gorm.DB, trader Trader) ([]*Order, error) {
	var orders []*Order
	if err := db.Where("trader_id = ?", trader.ID).Find(&orders).Limit(50).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (t Trade) ListWithCard(db *gorm.DB, card Card) ([]*Trade, error) {
	var trades []*Trade
	if err := db.Where("card_id = ?", card.ID).Find(&trades).Limit(50).Error; err != nil {
		return nil, err
	}
	return trades, nil
}

func (o Order) GetMaxBuyWithCard(db *gorm.DB, card Card) (*Order, error) {
	order := new(Order)
	if err := db.Where("caed_id = ? AND type = ? AND status = ?", card.ID, "buy", false).Order("price DESC").Find(&order).Limit(1).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (o Order) GetMinSellWithCard(db *gorm.DB, card Card) (*Order, error) {
	order := new(Order)
	if err := db.Where("caed_id = ? AND type = ? AND status = ?", card.ID, "sell", false).Order("price").Find(&order).Limit(1).Error; err != nil {
		return order, err
	}
	return order, nil
}
