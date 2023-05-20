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

func TestHandler_CreateBuyer(t *testing.T) {

	testTable := []struct{
		name 				string
		reqBody 			model.Buyer
		expectedStatusCode 	int
		expectedResponse	string
	} {
		{
			name: "OK",
			reqBody: model.Buyer{
				Name: "John Tester",
				Contact: "john.tester@example.com",
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: fmt.Sprintln("1"),
		},
		{
			name: "Empty fields",
			reqBody: model.Buyer{
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: fmt.Sprintln("error: empty fields"),
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
			router.Add("POST /buyers", h.CreateBuyer)

			// Test Request
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/buyers", strings.NewReader(string(reqBodyJSON)))

			// Make Request
			router.ServeHTTP(rr, req)

			// Assert
			if testCase.expectedStatusCode != rr.Code {
				t.Errorf("wrong response code: got %v, wanted %v", rr.Code, testCase.expectedStatusCode)
			}

			if testCase.expectedResponse != rr.Body.String() {
				t.Errorf("empty response body: got %v, wanted %v", rr.Body.String(), testCase.expectedResponse)
			}
		})
	}
}

func TestHandler_GetAllBuyers(t *testing.T) {

	testTable := []struct{
		name 				string
		expectedStatusCode 	int
		expectedResponse	string
	} {
		{
			name: "OK",
			expectedStatusCode: http.StatusOK,
			expectedResponse: fmt.Sprintln(`[{"id":"1","name":"Test Buyer","contact":"test.buyer@example.com"}]`),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r := mocks_repo.NewMockRepository()
			s := service.New(r)
			h := New(s)

			// Test Server
			router := router.New()
			router.Add("GET /buyers", h.GetAllBuyers)

			// Test Request
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/buyers", nil)

			// Make Request
			router.ServeHTTP(rr, req)

			// Assert
			if testCase.expectedStatusCode != rr.Code {
				t.Errorf("wrong response code: got %v, wanted %v", rr.Code, testCase.expectedStatusCode)
			}

			if testCase.expectedResponse != rr.Body.String() {
				t.Errorf("empty response body: got %v, wanted %v", rr.Body.String(), testCase.expectedResponse)
			}
		})
	}
}

func TestHandler_GetBuyerById(t *testing.T) {

	testTable := []struct{
		name 				string
		queryParam			map[string]string
		expectedStatusCode 	int
		expectedResponse	string
	} {
		{
			name: "OK",
			queryParam: map[string]string{
				"id":"1",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: fmt.Sprintln(`{"id":"1","name":"Test Buyer","contact":"test.buyer@example.com"}`),
		},
		{
			name: "NO ID",
			queryParam: map[string]string{
				"id":"17",
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse: fmt.Sprintln(`buyer with given id does not exist`),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r := mocks_repo.NewMockRepository()
			s := service.New(r)
			h := New(s)

			// Test Server
			router := router.New()
			router.Add("GET /buyers/get", h.GetBuyerById)

			// Test Request
			rr := httptest.NewRecorder()
			query := fmt.Sprintf("/buyers/get?id=%s", testCase.queryParam["id"])
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

func TestHandler_Updatebuyer(t *testing.T) {

	testTable := []struct{
		name 				string
		reqBody 			model.Buyer
		expectedStatusCode 	int
		expectedResponse	string
	} {
		{
			name: "OK",
			reqBody: model.Buyer{
				ID: 1,
				Name: "John Tester",
				Contact: "john.tester@example.com",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: fmt.Sprintln(`{"id":"1","name":"Test Buyer","contact":"test.buyer@example.com"}`),
		},
		{
			name: "Empty fields",
			reqBody: model.Buyer{
				ID: 1,
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: fmt.Sprintln("error: empty fields"),
		},
		{
			name: "NO ID",
			reqBody: model.Buyer{
				
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: fmt.Sprintln("error: empty fields"),
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
			router.Add("PUT /buyers", h.UpdateBuyer)

			// Test Request
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/buyers", strings.NewReader(string(reqBodyJSON)))

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

func TestHandler_DeleteBuyer(t *testing.T) {

	testTable := []struct{
		name 				string
		queryParam			map[string]string
		expectedStatusCode 	int
		expectedResponse	string
	} {
		{
			name: "OK",
			queryParam: map[string]string{
				"id":"1",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: fmt.Sprintln(`Buyer and his purchases deleted!`),
		},
		{
			name: "NO ID",
			queryParam: map[string]string{
				"id":"17",
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse: fmt.Sprintln(`buyer with given id does not exist`),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			r := mocks_repo.NewMockRepository()
			s := service.New(r)
			h := New(s)

			// Test Server
			router := router.New()
			router.Add("DELETE /buyers", h.DeleteBuyer)

			// Test Request
			rr := httptest.NewRecorder()
			query := fmt.Sprintf("/buyers?id=%s", testCase.queryParam["id"])
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