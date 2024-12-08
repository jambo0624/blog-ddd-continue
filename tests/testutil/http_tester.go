// testutils/HttpTester.go
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

// HttpTester is a test utility for making HTTP requests to a Gin router.
type HttpTester struct {
	t        *testing.T
	router   Router
	engine   *gin.Engine
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

func NewHttpTester(t *testing.T, router Router) *HttpTester {
	t.Helper()
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	
	// Register routes
	router.Register(engine.Group("/api"))

	return &HttpTester{
		t:        t,
		router:   router,
		engine:   engine,
		headers:  map[string]string{"Content-Type": "application/json"},
		response: httptest.NewRecorder(),
	}
}

func (a *HttpTester) RegisterRoute(httpMethod string, path string, handler gin.HandlerFunc) *HttpTester {
	// dynamically register routes based on the HTTP method
	switch httpMethod {
	case http.MethodGet:
		a.engine.GET(path, handler)
	case http.MethodPost:
		a.engine.POST(path, handler)
	case http.MethodPut:
		a.engine.PUT(path, handler)
	case http.MethodDelete:
		a.engine.DELETE(path, handler)
	case http.MethodPatch:
		a.engine.PATCH(path, handler)
	case http.MethodOptions:
		a.engine.OPTIONS(path, handler)
	case http.MethodHead:
		a.engine.HEAD(path, handler)
	case "ANY": // Use gin.Any() for custom HTTP methods
		a.engine.Any(path, handler)
	default:
		// Support for custom HTTP methods
		a.engine.Handle(httpMethod, path, handler)
	}

	return a
}

func (a *HttpTester) RegisterRoutes(routes []Route) *HttpTester {
	for _, route := range routes {
		a.RegisterRoute(route.Method, route.Path, route.Handler)
	}

	return a
}

func (a *HttpTester) WithHeader(key, value string) *HttpTester {
	a.headers[key] = value

	return a
}

func (a *HttpTester) Get(path string, params map[string]string) *HttpTester {
	a.method = "GET"
	a.path = path

	if len(params) > 0 {
		a.path += "?" + encodeParams(params)
	}

	return a.execute()
}

func (a *HttpTester) Post(path string) *HttpTester {
	a.method = "POST"
	a.path = path

	return a.execute()
}

func (a *HttpTester) Put(path string) *HttpTester {
	a.method = "PUT"
	a.path = path

	return a.execute()
}

func (a *HttpTester) Delete(path string) *HttpTester {
	a.method = "DELETE"
	a.path = path

	return a.execute()
}

func (a *HttpTester) SeeStatus(expectedStatus int) *HttpTester {
	if a.response.Code != expectedStatus {
		a.t.Fatalf("Expected status code %d, but got %d", expectedStatus, a.response.Code)
	}

	return a
}

// DumpResponse logs the response body to the test output.
func (a *HttpTester) DumpResponse() *HttpTester {
	a.t.Logf("Response status: %d", a.response.Code)
	a.t.Logf("Response: %s", a.response.Body.String())

	return a
}

func (a *HttpTester) WithJSONBody(body interface{}) *HttpTester {
	jsonData, err := json.Marshal(body)
	if err != nil {
		a.t.Fatalf("Failed to marshal body: %v", err)
	}

	a.body = bytes.NewReader(jsonData)

	return a
}

func (a *HttpTester) reset() {
	a.response = httptest.NewRecorder()
}

func (a *HttpTester) execute() *HttpTester {
	a.reset()
	req := httptest.NewRequest(a.method, a.path, a.body)

	for key, value := range a.headers {
		req.Header.Set(key, value)
	}

	a.engine.ServeHTTP(a.response, req)
	return a
}

func encodeParams(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}

	return values.Encode()
}
