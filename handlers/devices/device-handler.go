package devicehandler

import (
	"github.com/gin-gonic/gin"
	"main.go/schemas"
	"main.go/utils"
	"net/http"
)

func (h *stockHandler) CreateStock(ctx *gin.Context) {
	var request schemas.CreateStockRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	user, err := h.stockService.CreateStock(request)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *stockHandler) GetStockByID(ctx *gin.Context) {
	id := ctx.Param("id")

	_, user, err := h.stockService.GetStockByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *stockHandler) UpdateStock(ctx *gin.Context) {
	id := ctx.Param("id")
	var request schemas.UpdateStockRequest

	err := ctx.ShouldBind(&request)
	if err != nil {

		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	user, _, err := h.stockService.GetStockByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	user, err = h.stockService.UpdateStock(request, *user)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *stockHandler) DeleteStock(ctx *gin.Context) {
	ID := ctx.Param("id")

	_, _, err := h.stockService.GetStockByID(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	err = h.stockService.DeleteStock(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", nil, nil)
}
