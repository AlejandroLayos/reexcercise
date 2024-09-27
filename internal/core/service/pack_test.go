package service

import (
	"context"
	"repartners/internal/core/domain"

	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	gomock "go.uber.org/mock/gomock"
)

func TestServicesPackSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pack Services Test Suite")
}

var _ = Describe("Pack Services Tests Suite", func() {

	var (
		mockCtrl *gomock.Controller
		service  *PackServiceImpl
		logger   *logrus.Entry
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		logger = logrus.New().WithContext(context.Background())
		service = &PackServiceImpl{}

	})

	AfterEach(func() {
		defer mockCtrl.Finish()
	})
	Context("CalculatePackages method", func() {
		packListCheck := domain.PackList{
			Packs: []domain.Packs{
				{Quantity: 3, Value: 3},
				{Quantity: 1, Value: 1},
			},
		}
		It("Success creation", func() {
			packList, err := service.CalculatePackages(10, []int{3, 2, 1}, logger)
			Expect(packList).To(Equal(packListCheck))
			Expect(err).To(BeNil())
		})

	})

})
