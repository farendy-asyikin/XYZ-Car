package schemas

type CreateStockRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type UpdateStockRequest struct {
	Name *string `json:"name" binding:"required"`
	Type *string `json:"type" binding:"required"`
}

type DetailStockResponse struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	IsActive bool   `json:"is_active"`
}
