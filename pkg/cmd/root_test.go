package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

// findSubcommand searches for a subcommand by name within a given command.
func findSubcommand(cmd *cobra.Command, name string) *cobra.Command {
	for _, c := range cmd.Commands() {
		if c.Name() == name {
			return c
		}
	}
	return nil
}

// assertSubcommands verifies that all expected subcommand names are present
// under the given parent command.
func assertSubcommands(t *testing.T, cmd *cobra.Command, expected []string) {
	t.Helper()
	names := make(map[string]bool)
	for _, c := range cmd.Commands() {
		names[c.Name()] = true
	}
	for _, name := range expected {
		if !names[name] {
			t.Errorf("expected subcommand %q not found in %q", name, cmd.Name())
		}
	}
}

// assertAlias verifies that the given command has the expected alias.
func assertAlias(t *testing.T, cmd *cobra.Command, expectedAlias string) {
	t.Helper()
	for _, a := range cmd.Aliases {
		if a == expectedAlias {
			return
		}
	}
	t.Errorf("expected alias %q not found in command %q (aliases: %v)", expectedAlias, cmd.Name(), cmd.Aliases)
}

func TestNewRootCmd(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")

	if cmd.Use != "recotem" {
		t.Errorf("expected Use to be %q, got %q", "recotem", cmd.Use)
	}

	if cmd.Short != "CLI for recotem recommendation system" {
		t.Errorf("expected Short to be %q, got %q", "CLI for recotem recommendation system", cmd.Short)
	}

	if !cmd.SilenceUsage {
		t.Error("expected SilenceUsage to be true")
	}
}

func TestRootCmdPersistentFlags(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")

	outputFlag := cmd.PersistentFlags().Lookup("output")
	if outputFlag == nil {
		t.Fatal("expected persistent flag --output to exist")
	}
	if outputFlag.DefValue != "text" {
		t.Errorf("expected --output default to be %q, got %q", "text", outputFlag.DefValue)
	}
	if outputFlag.Shorthand != "o" {
		t.Errorf("expected --output shorthand to be %q, got %q", "o", outputFlag.Shorthand)
	}

	apiKeyFlag := cmd.PersistentFlags().Lookup("api-key")
	if apiKeyFlag == nil {
		t.Fatal("expected persistent flag --api-key to exist")
	}
	if apiKeyFlag.DefValue != "" {
		t.Errorf("expected --api-key default to be empty, got %q", apiKeyFlag.DefValue)
	}
}

func TestRootCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")

	expectedSubcommands := []string{
		"login",
		"logout",
		"version",
		"completion",
		"ping",
		"project",
		"training-data",
		"item-meta-data",
		"trained-model",
		"model-configuration",
		"evaluation-config",
		"split-config",
		"parameter-tuning-job",
		"api-key",
		"deployment-slot",
		"ab-test",
		"conversion-event",
		"retraining-schedule",
		"retraining-run",
		"task-log",
		"user",
	}

	assertSubcommands(t, cmd, expectedSubcommands)

	// Verify the total count of registered subcommands.
	// Cobra may add a built-in "help" command, so we check that at least
	// all 21 explicitly registered commands are present.
	registered := cmd.Commands()
	if len(registered) < len(expectedSubcommands) {
		t.Errorf("expected at least %d subcommands, got %d", len(expectedSubcommands), len(registered))
	}
}

func TestProjectCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	projectCmd := findSubcommand(cmd, "project")
	if projectCmd == nil {
		t.Fatal("expected project command to exist")
	}

	assertAlias(t, projectCmd, "p")

	expected := []string{"list", "create", "delete", "summary"}
	assertSubcommands(t, projectCmd, expected)
}

func TestTrainingDataCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	tdCmd := findSubcommand(cmd, "training-data")
	if tdCmd == nil {
		t.Fatal("expected training-data command to exist")
	}

	assertAlias(t, tdCmd, "td")

	expected := []string{"list", "upload", "delete", "download", "preview"}
	assertSubcommands(t, tdCmd, expected)
}

func TestItemMetaDataCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	imdCmd := findSubcommand(cmd, "item-meta-data")
	if imdCmd == nil {
		t.Fatal("expected item-meta-data command to exist")
	}

	assertAlias(t, imdCmd, "imd")

	expected := []string{"list", "upload", "delete", "download"}
	assertSubcommands(t, imdCmd, expected)
}

func TestTrainedModelCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	tmCmd := findSubcommand(cmd, "trained-model")
	if tmCmd == nil {
		t.Fatal("expected trained-model command to exist")
	}

	assertAlias(t, tmCmd, "tm")

	expected := []string{"list", "create", "delete", "download", "recommend", "sample-recommend", "recommend-profile"}
	assertSubcommands(t, tmCmd, expected)
}

func TestModelConfigurationCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	mcCmd := findSubcommand(cmd, "model-configuration")
	if mcCmd == nil {
		t.Fatal("expected model-configuration command to exist")
	}

	assertAlias(t, mcCmd, "mc")

	expected := []string{"list", "create", "delete", "update"}
	assertSubcommands(t, mcCmd, expected)
}

func TestEvaluationConfigCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	ecCmd := findSubcommand(cmd, "evaluation-config")
	if ecCmd == nil {
		t.Fatal("expected evaluation-config command to exist")
	}

	assertAlias(t, ecCmd, "ec")

	expected := []string{"list", "create", "delete", "update"}
	assertSubcommands(t, ecCmd, expected)
}

func TestSplitConfigCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	scCmd := findSubcommand(cmd, "split-config")
	if scCmd == nil {
		t.Fatal("expected split-config command to exist")
	}

	assertAlias(t, scCmd, "sc")

	expected := []string{"list", "create", "delete", "update"}
	assertSubcommands(t, scCmd, expected)
}

func TestParameterTuningJobCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	ptjCmd := findSubcommand(cmd, "parameter-tuning-job")
	if ptjCmd == nil {
		t.Fatal("expected parameter-tuning-job command to exist")
	}

	assertAlias(t, ptjCmd, "ptj")

	expected := []string{"list", "create", "delete"}
	assertSubcommands(t, ptjCmd, expected)
}

func TestApiKeyCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	akCmd := findSubcommand(cmd, "api-key")
	if akCmd == nil {
		t.Fatal("expected api-key command to exist")
	}

	assertAlias(t, akCmd, "ak")

	expected := []string{"list", "create", "get", "revoke", "delete"}
	assertSubcommands(t, akCmd, expected)
}

func TestDeploymentSlotCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	dsCmd := findSubcommand(cmd, "deployment-slot")
	if dsCmd == nil {
		t.Fatal("expected deployment-slot command to exist")
	}

	assertAlias(t, dsCmd, "ds")

	expected := []string{"list", "create", "get", "update", "delete"}
	assertSubcommands(t, dsCmd, expected)
}

func TestAbTestCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	abCmd := findSubcommand(cmd, "ab-test")
	if abCmd == nil {
		t.Fatal("expected ab-test command to exist")
	}

	assertAlias(t, abCmd, "ab")

	expected := []string{"list", "create", "get", "update", "delete", "start", "stop", "results", "promote-winner"}
	assertSubcommands(t, abCmd, expected)
}

func TestConversionEventCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	ceCmd := findSubcommand(cmd, "conversion-event")
	if ceCmd == nil {
		t.Fatal("expected conversion-event command to exist")
	}

	assertAlias(t, ceCmd, "ce")

	expected := []string{"list", "create", "batch-create", "get"}
	assertSubcommands(t, ceCmd, expected)
}

func TestRetrainingScheduleCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	rsCmd := findSubcommand(cmd, "retraining-schedule")
	if rsCmd == nil {
		t.Fatal("expected retraining-schedule command to exist")
	}

	assertAlias(t, rsCmd, "rs")

	expected := []string{"list", "create", "get", "update", "delete", "trigger"}
	assertSubcommands(t, rsCmd, expected)
}

func TestRetrainingRunCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	rrCmd := findSubcommand(cmd, "retraining-run")
	if rrCmd == nil {
		t.Fatal("expected retraining-run command to exist")
	}

	assertAlias(t, rrCmd, "rr")

	expected := []string{"list", "get"}
	assertSubcommands(t, rrCmd, expected)
}

func TestTaskLogCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	tlCmd := findSubcommand(cmd, "task-log")
	if tlCmd == nil {
		t.Fatal("expected task-log command to exist")
	}

	assertAlias(t, tlCmd, "tl")

	expected := []string{"list"}
	assertSubcommands(t, tlCmd, expected)
}

func TestUserCmdSubcommands(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	userCmd := findSubcommand(cmd, "user")
	if userCmd == nil {
		t.Fatal("expected user command to exist")
	}

	assertAlias(t, userCmd, "u")

	expected := []string{"list", "create", "get", "update", "deactivate", "activate", "reset-password"}
	assertSubcommands(t, userCmd, expected)
}

func TestGetOutputFormatDefault(t *testing.T) {
	// Save the original value and restore it after the test.
	original := outputFormat
	defer func() { outputFormat = original }()

	outputFormat = ""
	result := getOutputFormat()
	if result != "text" {
		t.Errorf("expected getOutputFormat() to return %q when outputFormat is empty, got %q", "text", result)
	}

	outputFormat = "json"
	result = getOutputFormat()
	if result != "json" {
		t.Errorf("expected getOutputFormat() to return %q, got %q", "json", result)
	}

	outputFormat = "yaml"
	result = getOutputFormat()
	if result != "yaml" {
		t.Errorf("expected getOutputFormat() to return %q, got %q", "yaml", result)
	}
}

func TestVersionCmd(t *testing.T) {
	cmd := NewRootCmd("1.0.0", "abc123", "2024-01-01")
	versionCmd := findSubcommand(cmd, "version")
	if versionCmd == nil {
		t.Fatal("expected version command to exist")
	}

	if versionCmd.Use != "version" {
		t.Errorf("expected version command Use to be %q, got %q", "version", versionCmd.Use)
	}

	if versionCmd.Short != "Print version information" {
		t.Errorf("expected version command Short to be %q, got %q", "Print version information", versionCmd.Short)
	}
}
