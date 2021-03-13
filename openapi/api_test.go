package openapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var insertCountryINMobile, _ = ioutil.ReadFile("../testdata/insert-country-IN-Mobile.json")
var insertCountryINTablet, _ = ioutil.ReadFile("../testdata/insert-country-IN-Tablet.json")
var insertCountryUSTablet, _ = ioutil.ReadFile("../testdata/insert-country-US-Tablet.json")
var insertBadRequest, _ = ioutil.ReadFile("../testdata/insert-bad-request.json")
var insertCountryUSDeskAdditionalMatrics, _ = ioutil.ReadFile("../testdata/insert-country-US-Desk.json")
var getCountryIN, _ = ioutil.ReadFile("../testdata/get-country-IN.json")
var getCountryINMobile, _ = ioutil.ReadFile("../testdata/get-country-IN-Mobile.json")
var getCountryUK, _ = ioutil.ReadFile("../testdata/get-country-UK.json")
var getCountryUSMobile, _ = ioutil.ReadFile("../testdata/get-country-US-Mobile.json")

var router *mux.Router

func init() {
	InsertApiService := NewInsertApiService()
	InsertApiController := NewInsertApiController(InsertApiService)

	QueryApiService := NewQueryApiService()
	QueryApiController := NewQueryApiController(QueryApiService)

	router = NewRouter(InsertApiController, QueryApiController)
}

func TestAPI(t *testing.T) {
	t.Run("insert", func(fp *testing.T) {
		fp.Run("AddingCountryINWithMobileDataValidTest", AddCountryIN)
		fp.Run("AddingCountryINWithTabletDataValidTest", AddCountryINTablet)
		fp.Run("AddingCountryUSWithTabletDataValidTest", AddCountryUSTablet)
		fp.Run("AddingCountryUSWithAdditionalMetricsValidTest", AddCountryUSAdditionalMetrics)
		fp.Run("AddingCountryBadRequestInvalidTest", AddBadRequest)
	})
	t.Run("query", func(fp *testing.T) {
		fp.Run("GetDataWhenCountryExists", getCountryINTest)
		fp.Run("GetDataForDeviceInCountry", getCountryINMobileTest)
		fp.Run("GetDataWhenDeviceNotExist", getCountryUSMobileTest)
		fp.Run("GetDataWhenCountryNotExist", getCountryUKTest)
	})
}

func AddCountryIN(t *testing.T) {

	req, _ := http.NewRequest(http.MethodPut, "/v1/insert", bytes.NewReader(insertCountryINMobile))
	res := executeRequest(req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func AddCountryINTablet(t *testing.T) {

	req, _ := http.NewRequest(http.MethodPut, "/v1/insert", bytes.NewReader(insertCountryINTablet))
	res := executeRequest(req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func AddCountryUSTablet(t *testing.T) {

	req, _ := http.NewRequest(http.MethodPut, "/v1/insert", bytes.NewReader(insertCountryUSTablet))
	res := executeRequest(req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func AddCountryUSAdditionalMetrics(t *testing.T) {

	req, _ := http.NewRequest(http.MethodPut, "/v1/insert", bytes.NewReader(insertCountryUSDeskAdditionalMatrics))
	res := executeRequest(req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func AddBadRequest(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPut, "/v1/insert", bytes.NewReader(insertBadRequest))
	res := executeRequest(req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func getCountryINTest(t *testing.T) {

	req, _ := http.NewRequest(http.MethodPost, "/v1/query", bytes.NewReader(getCountryIN))
	res := executeRequest(req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func getCountryINMobileTest(t *testing.T) {

	req, _ := http.NewRequest(http.MethodPost, "/v1/query", bytes.NewReader(getCountryINMobile))
	res := executeRequest(req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func getCountryUSMobileTest(t *testing.T) {

	req, _ := http.NewRequest(http.MethodPost, "/v1/query", bytes.NewReader(getCountryUSMobile))
	res := executeRequest(req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func getCountryUKTest(t *testing.T) {

	req, _ := http.NewRequest(http.MethodPost, "/v1/query", bytes.NewReader(getCountryUK))
	res := executeRequest(req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}
