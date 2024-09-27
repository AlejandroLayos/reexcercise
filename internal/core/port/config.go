package port

import (
	"context"
	"repartners/internal/core/domain"

	logrus "github.com/sirupsen/logrus"
)

type ConfigService interface {
	GetConfigOrdered(ctx context.Context, logger *logrus.Entry) (*domain.Config, error)
	UpdateConfig(ctx context.Context, config *domain.Config, logger *logrus.Entry) error
}

var configService ConfigService

// GetInstance get an instance of the package service interface implementation
func GetConfigServiceInstance() ConfigService {
	return configService
}

func SetConfigServiceInstance(instance ConfigService) {
	configService = instance
}
