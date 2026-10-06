package cmd

import (
	"testing"

	"github.com/anthropics/anthropic-cli/internal/mocktest"
)

func TestBetaOrganizationSpendLimitsEffectiveList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:organization:spend-limits:effective", "list",
			"--max-items", "10",
			"--limit", "1",
			"--page", "page",
			"--period", "daily",
			"--user-id", "string",
		)
	})
}
