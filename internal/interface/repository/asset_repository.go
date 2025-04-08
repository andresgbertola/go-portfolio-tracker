package repository

import "go-portfolio-tracker/internal/domain"

type AssetRepository interface {
	SaveAsset(asset domain.Asset) error
	GetAllAssets() ([]domain.Asset, error)
}
