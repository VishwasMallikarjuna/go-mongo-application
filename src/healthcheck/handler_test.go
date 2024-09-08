package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/VishwasMallikarjuna/go-mongo-application/common/config"
	"github.com/VishwasMallikarjuna/go-mongo-application/common/logwrapper"
	"github.com/VishwasMallikarjuna/go-mongo-application/common/response"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	config := config.Config{
		ConfigPath: "",
	}

	handler := NewHandler(config).(*theHandler)
	assert.True(t, reflect.DeepEqual(config, handler.config))

	// Can't check partitionReaderFromConfig, because it's an anonymous function
	// This asserts that they are the same function by memory address
	assert.Equal(t, reflect.ValueOf(GetCheck), reflect.ValueOf(handler.healthcheck))
}

func TestHealthcheckHandler(t *testing.T) {
	validConfig := config.Config{
		ConfigPath: "",
	}

	tests := []struct {
		name         string
		handler      *theHandler
		expectedCode int
		expectedBody string
	}{
		{
			name: "Good healthcheck",
			handler: &theHandler{
				config: validConfig,
				hriHealthcheck: func(requestId string, healthChecker kafka.HealthChecker) (int, *response.ErrorDetail) {
					return http.StatusOK, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedBody: "",
		},
		{
			name: "Bad healthcheck",
			handler: &theHandler{
				config: config.Config{},
				hriHealthcheck: func(requestId string, healthChecker kafka.HealthChecker) (int, *response.ErrorDetail) {
					return http.StatusServiceUnavailable, response.NewErrorDetail(requestId, "Cosmos not available")
				},
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "{\"errorEventId\":\"" + requestId + "\",\"errorDescription\":\"Cosmos not available\"}\n",
		},
	}

	logwrapper.Initialize("error", os.Stdout)
	e := test.GetTestServer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
			context, recorder := test.PrepareHeadersContextRecorder(request, e)
			context.Response().Header().Add(echo.HeaderXRequestID, requestId)

			if assert.NoError(t, tt.handler.HriHealthcheck(context)) {
				assert.Equal(t, tt.expectedCode, recorder.Code)
				assert.Equal(t, tt.expectedBody, recorder.Body.String())
			}
		})
	}
}
