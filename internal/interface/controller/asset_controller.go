package controller

import (
	"encoding/json"
	"go-portfolio-tracker/internal/usecase/query"
	"net/http"
)

type AssetController struct {
	getAllAssetsQuery *query.GetAllAssetsQuery
}

func NewAssetController(getAllAssetsQuery *query.GetAllAssetsQuery) *AssetController {
	return &AssetController{getAllAssetsQuery: getAllAssetsQuery}
}

func (h *AssetController) GetAllAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := h.getAllAssetsQuery.Handle()
	if err != nil {
		http.Error(w, "Failed to fetch assets", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(assets)
}
