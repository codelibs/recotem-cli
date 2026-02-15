package api

import (
	"net/http"
	"testing"
)

func TestGetDeploymentSlotsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"count": 1,
			"results": []map[string]any{
				{"id": 1, "name": "slot1", "project": 1, "is_active": true},
			},
		})
	})
	defer server.Close()

	result, err := client.GetDeploymentSlots(nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Count == nil || *result.Count != 1 {
		t.Fatalf("expected count 1, got %v", result.Count)
	}
	if result.Results == nil || len(*result.Results) != 1 {
		t.Fatal("expected 1 result")
	}
	slot := (*result.Results)[0]
	if slot.Name != "slot1" {
		t.Errorf("expected name 'slot1', got %s", slot.Name)
	}
	if slot.Project != 1 {
		t.Errorf("expected project 1, got %d", slot.Project)
	}
	if slot.IsActive != true {
		t.Errorf("expected is_active true, got %v", slot.IsActive)
	}
}

func TestCreateDeploymentSlotSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":        1,
			"name":      "slot1",
			"project":   1,
			"is_active": true,
		})
	})
	defer server.Close()

	result, err := client.CreateDeploymentSlot("slot1", 1, nil, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Name != "slot1" {
		t.Errorf("expected name 'slot1', got %s", result.Name)
	}
	if result.Project != 1 {
		t.Errorf("expected project 1, got %d", result.Project)
	}
	if result.IsActive != true {
		t.Errorf("expected is_active true, got %v", result.IsActive)
	}
}

func TestGetDeploymentSlotSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":        1,
			"name":      "slot1",
			"project":   1,
			"is_active": true,
		})
	})
	defer server.Close()

	result, err := client.GetDeploymentSlot(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Name != "slot1" {
		t.Errorf("expected name 'slot1', got %s", result.Name)
	}
}

func TestUpdateDeploymentSlotSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":        1,
			"name":      "updated-slot",
			"project":   1,
			"is_active": true,
		})
	})
	defer server.Close()

	result, err := client.UpdateDeploymentSlot(1, stringPtr("updated-slot"), nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Name != "updated-slot" {
		t.Errorf("expected name 'updated-slot', got %s", result.Name)
	}
}

func TestDeleteDeploymentSlotSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteDeploymentSlot(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
