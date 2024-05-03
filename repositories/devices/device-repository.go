package devicerepository

import (
	"errors"
	"gorm.io/gorm"
	"main.go/models"
)

func (r *stockRepository) CreateStock(sensor models.Stock) (*models.Stock, error) {

	err := r.db.Create(&sensor).Error
	if err != nil {
		return nil, err
	}

	return &sensor, nil
}

func (r *stockRepository) GetStockByID(ID string) (*models.Stock, error) {
	var stock models.Stock
	err := r.db.First(&stock, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("stock not found")
		}

		return nil, err
	}

	return &stock, nil
}

func (r *stockRepository) UpdateStock(stock models.Stock) (*models.Stock, error) {
	err := r.db.Model(&stock).Updates(map[string]any{
		"name":      stock.Brand,
		"type":      stock.Type,
		"is_active": stock.IsActive,
	}).Error
	if err != nil {
		return nil, err
	}

	return &stock, nil
}

func (r *stockRepository) DeleteStock(ID string) error {
	var stock models.Stock
	err := r.db.Model(&stock).Where("id = ?", ID).Updates(map[string]interface{}{"is_active": false}).Error
	if err != nil {
		return err
	}

	return nil
}
