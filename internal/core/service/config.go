package service

import (
	"context"
	"encoding/json"
	"repartners/internal/core/domain"
	"repartners/internal/core/port"
	"sort"
	"time"

	"github.com/sirupsen/logrus"
)

type ConfigServiceImpl struct {
	CacheRepository port.CacheRepository
}

func (p *ConfigServiceImpl) GetConfigOrdered(ctx context.Context, logger *logrus.Entry) (*domain.Config, error) {
	logger.Debug("CalculatePackages")
	val, err := p.CacheRepository.Get(ctx, "config:1")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into a Config struct
	var config domain.Config
	if err := json.Unmarshal([]byte(val), &config); err != nil {
		return nil, err
	}
	sort.Sort(sort.Reverse(sort.IntSlice(config.PacksConfig)))

	return &config, nil
}

func (p *ConfigServiceImpl) UpdateConfig(ctx context.Context, config *domain.Config, logger *logrus.Entry) error {
	logger.Debug("UpdateConfig")
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}
	p.CacheRepository.Set(ctx, "config:1", jsonConfig, time.Hour*24)
	return nil
}
