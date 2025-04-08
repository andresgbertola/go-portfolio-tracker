package repository

import (
	"go-portfolio-tracker/internal/domain"

	"gorm.io/gorm"
)

type AssetRepositoryGorm struct {
	db *gorm.DB
}

func NewAssetRepositoryGorm(db *gorm.DB) *AssetRepositoryGorm {
	return &AssetRepositoryGorm{db: db}
}

func (a *AssetRepositoryGorm) SaveAsset(asset domain.Asset) error {
	return a.db.Create(&asset).Error
}

func (a *AssetRepositoryGorm) GetAllAssets() ([]domain.Asset, error) {

	var assets []domain.Asset
	err := a.db.Find(&assets).Error

	return assets, err
}
