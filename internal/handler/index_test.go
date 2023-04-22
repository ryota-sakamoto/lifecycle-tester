package handler_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func TestGetIndex(t *testing.T) {
	sm := state.NewStateManager()
	server := setupTestServer(sm)
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, 200, res.StatusCode)

	type BodyType struct {
		Hostname string
		Request  map[string]interface{}
		State    state.State
	}

	body := BodyType{}
	err = json.NewDecoder(res.Body).Decode(&body)
	assert.NoError(t, err)

	hostname, err := os.Hostname()
	assert.NoError(t, err)

	expected := BodyType{
		Hostname: hostname,
		Request: map[string]interface{}{
			"header": map[string]interface{}{
				"Accept-Encoding": []interface{}{"gzip"},
				"User-Agent":      []interface{}{"Go-http-client/1.1"},
			},
			"host":        res.Request.Host,
			"method":      "GET",
			"remote_addr": body.Request["remote_addr"],
			"request_uri": "/",
		},
		State: state.State{
			IsFailedLiveness:     false,
			IsFailedReadiness:    false,
			ShutdownDelaySeconds: 0,
		},
	}
	assert.Equal(t, expected, body)
}
