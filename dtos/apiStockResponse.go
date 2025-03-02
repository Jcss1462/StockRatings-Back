package dtos

type ApiStockResponse struct {
	Items     []StockDto
	Next_page string
}
