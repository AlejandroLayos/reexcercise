package service

import (
	"context"
	"encoding/json"
	"errors"
	"repartners/internal/core/domain"
	"repartners/internal/core/port"

	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	gomock "go.uber.org/mock/gomock"
)

func TestServicesConfigSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Services Test Suite")
}

var _ = Describe("Config Services Tests Suite", func() {

	var (
		mockCtrl        *gomock.Controller
		mockConfigCache *port.MockCacheRepository
		service         *ConfigServiceImpl
		logger          *logrus.Entry
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockConfigCache = port.NewMockCacheRepository(mockCtrl)

		logger = logrus.New().WithContext(context.Background())
		service = &ConfigServiceImpl{CacheRepository: mockConfigCache}

	})

	AfterEach(func() {
		defer mockCtrl.Finish()
	})
	Context("GetConfigOrdered method", func() {
		jsonConfigReturned, _ := json.Marshal(domain.Config{PacksConfig: []int{1, 2, 3}})
		configCheck := domain.Config{PacksConfig: []int{3, 2, 1}}
		It("Success", func() {
			mockConfigCache.EXPECT().Get(gomock.Any(), gomock.Any()).Return(jsonConfigReturned, nil)
			configOrdered, err := service.GetConfigOrdered(context.Background(), logger)
			Expect(configOrdered).To(Equal(&configCheck))
			Expect(err).To(BeNil())
		})
		It("Error", func() {
			mockConfigCache.EXPECT().Get(gomock.Any(), gomock.Any()).Return(jsonConfigReturned, errors.New("error"))
			configOrdered, err := service.GetConfigOrdered(context.Background(), logger)
			Expect(configOrdered).To(BeNil())
			Expect(err).To(HaveOccurred())
		})

	})

})
