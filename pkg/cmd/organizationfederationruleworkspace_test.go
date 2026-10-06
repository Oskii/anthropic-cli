package cmd

import (
	"testing"

	"github.com/anthropics/anthropic-cli/internal/mocktest"
)

func TestOrganizationFederationRulesWorkspacesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:federation:rules:workspaces", "list",
			"--max-items", "10",
			"--federation-rule-id", "federation_rule_id",
			"--limit", "1",
			"--page", "page",
		)
	})
}

func TestOrganizationFederationRulesWorkspacesAdd(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:federation:rules:workspaces", "add",
			"--federation-rule-id", "federation_rule_id",
			"--workspace-id", "workspace_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("workspace_id: workspace_id")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organization:federation:rules:workspaces", "add",
			"--federation-rule-id", "federation_rule_id",
		)
	})
}

func TestOrganizationFederationRulesWorkspacesRemove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organization:federation:rules:workspaces", "remove",
			"--federation-rule-id", "federation_rule_id",
			"--workspace-id", "workspace_id",
		)
	})
}
