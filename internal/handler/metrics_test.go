package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/prometheus/common/expfmt"
	"github.com/stretchr/testify/assert"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func TestMetrics(t *testing.T) {
	sm := state.NewStateManager()
	server := setupTestServer(sm)
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/metrics", server.URL))
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	defer res.Body.Close()

	parser := expfmt.TextParser{}
	mf, err := parser.TextToMetricFamilies(res.Body)
	assert.NoError(t, err)
	assert.NotNil(t, mf["promhttp_metric_handler_requests_total"])
}
