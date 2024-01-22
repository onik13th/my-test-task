package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/onik13th/my-test-task/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"name":       "Лев",
				"surname":    "Толстой",
				"patronymic": "Николаевич",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"name":       "123",
				"surname":    "321",
				"patronymic": "312",
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "valid",
			payload: map[string]string{
				"name":       "Фридрих",
				"surname":    "Энгельс",
				"patronymic": "",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"name":       "Стендаль",
				"surname":    "",
				"patronymic": "",
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest("POST", "/api/users/", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleFoundAllUsers(t *testing.T) {
	s := newServer(teststore.New())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users/", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

//func TestServer_HandleFoundById(t *testing.T) {
//	s := newServer(teststore.New())
//
//	testCases := []struct {
//		name         string
//		id           string
//		expectedCode int
//	}{
//		{
//			name:         "valid id",
//			id:           "1",
//			expectedCode: http.StatusOK,
//		},
//		{
//			name:         "invalid id",
//			id:           "invalid",
//			expectedCode: http.StatusInternalServerError,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			rec := httptest.NewRecorder()
//			req, _ := http.NewRequest("GET", "/api/users/"+tc.id, nil)
//			s.ServeHTTP(rec, req)
//			assert.Equal(t, tc.expectedCode, rec.Code)
//		})
//	}
//}
