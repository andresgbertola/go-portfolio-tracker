package dto

import (
	"go-portfolio-tracker/internal/domain"
	"time"
)

type AssetDTO struct {
	AssetName          string  `json:"assetName"`
	Quantity           float64 `json:"quantity"`
	ExchangedQuantity  float64 `json:"exchangedQuantity"`
	ExchangedAssetName string  `json:"exchangedAssetName"`
	UnitExchangeValue  float64 `json:"unitExchangeValue"`
	EntryDate          string  `json:"entryDate"`
}

func ToDto(a domain.Asset) AssetDTO {
	return AssetDTO{
		AssetName:          a.AssetName,
		Quantity:           a.Quantity,
		ExchangedQuantity:  a.ExchangedQuantity,
		ExchangedAssetName: a.ExchangedAssetName,
		UnitExchangeValue:  a.UnitExchangeValue,
		EntryDate:          a.EntryDate.Format(time.RFC3339),
	}
}
