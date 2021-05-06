package global

type Card struct {
	ID   int    `gorm:"primary_key;auto_increase"`
	Name string `json:"name"`
}

type Trader struct {
	ID int `gorm:"primary_key;auto_increase"`
}

type Order struct {
	ID       int     `gorm:"primary_key;auto_increase"`
	Price    float32 `json:"price"`
	CardID   int     `json:"card_ID"`
	Card     Card
	TraderID int `json:"trader_ID"`
	Trader   Trader
	status   bool `json:"status"`
}

type Trade struct {
	ID     int     `gorm:"primary_key;auto_increase"`
	Price  float32 `json:"price"`
	CardID int     `json:"card_ID"`
	Card   Card
}
