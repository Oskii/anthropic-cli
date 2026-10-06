package cmd

import (
	"testing"

	"github.com/anthropics/anthropic-cli/internal/mocktest"
)

func TestBetaOrganizationRBACRolesPermissionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:organization:rbac-roles:permissions", "list",
			"--max-items", "10",
			"--rbac-role-id", "rbac_role_id",
			"--limit", "1",
			"--page", "eyJjdXJzb3IiOiAicmJhY19yb2xlXzAxIn0",
		)
	})
}
