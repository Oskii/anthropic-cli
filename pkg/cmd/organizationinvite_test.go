package cmd

import (
	"testing"

	"github.com/anthropics/anthropic-cli/internal/mocktest"
)

func TestOrganizationInvitesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:invites", "create",
			"--email", "user@emaildomain.com",
			"--role", "user",
			"--rbac-group-id", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"email: user@emaildomain.com\n" +
			"role: user\n" +
			"rbac_group_ids:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organization:invites", "create",
		)
	})
}

func TestOrganizationInvitesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:invites", "retrieve",
			"--invite-id", "invite_id",
		)
	})
}

func TestOrganizationInvitesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:invites", "list",
			"--max-items", "10",
			"--after-id", "after_id",
			"--before-id", "before_id",
			"--email", "dev@stainless.com",
			"--limit", "1",
			"--role", "string",
			"--status", "accepted",
		)
	})
}

func TestOrganizationInvitesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:invites", "delete",
			"--invite-id", "invite_id",
		)
	})
}
