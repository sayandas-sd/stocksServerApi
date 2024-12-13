package model

type Stock struct {
	StockId int64  `json:"stockid"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Company string `json:"company"`
}
