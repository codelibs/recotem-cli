package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

// assertFlag checks that a flag exists on the command with the expected shorthand and default value.
func assertFlag(t *testing.T, cmd *cobra.Command, name, shorthand, defValue string) {
	t.Helper()
	flag := cmd.Flags().Lookup(name)
	if flag == nil {
		t.Errorf("expected flag --%s to exist on %q", name, cmd.Name())
		return
	}
	if flag.Shorthand != shorthand {
		t.Errorf("flag --%s: expected shorthand %q, got %q", name, shorthand, flag.Shorthand)
	}
	if flag.DefValue != defValue {
		t.Errorf("flag --%s: expected default %q, got %q", name, defValue, flag.DefValue)
	}
}

// assertRequiredFlag checks that a flag is marked as required.
func assertRequiredFlag(t *testing.T, cmd *cobra.Command, name string) {
	t.Helper()
	flag := cmd.Flags().Lookup(name)
	if flag == nil {
		t.Fatalf("expected flag --%s to exist on %q", name, cmd.Name())
	}
	ann := flag.Annotations
	if ann == nil {
		t.Errorf("expected flag --%s to be required on %q", name, cmd.Name())
		return
	}
	if _, ok := ann[cobra.BashCompOneRequiredFlag]; !ok {
		t.Errorf("expected flag --%s to be required on %q", name, cmd.Name())
	}
}

// assertNotRequiredFlag checks that a flag is NOT marked as required.
func assertNotRequiredFlag(t *testing.T, cmd *cobra.Command, name string) {
	t.Helper()
	flag := cmd.Flags().Lookup(name)
	if flag == nil {
		t.Fatalf("expected flag --%s to exist on %q", name, cmd.Name())
	}
	ann := flag.Annotations
	if ann != nil {
		if _, ok := ann[cobra.BashCompOneRequiredFlag]; ok {
			t.Errorf("expected flag --%s to NOT be required on %q", name, cmd.Name())
		}
	}
}

// --- Ping Command ---

func TestPingCmdStructure(t *testing.T) {
	cmd := newPingCmd()

	if cmd.Use != "ping" {
		t.Errorf("expected Use %q, got %q", "ping", cmd.Use)
	}
	if cmd.Short != "Check server health" {
		t.Errorf("expected Short %q, got %q", "Check server health", cmd.Short)
	}
	if cmd.RunE == nil {
		t.Error("expected RunE to be set")
	}
}

// --- Project Command ---

func TestProjectListCmdFlags(t *testing.T) {
	cmd := newProjectListCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "name", "n", "")
	assertNotRequiredFlag(t, cmd, "id")
	assertNotRequiredFlag(t, cmd, "name")
}

func TestProjectCreateCmdFlags(t *testing.T) {
	cmd := newProjectCreateCmd()

	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "user-column", "u", "")
	assertFlag(t, cmd, "item-column", "i", "")
	assertFlag(t, cmd, "time-column", "t", "")

	assertRequiredFlag(t, cmd, "name")
	assertRequiredFlag(t, cmd, "user-column")
	assertRequiredFlag(t, cmd, "item-column")
	assertNotRequiredFlag(t, cmd, "time-column")
}

