package port

import (
	"repartners/internal/core/domain"

	"github.com/sirupsen/logrus"
)

type PackageService interface {
	CalculatePackages(items int, packagesCofig []int, logger *logrus.Entry) (domain.PackList, error)
}

var packageService PackageService

// GetInstance get an instance of the package service interface implementation
func GetPackageServiceInstance() PackageService {
	return packageService
}

// SetInstanceCb set a package service interface implementation
func SetPackageServiceInstance(instance PackageService) {
	packageService = instance
}
