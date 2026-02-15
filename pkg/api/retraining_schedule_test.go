package api

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetRetrainingSchedulesSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/v1/retraining-schedule") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"count": 1,
			"next":  nil,
			"previous": nil,
			"results": []map[string]any{
				{
					"id":              1,
					"deployment_slot": 10,
					"cron_expression": "0 0 * * *",
					"is_active":       true,
					"last_run_at":     nil,
					"next_run_at":     nil,
				},
			},
		})
	})
	defer server.Close()

	result, err := client.GetRetrainingSchedules(nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Count == nil || *result.Count != 1 {
		t.Errorf("expected count 1, got %v", result.Count)
	}
	if result.Results == nil || len(*result.Results) != 1 {
		t.Fatal("expected 1 result")
	}
	schedule := (*result.Results)[0]
	if schedule.Id == nil || *schedule.Id != 1 {
		t.Errorf("expected id 1, got %v", schedule.Id)
	}
	if schedule.DeploymentSlot != 10 {
		t.Errorf("expected deployment_slot 10, got %d", schedule.DeploymentSlot)
	}
	if schedule.CronExpression != "0 0 * * *" {
		t.Errorf("expected cron_expression '0 0 * * *', got %s", schedule.CronExpression)
	}
}

func TestCreateRetrainingScheduleSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/v1/retraining-schedule") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":              1,
			"deployment_slot": 10,
			"cron_expression": "0 0 * * *",
			"is_active":       true,
			"last_run_at":     nil,
			"next_run_at":     nil,
		})
	})
	defer server.Close()

	result, err := client.CreateRetrainingSchedule(10, "0 0 * * *", true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.DeploymentSlot != 10 {
		t.Errorf("expected deployment_slot 10, got %d", result.DeploymentSlot)
	}
	if result.CronExpression != "0 0 * * *" {
		t.Errorf("expected cron_expression '0 0 * * *', got %s", result.CronExpression)
	}
	if !result.IsActive {
		t.Error("expected is_active true")
	}
}

func TestGetRetrainingScheduleSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/v1/retraining-schedule") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":              1,
			"deployment_slot": 10,
			"cron_expression": "0 0 * * *",
			"is_active":       true,
			"last_run_at":     nil,
			"next_run_at":     nil,
		})
	})
	defer server.Close()

	result, err := client.GetRetrainingSchedule(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.DeploymentSlot != 10 {
		t.Errorf("expected deployment_slot 10, got %d", result.DeploymentSlot)
	}
}

func TestUpdateRetrainingScheduleSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/v1/retraining-schedule") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPatch {
			t.Errorf("expected PATCH, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":              1,
			"deployment_slot": 10,
			"cron_expression": "0 12 * * *",
			"is_active":       false,
			"last_run_at":     nil,
			"next_run_at":     nil,
		})
	})
	defer server.Close()

	cronExpr := "0 12 * * *"
	isActive := false
	result, err := client.UpdateRetrainingSchedule(1, &cronExpr, &isActive)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.CronExpression != "0 12 * * *" {
		t.Errorf("expected cron_expression '0 12 * * *', got %s", result.CronExpression)
	}
	if result.IsActive {
		t.Error("expected is_active false")
	}
}

func TestDeleteRetrainingScheduleSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/v1/retraining-schedule") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteRetrainingSchedule(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestTriggerRetrainingScheduleSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/api/v1/retraining-schedule") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if !strings.Contains(r.URL.Path, "trigger") {
			t.Errorf("expected path to contain 'trigger', got %s", r.URL.Path)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"status": "triggered",
		})
	})
	defer server.Close()

	err := client.TriggerRetrainingSchedule(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
