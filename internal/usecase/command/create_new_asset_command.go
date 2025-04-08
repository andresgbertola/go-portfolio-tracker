package command

import (
	"go-portfolio-tracker/internal/domain"
	"go-portfolio-tracker/internal/interface/dto"
	"go-portfolio-tracker/internal/interface/repository"
	"time"
)

type CreateNewAssetCommand struct {
	repo repository.AssetRepository
}

func NewCreateNewAssetCommand(repo repository.AssetRepository) *CreateNewAssetCommand {
	return &CreateNewAssetCommand{repo: repo}
}

func (c CreateNewAssetCommand) Handle(assetDto dto.AssetDTO) (dto.AssetDTO, error) {
	entryDate, err := time.Parse(time.RFC3339, assetDto.EntryDate)
	if err != nil {
		return dto.AssetDTO{}, err
	}

	asset := domain.Asset{
		AssetName:          assetDto.AssetName,
		Quantity:           assetDto.Quantity,
		ExchangedAssetName: assetDto.ExchangedAssetName,
		ExchangedQuantity:  assetDto.ExchangedQuantity,
		UnitExchangeValue:  assetDto.UnitExchangeValue,
		EntryDate:          entryDate,
	}

	c.repo.SaveAsset(asset)

	return dto.ToDto(asset), nil
}
