package global

type card struct {
	ID   int    `gorm:"primary_key;auto_increase"`
	name string `json:"name"`
}

type orders struct {
	ID      int     `gorm:"primary_key;auto_increase"`
	trader  string  `json:"trader"`
	card_ID int     `json:"card_ID"`
	price   float32 `json:"price"`
}

type trades struct {
	ID      int     `gorm:"primary_key;auto_increase"`
	buy     string  `json:"buy"`
	sell    string  `json:"sell"`
	card_ID int     `json:"card_ID"`
	price   float32 `json:"price"`
}
