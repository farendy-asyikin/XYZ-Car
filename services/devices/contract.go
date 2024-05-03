package deviceservice

import (
	"main.go/models"
	"main.go/schemas"
)

type DeviceService interface {
	CreateStock(request schemas.CreateStockeRequest) (*models.Stock, error)
	UpdateStock(request schemas.UpdateStockRequest, device models.Stock) (*models.Stock, error)
	DeleteStock(ID string) error
	GetStockByID(ID string) (*models.Stock, *schemas.DetailStockResponse, error)
}

type stockService struct {
	stockRepository stockrepository.StockRepository
}

func NewStockService(
	stockRepository stockrepository.StockRepository,
) *stockService {
	return &stockService{
		stockRepository: stockRepository,
	}
}
