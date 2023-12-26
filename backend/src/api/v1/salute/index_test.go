package salute

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateSaluteSubrouter(t *testing.T) {
	// define the table driven test cases
	testCases := []struct {
		name     string
		path     string
		expected int
	}{
		{
			name:     "Root Path",
			path:     "/v1/salute",
			expected: http.StatusOK,
		},
		{
			name:     "Not Found",
			path:     "/v1/notSalute",
			expected: http.StatusNotFound,
		},
	}

	// iterate over each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.path, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			CreateSaluteSubrouter(router.PathPrefix("/v1").Subrouter())

			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expected, rr.Result().StatusCode)
		})
	}
}
