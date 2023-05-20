package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Marif226/product-crud/internal/model"
	mocks_repo "github.com/Marif226/product-crud/internal/repository/mocks"
	"github.com/Marif226/product-crud/internal/service"
	"github.com/Marif226/product-crud/pkg/router"
)

func TestHandler_CreatePurchase(t *testing.T) {

	testTable := []struct {
		name               string
		reqBody            model.Purchase
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "OK",
			reqBody: model.Purchase{
				Name:        "Test Purchase",
				Description: "Test description",
				Quantity:    1,
				Price:       10,
				BuyerID:     1,
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   fmt.Sprintln("1"),
		},
		{
			name:               "Empty fields",
			reqBody:            model.Purchase{},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   fmt.Sprintln("error: empty fields"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r := mocks_repo.NewMockRepository()
			s := service.New(r)
			h := New(s)

			reqBodyJSON, err := json.Marshal(testCase.reqBody)
			if err != nil {
				t.Fatal(err)
			}

			// Test Server
			router := router.New()
			router.Add("POST /purchases", h.CreatePurchase)

			// Test Request
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/purchases", strings.NewReader(string(reqBodyJSON)))

			// Make Request
			router.ServeHTTP(rr, req)

			// Assert
			if testCase.expectedStatusCode != rr.Code {
				t.Errorf("wrong response code: got %v, wanted %v", rr.Code, testCase.expectedStatusCode)
			}

			if testCase.expectedResponse != rr.Body.String() {
				t.Errorf("wrong response body: got %v, wanted %v", rr.Body.String(), testCase.expectedResponse)
			}
		})
	}
}

func TestHandler_GetAllPurchases(t *testing.T) {

	testTable := []struct {
		name               string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:               "OK",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintln(`[{"ID":1,"Name":"Test Purchase","Description":"Test Description","Quantity":1,"Price":10,"Buyer":{"id":"1","name":"Test Buyer","contact":"test.buyer@example.com"}}]`),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r := mocks_repo.NewMockRepository()
			s := service.New(r)
			h := New(s)

			// Test Server
			router := router.New()
			router.Add("GET /purchases", h.GetAllPurchases)

			// Test Request
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/purchases", nil)

			// Make Request
			router.ServeHTTP(rr, req)

			// Assert
			if testCase.expectedStatusCode != rr.Code {
				t.Errorf("wrong response code: got %v, wanted %v", rr.Code, testCase.expectedStatusCode)
			}

			if testCase.expectedResponse != rr.Body.String() {
				t.Errorf("wrong response body: got %v, wanted %v", rr.Body.String(), testCase.expectedResponse)
			}
		})
	}
}

func TestHandler_GetPurchaseById(t *testing.T) {

	testTable := []struct {
		name               string
		queryParam         map[string]string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "OK",
			queryParam: map[string]string{
				"id": "1",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintln(`{"ID":1,"Name":"Test Purchase","Description":"Test Description","Quantity":1,"Price":10,"Buyer":{"id":"1","name":"Test Buyer","contact":"test.buyer@example.com"}}`),
		},
		{
			name: "NO ID",
			queryParam: map[string]string{
				"id": "17",
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   fmt.Sprintln(`purchase with given id does not exist`),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r := mocks_repo.NewMockRepository()
			s := service.New(r)
			h := New(s)

			// Test Server
			router := router.New()
			router.Add("GET /purchases/get", h.GetPurchaseById)

			// Test Request
			rr := httptest.NewRecorder()
			query := fmt.Sprintf("/purchases/get?id=%s", testCase.queryParam["id"])
			req := httptest.NewRequest("GET", query, nil)

			// Make Request
			router.ServeHTTP(rr, req)

			// Assert
			if testCase.expectedStatusCode != rr.Code {
				t.Errorf("wrong response code: got %v, wanted %v", rr.Code, testCase.expectedStatusCode)
			}

			if testCase.expectedResponse != rr.Body.String() {
				t.Errorf("wrong response body: got %v, wanted %v", rr.Body.String(), testCase.expectedResponse)
			}
		})
	}
}

func TestHandler_UpdatePurchase(t *testing.T) {

	testTable := []struct {
		name               string
		reqBody            model.Purchase
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "OK",
			reqBody: model.Purchase{
				ID:          1,
				Name:        "Test Purchase",
				Description: "Test description",
				Quantity:    1,
				Price:       10,
				BuyerID:     1,
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintln(`{"ID":1,"Name":"Test Purchase","Description":"Test Description","Quantity":1,"Price":10,"Buyer":{"id":"1","name":"Test Buyer","contact":"test.buyer@example.com"}}`),
		},
		{
			name: "Empty fields",
			reqBody: model.Purchase{
				ID: 1,
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   fmt.Sprintln("error: empty fields"),
		},
		{
			name: "NO ID",
			reqBody: model.Purchase{
				Name:        "Test Purchase",
				Description: "Test description",
				Quantity:    1,
				Price:       10,
				BuyerID:     1,
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   fmt.Sprintln("purchase with given id does not exist"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r := mocks_repo.NewMockRepository()
			s := service.New(r)
			h := New(s)

			reqBodyJSON, err := json.Marshal(testCase.reqBody)
			if err != nil {
				t.Fatal(err)
			}

			// Test Server
			router := router.New()
			router.Add("PUT /purchases", h.UpdatePurchase)

			// Test Request
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/purchases", strings.NewReader(string(reqBodyJSON)))

			// Make Request
			router.ServeHTTP(rr, req)

			// Assert
			if testCase.expectedStatusCode != rr.Code {
				t.Errorf("wrong response code: got %v, wanted %v", rr.Code, testCase.expectedStatusCode)
			}

			if testCase.expectedResponse != rr.Body.String() {
				t.Errorf("wrong response body: got %v, wanted %v", rr.Body.String(), testCase.expectedResponse)
			}
		})
	}
}

func TestHandler_DeletePurchase(t *testing.T) {

	testTable := []struct {
		name               string
		queryParam         map[string]string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "OK",
			queryParam: map[string]string{
				"id": "1",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprint(`Purchase deleted!`),
		},
		{
			name: "NO ID",
			queryParam: map[string]string{
				"id": "17",
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   fmt.Sprintln(`purchase with given id does not exist`),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r := mocks_repo.NewMockRepository()
			s := service.New(r)
			h := New(s)

			// Test Server
			router := router.New()
			router.Add("DELETE /purchase", h.DeletePurchase)

			// Test Request
			rr := httptest.NewRecorder()
			query := fmt.Sprintf("/purchase?id=%s", testCase.queryParam["id"])
			req := httptest.NewRequest("DELETE", query, nil)

			// Make Request
			router.ServeHTTP(rr, req)

			// Assert
			if testCase.expectedStatusCode != rr.Code {
				t.Errorf("wrong response code: got %v, wanted %v", rr.Code, testCase.expectedStatusCode)
			}

			if testCase.expectedResponse != rr.Body.String() {
				t.Errorf("wrong response body: got %v, wanted %v", rr.Body.String(), testCase.expectedResponse)
			}
		})
	}
}
