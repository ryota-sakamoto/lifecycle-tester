package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func TestReadiness(t *testing.T) {
	sm := state.NewStateManager()
	server := setupTestServer(sm)
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/readiness", server.URL))
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

func TestReadiness503(t *testing.T) {
	sm := state.NewStateManager()
	sm.SetState(state.State{
		IsFailedReadiness: true,
	})
	server := setupTestServer(sm)
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/readiness", server.URL))
	assert.NoError(t, err)
	assert.Equal(t, 503, res.StatusCode)
}
