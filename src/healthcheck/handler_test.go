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