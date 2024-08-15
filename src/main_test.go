package main

import (
	"errors"
	"testing"

	"github.com/VishwasMallikarjuna/go-mongo-application/common/test"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestConfigureMgmtServerErrors(t *testing.T) {
	configPath := test.FindConfigPath(t)
	e := echo.New()

	tests := []struct {
		name               string
		args               []string
		expectedError      error
		expectedReturnCode int
	}{
		{
			name:               "Bad Config",
			expectedReturnCode: 1,
			args:               []string{"--validation=notABool"},
			expectedError:      errors.New("error parsing commandline args: invalid boolean value \"notABool\" for -validation: parse error"),
		},
		// {
		// 	name:               "Bad Log Level",
		// 	expectedReturnCode: 3,
		// 	args:               []string{"--log-level=notALevel"},
		// 	expectedError:      errors.New("ERROR: Could NOT initialize Logger: error parsing log Level - not a valid logrus Level: \"notALevel\""),
		// },
		{
			name:               "Bad New Relic License",
			expectedReturnCode: 1,
			args:               []string{"--new-relic-license-key=notLongEnough"},
			expectedError:      errors.New("ERROR CONFIGURING NEW RELIC: license length is not 40"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rc, startFunc, err := configureMgmtServer(e, append([]string{"--config-path=" + configPath}, tc.args...))
			assert.Nil(t, startFunc)
			if tc.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err, tc.expectedError)
			}

			assert.Equal(t, tc.expectedReturnCode, rc)
		})
	}
}
