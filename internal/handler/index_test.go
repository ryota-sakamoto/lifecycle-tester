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
	type BodyType struct {
		Hostname string
		Request  map[string]interface{}
		State    state.State
	}

	hostname, err := os.Hostname()
	assert.NoError(t, err)

	tt := []struct {
		name     string
		state    *state.State
		expected func(*http.Response, BodyType) BodyType
	}{
		{
			name:  "normal",
			state: nil,
			expected: func(res *http.Response, body BodyType) BodyType {
				return BodyType{
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
						IsFailedReadiness:    false,
						IsFailedLiveness:     false,
						ShutdownDelaySeconds: 0,
					},
				}
			},
		},
		{
			name: "change state",
			state: &state.State{
				IsFailedReadiness:    true,
				IsFailedLiveness:     false,
				ShutdownDelaySeconds: 10,
			},
			expected: func(res *http.Response, body BodyType) BodyType {
				return BodyType{
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
						IsFailedReadiness:    true,
						IsFailedLiveness:     false,
						ShutdownDelaySeconds: 10,
					},
				}
			},
		},
	}

	for _, data := range tt {
		t.Run(data.name, func(t *testing.T) {
			sm := state.NewStateManager()
			if data.state != nil {
				sm.SetState(*data.state)
			}

			server := setupTestServer(sm)
			defer server.Close()

			req, err := http.NewRequest(http.MethodGet, server.URL, nil)
			assert.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			defer res.Body.Close()

			assert.Equal(t, 200, res.StatusCode)

			body := BodyType{}
			err = json.NewDecoder(res.Body).Decode(&body)
			assert.NoError(t, err)

			assert.Equal(t, data.expected(res, body), body)
		})
	}
}
