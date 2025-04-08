package domain

import "time"

type Asset struct {
	AssetName          string
	Quantity           float64
	ExchangedQuantity  float64
	ExchangedAssetName string
	UnitExchangeValue  float64
	EntryDate          time.Time
}
