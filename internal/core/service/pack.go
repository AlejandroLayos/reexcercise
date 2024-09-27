package service

import (
	"repartners/internal/core/domain"

	"github.com/sirupsen/logrus"
)

type PackServiceImpl struct {
}

func (p *PackServiceImpl) CalculatePackages(items int, packagesCofig []int, logger *logrus.Entry) (domain.PackList, error) {
	logger.Debug("CalculatePackages")

	var result []domain.Packs
	remainingItems := items

	for _, packSize := range packagesCofig {
		if remainingItems == 0 {
			break
		}
		quantity := remainingItems / packSize
		if quantity > 0 {
			result = append(result, domain.Packs{Quantity: quantity, Value: packSize})
			remainingItems -= quantity * packSize
		}
	}

	// If there are remaining items that couldn't be packed, add the smallest pack
	if remainingItems > 0 {
		result = append(result, domain.Packs{Quantity: 1, Value: packagesCofig[len(packagesCofig)-1]})
	}

	return domain.PackList{Packs: result}, nil
}
