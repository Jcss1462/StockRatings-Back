package dtos

type ApiStockResponseDto struct {
	Items     []StockDto
	Next_page string
}
