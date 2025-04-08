package controller

import (
	"encoding/json"
	"go-portfolio-tracker/internal/interface/dto"
	"go-portfolio-tracker/internal/usecase/command"
	"go-portfolio-tracker/internal/usecase/query"
	"net/http"
)

type AssetController struct {
	getAllAssetsQuery     *query.GetAllAssetsQuery
	createNewAssetCommand *command.CreateNewAssetCommand
}

func NewAssetController(getAllAssetsQuery *query.GetAllAssetsQuery,
	createNewAssetCommand *command.CreateNewAssetCommand) *AssetController {
	return &AssetController{
		getAllAssetsQuery:     getAllAssetsQuery,
		createNewAssetCommand: createNewAssetCommand,
	}
}

func (h *AssetController) GetAllAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := h.getAllAssetsQuery.Handle()
	if err != nil {
		http.Error(w, "Failed to fetch assets", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(assets)
}

func (h *AssetController) CreateNewAsset(w http.ResponseWriter, r *http.Request) {
	var assetDto dto.AssetDTO
	err := json.NewDecoder(r.Body).Decode(&assetDto)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	assetDtoResponse, err := h.createNewAssetCommand.Handle(assetDto)
	if err != nil {
		http.Error(w, "Failed to create asset", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assetDtoResponse)
}
