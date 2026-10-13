package cmd

import (
	"testing"

	"github.com/anthropics/anthropic-cli/internal/mocktest"
)

func TestBetaOrganizationAnalyticsSummariesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:organization:analytics:summaries", "list",
			"--max-items", "10",
			"--starting-date", "'2019-12-27'",
			"--ending-date", "'2019-12-27'",
			"--filter", "string",
			"--limit", "1",
			"--page", "page",
		)
	})
}
