// testutils/HTTPTester.go
package testutil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

// HTTPTester is a test utility for making HTTP requests to a Gin router.
type HTTPTester struct {
	t        *testing.T
	router   *gin.Engine
	method   string
	path     string
	body     io.Reader
	headers  map[string]string
	response *httptest.ResponseRecorder
}

// Route is a struct that represents a route in a Gin router.
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func NewHTTPTester(t *testing.T, register func(api *gin.RouterGroup)) *HTTPTester {
	t.Helper()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	// Register routes
	register(router.Group("/api"))

	return &HTTPTester{
		t:        t,
		router:   router,
		headers:  map[string]string{"Content-Type": "application/json"},
		response: httptest.NewRecorder(),
	}
}

func (a *HTTPTester) RegisterRoute(httpMethod string, path string, handler gin.HandlerFunc) *HTTPTester {
	// dynamically register routes based on the HTTP method
	switch httpMethod {
	case http.MethodGet:
		a.router.GET(path, handler)
	case http.MethodPost:
		a.router.POST(path, handler)
	case http.MethodPut:
		a.router.PUT(path, handler)
	case http.MethodDelete:
		a.router.DELETE(path, handler)
	case http.MethodPatch:
		a.router.PATCH(path, handler)
	case http.MethodOptions:
		a.router.OPTIONS(path, handler)
	case http.MethodHead:
		a.router.HEAD(path, handler)
	case "ANY": // Use gin.Any() for custom HTTP methods
		a.router.Any(path, handler)
	default:
		// Support for custom HTTP methods
		a.router.Handle(httpMethod, path, handler)
	}

	return a
}

func (a *HTTPTester) RegisterRoutes(routes []Route) *HTTPTester {
	for _, route := range routes {
		a.RegisterRoute(route.Method, route.Path, route.Handler)
	}

	return a
}

func (a *HTTPTester) WithHeader(key, value string) *HTTPTester {
	a.headers[key] = value

	return a
}

func (a *HTTPTester) Get(path string, params map[string]string) *HTTPTester {
	a.method = "GET"
	a.path = path

	if len(params) > 0 {
		a.path += "?" + encodeParams(params)
	}

	return a.execute()
}

func (a *HTTPTester) Post(path string) *HTTPTester {
	a.method = "POST"
	a.path = path

	return a.execute()
}

func (a *HTTPTester) Put(path string) *HTTPTester {
	a.method = "PUT"
	a.path = path

	return a.execute()
}

func (a *HTTPTester) Delete(path string) *HTTPTester {
	a.method = "DELETE"
	a.path = path

	return a.execute()
}

func (a *HTTPTester) SeeStatus(expectedStatus int) *HTTPTester {
	if a.response.Code != expectedStatus {
		a.t.Fatalf("Expected status code %d, but got %d", expectedStatus, a.response.Code)
	}

	return a
}

// DumpResponse logs the response body to the test output.
func (a *HTTPTester) DumpResponse() *HTTPTester {
	a.t.Logf("Response status: %d", a.response.Code)
	a.t.Logf("Response: %s", a.response.Body.String())

	return a
}

func (a *HTTPTester) WithJSONBody(body interface{}) *HTTPTester {
	jsonData, err := json.Marshal(body)
	if err != nil {
		a.t.Fatalf("Failed to marshal body: %v", err)
	}

	a.body = bytes.NewReader(jsonData)

	return a
}

func (a *HTTPTester) reset() {
	a.response = httptest.NewRecorder()
}

func (a *HTTPTester) execute() *HTTPTester {
	a.reset()
	req := httptest.NewRequest(a.method, a.path, a.body)

	for key, value := range a.headers {
		req.Header.Set(key, value)
	}

	a.router.ServeHTTP(a.response, req)
	return a
}

func encodeParams(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}

	return values.Encode()
}
