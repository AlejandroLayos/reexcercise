package http

import (
	"net/http"
	"strconv"

	"repartners/internal/core/domain"
	"repartners/internal/core/port"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// ProductHandler represents the HTTP handler for product-related requests
type PackageHandler struct {
	packageService port.PackageService
	configService  port.ConfigService
}

// NewProductHandler creates a new ProductHandler instance
func NewPackageHandler(service port.PackageService, configService port.ConfigService) *PackageHandler {
	return &PackageHandler{
		packageService: service,
		configService:  configService,
	}
}

// GetItemsHandler handles the GET /packs/:items endpoint
// GetItemsHandler handles the retrieval of a user by nickname
// It returns a number of packs.
// @Summary Retrieve a number of packs
// @Description retrieve a number of packs
// @ID get-items
// @Accept  json
// @Produce  json
// @Param items path int true "Items"
// @Success 200
// @Failure 400 "bad request"
// @Failure 500 "internal server error"
// @Router /{items} [get]
func (h *PackageHandler) GetItemsHandler(c *gin.Context) {
	logger := logrus.New().WithFields(logrus.Fields{
		"handler": "GetItemsHandler",
		"method":  "GET",
	})
	items := c.Param("items")
	if items == "" {
		logger.Error("items is required")
		c.JSON(http.StatusBadRequest, gin.H{"error to serve empty items": items})
		return
	}
	itemsToServe, err := strconv.Atoi(items)

	if err != nil {
		logger.Error("items is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": itemsToServe})
		return
	}
	config, err := h.configService.GetConfigOrdered(c, logger)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info(config)

	packageList, err := h.packageService.CalculatePackages(itemsToServe, config.PacksConfig, logger)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("package list retrieved")
	c.JSON(http.StatusOK, packageList)
}

// UpdateConfigHandler handles the PUT /packs/config endpoint
// UpdateConfigHandler handles the update of the config
// @Summary Update the config
// @Description update the config
// @ID update-config
// @Accept  json
// @Produce  json
// @Param config body domain.Config true "Config"
// @Success 200
// @Failure 400 "bad request"
// @Failure 500 "internal server error"
// @Router /config [put]
func (h *PackageHandler) UpdateConfigHandler(c *gin.Context) {
	logger := logrus.New().WithFields(logrus.Fields{
		"handler": "UpdateConfigHandler",
		"method":  "PUT",
	})
	var config domain.Config
	if err := c.ShouldBindJSON(&config); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(config.PacksConfig) == 0 {
		logger.Error("packs config is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "packs config is required"})
		return
	}
	err := h.configService.UpdateConfig(c, &config, logger)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("config updated")
	c.JSON(http.StatusOK, gin.H{"message": "config updated"})
}
