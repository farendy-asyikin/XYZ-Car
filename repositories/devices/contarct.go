package devicerepository

import (
	"gorm.io/gorm"
	"main.go/models"
)

type StockRepository interface {
	CreateStock(device models.Stock) (*models.Stock, error)
	GetStockByID(ID string) (*models.Stock, error)
	UpdateStock(device models.Stock) (*models.Stock, error)
	DeleteStock(ID string) error
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *stockRepository {
	return &stockRepository{db}
}
