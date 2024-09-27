package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"repartners/internal/core/domain"
	"repartners/internal/core/port"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	gomock "go.uber.org/mock/gomock"
)

func TestServicesSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Package Suite")
}

func createTestContext(method, url string, body interface{}) (*gin.Context, *httptest.ResponseRecorder, error) {
	// convert body to json in order to create the request
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, nil, err
	}

	// create a http request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// create a ResponseRecorder to capture the response
	w := httptest.NewRecorder()

	// create a gin context from the request and the ResponseRecorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	return c, w, nil
}

var _ = Describe("Handler Tests Suite", func() {

	var (
		mockCtrl           *gomock.Controller
		mockPackageService *port.MockPackageService
		mockConfigService  *port.MockConfigService
		packageHandler     *PackageHandler
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockPackageService = port.NewMockPackageService(mockCtrl)
		mockConfigService = port.NewMockConfigService(mockCtrl)

		port.SetPackageServiceInstance(mockPackageService)
		port.SetConfigServiceInstance(mockConfigService)
		packageHandler = NewPackageHandler(port.GetPackageServiceInstance(), port.GetConfigServiceInstance())

	})

	AfterEach(func() {
		defer mockCtrl.Finish()
	})
	Context("GetUser method", func() {
		packages := []domain.Packs{
			{
				Quantity: 10,
				Value:    500,
			},
			{
				Quantity: 20,
				Value:    1000,
			},
		}

		packageList := domain.PackList{
			Packs: packages,
		}

		It("Success get", func() {
			mockPackageService.EXPECT().CalculatePackages(gomock.Any(), gomock.Any(), gomock.Any()).Return(packageList, nil)
			mockConfigService.EXPECT().GetConfigOrdered(gomock.Any(), gomock.Any()).Return(&domain.Config{PacksConfig: []int{250, 500, 1000, 2000, 5000}}, nil)

			c, w, err := createTestContext(http.MethodGet, "/packs/1", nil)
			Expect(err).To(BeNil())
			c.Params = gin.Params{gin.Param{Key: "items", Value: "1"}}

			packageHandler.GetItemsHandler(c)
			Expect(w.Code).To(Equal(http.StatusOK))

			var response map[string]interface{}

			err = json.Unmarshal(w.Body.Bytes(), &response)
			Expect(err).To(BeNil())
		})
		It("Error internal GetConfigOrdered", func() {
			mockConfigService.EXPECT().GetConfigOrdered(gomock.Any(), gomock.Any()).Return(nil, errors.New("GetConfigOrdered internal error"))

			c, w, err := createTestContext(http.MethodGet, "/packs/2", nil)
			Expect(err).To(BeNil())
			c.Params = gin.Params{gin.Param{Key: "items", Value: "2"}}

			packageHandler.GetItemsHandler(c)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))

			var response map[string]interface{}

			err = json.Unmarshal(w.Body.Bytes(), &response)
			Expect(err).To(BeNil())
			Expect(response["error"]).To(Equal("GetConfigOrdered internal error"))
		})
		It("Error internal CalculatePackages", func() {
			mockConfigService.EXPECT().GetConfigOrdered(gomock.Any(), gomock.Any()).Return(&domain.Config{PacksConfig: []int{250, 500, 1000, 2000, 5000}}, nil)
			mockPackageService.EXPECT().CalculatePackages(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.PackList{}, errors.New("get internal error"))

			c, w, err := createTestContext(http.MethodGet, "/packs/3", nil)
			Expect(err).To(BeNil())
			c.Params = gin.Params{gin.Param{Key: "items", Value: "3"}}

			packageHandler.GetItemsHandler(c)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))

			var response map[string]interface{}

			err = json.Unmarshal(w.Body.Bytes(), &response)
			Expect(err).To(BeNil())
			Expect(response["error"]).To(Equal("get internal error"))
		})

	})
	Context("PutConfig method", func() {
		It("Success put", func() {
			mockConfigService.EXPECT().UpdateConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			c, w, err := createTestContext(http.MethodPut, "/config", domain.Config{PacksConfig: []int{250, 500, 1000, 2000, 5000}})
			Expect(err).To(BeNil())

			packageHandler.UpdateConfigHandler(c)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
		It("Error internal UpdateConfig", func() {
			mockConfigService.EXPECT().UpdateConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("UpdateConfig internal error"))

			c, w, err := createTestContext(http.MethodPut, "/config", domain.Config{PacksConfig: []int{250, 500, 1000, 2000, 5000}})
			Expect(err).To(BeNil())

			packageHandler.UpdateConfigHandler(c)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

})
