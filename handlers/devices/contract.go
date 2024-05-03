package devicehandler

import (
	"github.com/gin-gonic/gin"
)

type StockHandler interface {
	CreateStock(ctx *gin.Context)
	GetStockByID(ctx *gin.Context)
	UpdateStock(ctx *gin.Context)
	DeleteStock(ctx *gin.Context)
}

type stockHandler struct {
	stockService stockservice.StockService
}

func NewStockHandler(stockService stockservice.stockService) *stockHandler {
	return &stockHandler{stockService}
}
