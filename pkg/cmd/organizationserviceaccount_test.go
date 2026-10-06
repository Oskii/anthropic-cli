package cmd

import (
	"testing"

	"github.com/anthropics/anthropic-cli/internal/mocktest"
)

func TestOrganizationServiceAccountsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:service-accounts", "create",
			"--name", "ci-deploy-bot",
			"--description", "description",
			"--organization-role", "admin",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: ci-deploy-bot\n" +
			"description: description\n" +
			"organization_role: admin\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organization:service-accounts", "create",
		)
	})
}

func TestOrganizationServiceAccountsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:service-accounts", "retrieve",
			"--service-account-id", "service_account_id",
		)
	})
}

func TestOrganizationServiceAccountsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:service-accounts", "update",
			"--service-account-id", "service_account_id",
			"--description", "description",
			"--organization-role", "admin",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"organization_role: admin\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organization:service-accounts", "update",
			"--service-account-id", "service_account_id",
		)
	})
}

func TestOrganizationServiceAccountsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:service-accounts", "list",
			"--max-items", "10",
			"--include-archived=true",
			"--limit", "1",
			"--page", "page",
		)
	})
}

func TestOrganizationServiceAccountsArchive(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:service-accounts", "archive",
			"--service-account-id", "service_account_id",
		)
	})
}
