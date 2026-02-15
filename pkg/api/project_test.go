package api

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateProjectSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":          1,
			"name":        "test-project",
			"user_column": "user_id",
			"item_column": "item_id",
		})
	})
	defer server.Close()

	project, err := client.CreateProject("test-project", "user_id", "item_id", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if project == nil {
		t.Fatal("expected non-nil project")
	}
	if project.Name != "test-project" {
		t.Errorf("expected name 'test-project', got %q", project.Name)
	}
	if project.UserColumn != "user_id" {
		t.Errorf("expected user_column 'user_id', got %q", project.UserColumn)
	}
	if project.ItemColumn != "item_id" {
		t.Errorf("expected item_column 'item_id', got %q", project.ItemColumn)
	}
	if project.Id == nil || *project.Id != 1 {
		t.Errorf("expected id 1, got %v", project.Id)
	}
}

func TestCreateProjectWithTimeColumn(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":          2,
			"name":        "time-project",
			"user_column": "user_id",
			"item_column": "item_id",
			"time_column": "timestamp",
		})
	})
	defer server.Close()

	tc := stringPtr("timestamp")
	project, err := client.CreateProject("time-project", "user_id", "item_id", tc)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if project == nil {
		t.Fatal("expected non-nil project")
	}
	if project.TimeColumn == nil || *project.TimeColumn != "timestamp" {
		t.Errorf("expected time_column 'timestamp', got %v", project.TimeColumn)
	}
}

func TestCreateProjectError(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusBadRequest, map[string]any{
			"detail": "bad request",
		})
	})
	defer server.Close()

	project, err := client.CreateProject("bad-project", "user_id", "item_id", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if project != nil {
		t.Errorf("expected nil project on error, got %v", project)
	}
}

func TestDeleteProjectSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteProject(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDeleteProjectNotFound(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusNotFound, map[string]any{
			"detail": "not found",
		})
	})
	defer server.Close()

	err := client.DeleteProject(999)
	if err == nil {
		t.Fatal("expected error for not found, got nil")
	}
}

func TestGetProjectsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, []map[string]any{
			{
				"id":          1,
				"name":        "project-1",
				"user_column": "user_id",
				"item_column": "item_id",
			},
			{
				"id":          2,
				"name":        "project-2",
				"user_column": "uid",
				"item_column": "iid",
			},
		})
	})
	defer server.Close()

	projects, err := client.GetProjects(nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if projects == nil {
		t.Fatal("expected non-nil projects")
	}
	if len(*projects) != 2 {
		t.Fatalf("expected 2 projects, got %d", len(*projects))
	}
	if (*projects)[0].Name != "project-1" {
		t.Errorf("expected first project name 'project-1', got %q", (*projects)[0].Name)
	}
	if (*projects)[1].Name != "project-2" {
		t.Errorf("expected second project name 'project-2', got %q", (*projects)[1].Name)
	}
}

func TestGetProjectsWithFilter(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, []map[string]any{
			{
				"id":          1,
				"name":        "filtered",
				"user_column": "user_id",
				"item_column": "item_id",
			},
		})
	})
	defer server.Close()

	name := stringPtr("filtered")
	projects, err := client.GetProjects(nil, name)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if projects == nil {
		t.Fatal("expected non-nil projects")
	}
	if len(*projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(*projects))
	}
	if (*projects)[0].Name != "filtered" {
		t.Errorf("expected project name 'filtered', got %q", (*projects)[0].Name)
	}
}

func TestGetProjectSummarySuccess(t *testing.T) {
	summaryData := map[string]any{
		"project_id":  1,
		"total_users": 100,
		"total_items": 500,
	}
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, summaryData)
	})
	defer server.Close()

	body, err := client.GetProjectSummary(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if body == nil {
		t.Fatal("expected non-nil body")
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("failed to unmarshal summary body: %v", err)
	}
	if result["project_id"] != float64(1) {
		t.Errorf("expected project_id 1, got %v", result["project_id"])
	}
	if result["total_users"] != float64(100) {
		t.Errorf("expected total_users 100, got %v", result["total_users"])
	}
}

func TestGetProjectSummaryError(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusNotFound, map[string]any{
			"detail": "not found",
		})
	})
	defer server.Close()

	body, err := client.GetProjectSummary(999)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if body != nil {
		t.Errorf("expected nil body on error, got %v", body)
	}
}
