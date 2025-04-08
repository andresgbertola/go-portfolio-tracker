package repository

import (
	"go-portfolio-tracker/internal/domain"
	"time"
)

type AssetRepositoryMock struct{}

func (a *AssetRepositoryMock) SaveAsset(asset domain.Asset) error {
	return nil
}

func (a *AssetRepositoryMock) GetAllAssets() ([]domain.Asset, error) {
	assets := []domain.Asset{
		{
			AssetName:          "BTC",
			Quantity:           1,
			ExchangedQuantity:  90000.0,
			ExchangedAssetName: "Ethereum",
			UnitExchangeValue:  90000.0,
			EntryDate:          time.Now(),
		},
	}
	return assets, nil
}