func TestProjectDeleteCmdFlags(t *testing.T) {
	cmd := newProjectDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestProjectSummaryCmdFlags(t *testing.T) {
	cmd := newProjectSummaryCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- Trained Model Command ---

func TestTrainedModelListCmdFlags(t *testing.T) {
	cmd := newTrainedModelListCmd()

	assertFlag(t, cmd, "data-loc", "", "")
	assertFlag(t, cmd, "data-loc-project", "", "")
	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")

	assertNotRequiredFlag(t, cmd, "data-loc")
	assertNotRequiredFlag(t, cmd, "id")
	assertNotRequiredFlag(t, cmd, "page")
}

func TestTrainedModelCreateCmdFlags(t *testing.T) {
	cmd := newTrainedModelCreateCmd()

	assertFlag(t, cmd, "configuration", "c", "")
	assertFlag(t, cmd, "data-loc", "", "")
	assertFlag(t, cmd, "file", "f", "")
	assertFlag(t, cmd, "irspack-version", "", "")

	assertRequiredFlag(t, cmd, "configuration")
	assertRequiredFlag(t, cmd, "data-loc")
	assertNotRequiredFlag(t, cmd, "file")
	assertNotRequiredFlag(t, cmd, "irspack-version")
}

func TestTrainedModelDeleteCmdFlags(t *testing.T) {
	cmd := newTrainedModelDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestTrainedModelDownloadCmdFlags(t *testing.T) {
	cmd := newTrainedModelDownloadCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "output", "O", "")
	assertRequiredFlag(t, cmd, "id")
	assertRequiredFlag(t, cmd, "output")
}

func TestTrainedModelRecommendCmdFlags(t *testing.T) {
	cmd := newTrainedModelRecommendCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "user-id", "", "")
	assertFlag(t, cmd, "n-items", "n", "10")
	assertRequiredFlag(t, cmd, "id")
	assertRequiredFlag(t, cmd, "user-id")
	assertNotRequiredFlag(t, cmd, "n-items")
}

func TestTrainedModelSampleRecommendCmdFlags(t *testing.T) {
	cmd := newTrainedModelSampleRecommendCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestTrainedModelRecommendProfileCmdFlags(t *testing.T) {
	cmd := newTrainedModelRecommendProfileCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "n-items", "n", "10")
	assertRequiredFlag(t, cmd, "id")
	assertRequiredFlag(t, cmd, "item-ids")
	assertNotRequiredFlag(t, cmd, "n-items")
}

// --- Task Log Command ---

func TestTaskLogListCmdFlags(t *testing.T) {
	cmd := newTaskLogListCmd()

	assertFlag(t, cmd, "task", "", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")

	assertNotRequiredFlag(t, cmd, "task")
	assertNotRequiredFlag(t, cmd, "page")
	assertNotRequiredFlag(t, cmd, "page-size")
}

// --- API Key Command ---

func TestApiKeyListCmdFlags(t *testing.T) {
	cmd := newApiKeyListCmd()

	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
}

func TestApiKeyCreateCmdFlags(t *testing.T) {
	cmd := newApiKeyCreateCmd()

	assertFlag(t, cmd, "name", "n", "")
	assertRequiredFlag(t, cmd, "name")
}

func TestApiKeyGetCmdFlags(t *testing.T) {
	cmd := newApiKeyGetCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestApiKeyRevokeCmdFlags(t *testing.T) {
	cmd := newApiKeyRevokeCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestApiKeyDeleteCmdFlags(t *testing.T) {
	cmd := newApiKeyDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- Deployment Slot Command ---

func TestDeploymentSlotListCmdFlags(t *testing.T) {
	cmd := newDeploymentSlotListCmd()

	assertFlag(t, cmd, "project", "", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
}

func TestDeploymentSlotCreateCmdFlags(t *testing.T) {
	cmd := newDeploymentSlotCreateCmd()

	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "project", "", "")
	assertFlag(t, cmd, "trained-model", "", "")
	assertRequiredFlag(t, cmd, "name")
	assertRequiredFlag(t, cmd, "project")
	assertNotRequiredFlag(t, cmd, "trained-model")
}

func TestDeploymentSlotGetCmdFlags(t *testing.T) {
	cmd := newDeploymentSlotGetCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestDeploymentSlotUpdateCmdFlags(t *testing.T) {
	cmd := newDeploymentSlotUpdateCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "trained-model", "", "")
	assertRequiredFlag(t, cmd, "id")
	assertNotRequiredFlag(t, cmd, "name")
}

func TestDeploymentSlotDeleteCmdFlags(t *testing.T) {
	cmd := newDeploymentSlotDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- A/B Test Command ---

func TestAbTestListCmdFlags(t *testing.T) {
	cmd := newAbTestListCmd()

	assertFlag(t, cmd, "project", "", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
}

func TestAbTestCreateCmdFlags(t *testing.T) {
	cmd := newAbTestCreateCmd()

	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "project", "", "")
	assertFlag(t, cmd, "slots", "", "")
	assertRequiredFlag(t, cmd, "name")
	assertRequiredFlag(t, cmd, "project")
	assertRequiredFlag(t, cmd, "slots")
}

func TestAbTestGetCmdFlags(t *testing.T) {
	cmd := newAbTestGetCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestAbTestUpdateCmdFlags(t *testing.T) {
	cmd := newAbTestUpdateCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "name", "n", "")
	assertRequiredFlag(t, cmd, "id")
	assertNotRequiredFlag(t, cmd, "name")
}

func TestAbTestDeleteCmdFlags(t *testing.T) {
	cmd := newAbTestDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestAbTestStartCmdFlags(t *testing.T) {
	cmd := newAbTestStartCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestAbTestStopCmdFlags(t *testing.T) {
	cmd := newAbTestStopCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestAbTestResultsCmdFlags(t *testing.T) {
	cmd := newAbTestResultsCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestAbTestPromoteWinnerCmdFlags(t *testing.T) {
	cmd := newAbTestPromoteWinnerCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "slot-id", "", "")
	assertRequiredFlag(t, cmd, "id")
	assertRequiredFlag(t, cmd, "slot-id")
}

// --- Conversion Event Command ---

func TestConversionEventListCmdFlags(t *testing.T) {
	cmd := newConversionEventListCmd()

	assertFlag(t, cmd, "ab-test", "", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
}

func TestConversionEventCreateCmdFlags(t *testing.T) {
	cmd := newConversionEventCreateCmd()

	assertFlag(t, cmd, "ab-test", "", "")
	assertFlag(t, cmd, "slot", "", "")
	assertFlag(t, cmd, "user-id", "", "")
	assertFlag(t, cmd, "item-id", "", "")
	assertFlag(t, cmd, "event-type", "", "")
	assertRequiredFlag(t, cmd, "ab-test")
	assertRequiredFlag(t, cmd, "slot")
	assertRequiredFlag(t, cmd, "user-id")
	assertRequiredFlag(t, cmd, "event-type")
	assertNotRequiredFlag(t, cmd, "item-id")
}

func TestConversionEventBatchCreateCmdFlags(t *testing.T) {
	cmd := newConversionEventBatchCreateCmd()

	assertFlag(t, cmd, "file", "f", "")
	assertRequiredFlag(t, cmd, "file")
}

func TestConversionEventGetCmdFlags(t *testing.T) {
	cmd := newConversionEventGetCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- Retraining Schedule Command ---

func TestRetrainingScheduleListCmdFlags(t *testing.T) {
	cmd := newRetrainingScheduleListCmd()

	assertFlag(t, cmd, "deployment-slot", "", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
}

func TestRetrainingScheduleCreateCmdFlags(t *testing.T) {
	cmd := newRetrainingScheduleCreateCmd()

	assertFlag(t, cmd, "deployment-slot", "", "")
	assertFlag(t, cmd, "cron-expression", "", "")
	assertRequiredFlag(t, cmd, "deployment-slot")
	assertRequiredFlag(t, cmd, "cron-expression")
}

func TestRetrainingScheduleGetCmdFlags(t *testing.T) {
	cmd := newRetrainingScheduleGetCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestRetrainingScheduleUpdateCmdFlags(t *testing.T) {
	cmd := newRetrainingScheduleUpdateCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "cron-expression", "", "")
	assertFlag(t, cmd, "is-active", "", "")
	assertRequiredFlag(t, cmd, "id")
	assertNotRequiredFlag(t, cmd, "cron-expression")
	assertNotRequiredFlag(t, cmd, "is-active")
}

func TestRetrainingScheduleDeleteCmdFlags(t *testing.T) {
	cmd := newRetrainingScheduleDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestRetrainingScheduleTriggerCmdFlags(t *testing.T) {
	cmd := newRetrainingScheduleTriggerCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- Retraining Run Command ---

func TestRetrainingRunListCmdFlags(t *testing.T) {
	cmd := newRetrainingRunListCmd()

	assertFlag(t, cmd, "schedule", "", "")
	assertFlag(t, cmd, "status", "", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
}

func TestRetrainingRunGetCmdFlags(t *testing.T) {
	cmd := newRetrainingRunGetCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- User Command ---

func TestUserListCmdFlags(t *testing.T) {
	cmd := newUserListCmd()

	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
}

func TestUserCreateCmdFlags(t *testing.T) {
	cmd := newUserCreateCmd()

	assertFlag(t, cmd, "username", "", "")
	assertFlag(t, cmd, "email", "", "")
	assertFlag(t, cmd, "password", "", "")
	assertRequiredFlag(t, cmd, "username")
	assertRequiredFlag(t, cmd, "email")
	assertRequiredFlag(t, cmd, "password")
}

func TestUserGetCmdFlags(t *testing.T) {
	cmd := newUserGetCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestUserUpdateCmdFlags(t *testing.T) {
	cmd := newUserUpdateCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "email", "", "")
	assertFlag(t, cmd, "is-active", "", "")
	assertRequiredFlag(t, cmd, "id")
	assertNotRequiredFlag(t, cmd, "email")
	assertNotRequiredFlag(t, cmd, "is-active")
}

func TestUserDeactivateCmdFlags(t *testing.T) {
	cmd := newUserDeactivateCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestUserActivateCmdFlags(t *testing.T) {
	cmd := newUserActivateCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestUserResetPasswordCmdFlags(t *testing.T) {
	cmd := newUserResetPasswordCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "new-password", "", "")
	assertRequiredFlag(t, cmd, "id")
	assertRequiredFlag(t, cmd, "new-password")
}

// --- Login Command ---

func TestLoginCmdFlags(t *testing.T) {
	cmd := newLoginCmd()

	if cmd.Use != "login" {
		t.Errorf("expected Use %q, got %q", "login", cmd.Use)
	}
	assertFlag(t, cmd, "username", "u", "")
	assertFlag(t, cmd, "password", "p", "")
	assertNotRequiredFlag(t, cmd, "username")
	assertNotRequiredFlag(t, cmd, "password")
}

// --- Training Data Command ---

func TestTrainingDataListCmdFlags(t *testing.T) {
	cmd := newTrainingDataListCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
	assertFlag(t, cmd, "project", "", "")
}

func TestTrainingDataUploadCmdFlags(t *testing.T) {
	cmd := newTrainingDataUploadCmd()

	assertFlag(t, cmd, "project", "p", "")
	assertFlag(t, cmd, "file", "f", "")
	assertRequiredFlag(t, cmd, "project")
	assertRequiredFlag(t, cmd, "file")
}

func TestTrainingDataDeleteCmdFlags(t *testing.T) {
	cmd := newTrainingDataDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestTrainingDataDownloadCmdFlags(t *testing.T) {
	cmd := newTrainingDataDownloadCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "output", "O", "")
	assertRequiredFlag(t, cmd, "id")
	assertRequiredFlag(t, cmd, "output")
}

func TestTrainingDataPreviewCmdFlags(t *testing.T) {
	cmd := newTrainingDataPreviewCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- Item Meta Data Command ---

func TestItemMetaDataListCmdFlags(t *testing.T) {
	cmd := newItemMetaDataListCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
	assertFlag(t, cmd, "project", "", "")
}

func TestItemMetaDataUploadCmdFlags(t *testing.T) {
	cmd := newItemMetaDataUploadCmd()

	assertFlag(t, cmd, "project", "p", "")
	assertFlag(t, cmd, "file", "f", "")
	assertRequiredFlag(t, cmd, "project")
	assertRequiredFlag(t, cmd, "file")
}

func TestItemMetaDataDeleteCmdFlags(t *testing.T) {
	cmd := newItemMetaDataDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestItemMetaDataDownloadCmdFlags(t *testing.T) {
	cmd := newItemMetaDataDownloadCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "output", "O", "")
	assertRequiredFlag(t, cmd, "id")
	assertRequiredFlag(t, cmd, "output")
}

// --- Evaluation Config Command ---

func TestEvaluationConfigListCmdFlags(t *testing.T) {
	cmd := newEvaluationConfigListCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "unnamed", "u", "")
}

func TestEvaluationConfigCreateCmdFlags(t *testing.T) {
	cmd := newEvaluationConfigCreateCmd()

	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "cutoff", "c", "")
	assertFlag(t, cmd, "target-metric", "", "")
}

func TestEvaluationConfigDeleteCmdFlags(t *testing.T) {
	cmd := newEvaluationConfigDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestEvaluationConfigUpdateCmdFlags(t *testing.T) {
	cmd := newEvaluationConfigUpdateCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "cutoff", "c", "")
	assertFlag(t, cmd, "target-metric", "", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- Model Configuration Command ---

func TestModelConfigurationListCmdFlags(t *testing.T) {
	cmd := newModelConfigurationListCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
	assertFlag(t, cmd, "project", "", "")
}

func TestModelConfigurationCreateCmdFlags(t *testing.T) {
	cmd := newModelConfigurationCreateCmd()

	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "project", "p", "")
	assertFlag(t, cmd, "recommender-class-name", "", "")
	assertFlag(t, cmd, "parameters-json", "", "")
	assertRequiredFlag(t, cmd, "project")
	assertRequiredFlag(t, cmd, "recommender-class-name")
	assertRequiredFlag(t, cmd, "parameters-json")
	assertNotRequiredFlag(t, cmd, "name")
}

func TestModelConfigurationDeleteCmdFlags(t *testing.T) {
	cmd := newModelConfigurationDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestModelConfigurationUpdateCmdFlags(t *testing.T) {
	cmd := newModelConfigurationUpdateCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "recommender-class-name", "", "")
	assertFlag(t, cmd, "parameters-json", "", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- Split Config Command ---

func TestSplitConfigListCmdFlags(t *testing.T) {
	cmd := newSplitConfigListCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "unnamed", "u", "")
}

func TestSplitConfigCreateCmdFlags(t *testing.T) {
	cmd := newSplitConfigCreateCmd()

	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "scheme", "s", "")
	assertFlag(t, cmd, "heldout-ratio", "", "")
	assertFlag(t, cmd, "n-heldout", "", "")
	assertFlag(t, cmd, "test-user-ratio", "", "")
	assertFlag(t, cmd, "n-test-users", "", "")
	assertFlag(t, cmd, "random-seed", "", "")
}

func TestSplitConfigDeleteCmdFlags(t *testing.T) {
	cmd := newSplitConfigDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

func TestSplitConfigUpdateCmdFlags(t *testing.T) {
	cmd := newSplitConfigUpdateCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "name", "n", "")
	assertFlag(t, cmd, "scheme", "s", "")
	assertFlag(t, cmd, "heldout-ratio", "", "")
	assertFlag(t, cmd, "n-heldout", "", "")
	assertFlag(t, cmd, "test-user-ratio", "", "")
	assertFlag(t, cmd, "n-test-users", "", "")
	assertFlag(t, cmd, "random-seed", "", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- Parameter Tuning Job Command ---

func TestParameterTuningJobListCmdFlags(t *testing.T) {
	cmd := newParameterTuningJobListCmd()

	assertFlag(t, cmd, "data", "d", "")
	assertFlag(t, cmd, "data-project", "", "")
	assertFlag(t, cmd, "id", "i", "")
	assertFlag(t, cmd, "page", "p", "")
	assertFlag(t, cmd, "page-size", "", "")
}

func TestParameterTuningJobCreateCmdFlags(t *testing.T) {
	cmd := newParameterTuningJobCreateCmd()

	assertFlag(t, cmd, "data", "d", "")
	assertFlag(t, cmd, "split", "s", "")
	assertFlag(t, cmd, "evaluation", "e", "")
	assertRequiredFlag(t, cmd, "data")
	assertRequiredFlag(t, cmd, "split")
	assertRequiredFlag(t, cmd, "evaluation")

	// Optional flags
	assertFlag(t, cmd, "n-tasks-parallel", "", "")
	assertFlag(t, cmd, "n-trials", "", "")
	assertFlag(t, cmd, "memory-budget", "", "")
	assertFlag(t, cmd, "timeout-overall", "", "")
	assertFlag(t, cmd, "timeout-singlestep", "", "")
	assertFlag(t, cmd, "random-seed", "", "")
	assertFlag(t, cmd, "tried-algorithm-json", "", "")
	assertFlag(t, cmd, "irspack-version", "", "")
	assertFlag(t, cmd, "train-after-tuning", "", "")
	assertFlag(t, cmd, "best-score", "", "")
	assertFlag(t, cmd, "tuned-model", "", "")
	assertFlag(t, cmd, "best-config", "", "")
}

func TestParameterTuningJobDeleteCmdFlags(t *testing.T) {
	cmd := newParameterTuningJobDeleteCmd()

	assertFlag(t, cmd, "id", "i", "")
	assertRequiredFlag(t, cmd, "id")
}

// --- Completion Command ---

func TestCompletionCmdStructure(t *testing.T) {
	cmd := newCompletionCmd()

	if cmd.Use != "completion [bash|zsh|fish|powershell]" {
		t.Errorf("expected Use %q, got %q", "completion [bash|zsh|fish|powershell]", cmd.Use)
	}

	if len(cmd.ValidArgs) != 4 {
		t.Errorf("expected 4 valid args, got %d", len(cmd.ValidArgs))
	}

	validArgs := map[string]bool{
		"bash":       false,
		"zsh":        false,
		"fish":       false,
		"powershell": false,
	}
	for _, arg := range cmd.ValidArgs {
		if _, ok := validArgs[arg]; !ok {
			t.Errorf("unexpected valid arg: %s", arg)
		}
		validArgs[arg] = true
	}
	for arg, found := range validArgs {
		if !found {
			t.Errorf("missing valid arg: %s", arg)
		}
	}
}
