package query

import (
	"go-portfolio-tracker/internal/interface/dto"
	"go-portfolio-tracker/internal/interface/repository"
)

type GetAllAssetsQuery struct {
	repo repository.AssetRepository
}

func NewGetAllAssetsQuery(repo repository.AssetRepository) *GetAllAssetsQuery {
	return &GetAllAssetsQuery{repo: repo}
}

func (g GetAllAssetsQuery) Handle() ([]dto.AssetDTO, error) {

	g.repo.GetAllAssets()
	assets, err := g.repo.GetAllAssets()
	if err != nil {
		return nil, err
	}

	assetDTOs := make([]dto.AssetDTO, len(assets))
	for i, asset := range assets {
		assetDTOs[i] = dto.ToDto(asset)
	}
	return assetDTOs, nil
}
