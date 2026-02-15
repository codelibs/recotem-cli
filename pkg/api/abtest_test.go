package api

import (
	"net/http"
	"testing"

	"recotem.org/cli/recotem/pkg/openapi"
)

func TestGetAbTestsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, openapi.PaginatedAbTestList{
			Count: intPtr(1),
			Results: &[]openapi.AbTest{
				{
					Id:      intPtr(1),
					Name:    "test1",
					Project: 1,
					Slots:   []int{1, 2},
					Status:  "draft",
				},
			},
		})
	})
	defer server.Close()

	result, err := client.GetAbTests(nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if *result.Count != 1 {
		t.Errorf("expected count 1, got %d", *result.Count)
	}
	if len(*result.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(*result.Results))
	}
	abTest := (*result.Results)[0]
	if abTest.Name != "test1" {
		t.Errorf("expected name 'test1', got %s", abTest.Name)
	}
	if abTest.Project != 1 {
		t.Errorf("expected project 1, got %d", abTest.Project)
	}
}

func TestCreateAbTestSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, openapi.AbTest{
			Id:      intPtr(1),
			Name:    "new-test",
			Project: 1,
			Slots:   []int{1, 2},
			Status:  "draft",
		})
	})
	defer server.Close()

	result, err := client.CreateAbTest("new-test", 1, []int{1, 2})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Name != "new-test" {
		t.Errorf("expected name 'new-test', got %s", result.Name)
	}
	if *result.Id != 1 {
		t.Errorf("expected id 1, got %d", *result.Id)
	}
}

func TestGetAbTestSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, openapi.AbTest{
			Id:      intPtr(1),
			Name:    "test1",
			Project: 1,
			Slots:   []int{1, 2},
			Status:  "draft",
		})
	})
	defer server.Close()

	result, err := client.GetAbTest(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Name != "test1" {
		t.Errorf("expected name 'test1', got %s", result.Name)
	}
	if *result.Id != 1 {
		t.Errorf("expected id 1, got %d", *result.Id)
	}
}

func TestDeleteAbTestSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteAbTest(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestStartAbTestSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	err := client.StartAbTest(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestStopAbTestSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	err := client.StopAbTest(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGetAbTestResultsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, []openapi.AbTestResult{
			{
				SlotId:         1,
				ConversionRate: 0.5,
				Confidence:     0.95,
				Conversions:    50,
				Impressions:    100,
				SlotName:       "slot-a",
			},
		})
	})
	defer server.Close()

	result, err := client.GetAbTestResults(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if len(*result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(*result))
	}
	r := (*result)[0]
	if r.SlotId != 1 {
		t.Errorf("expected slot_id 1, got %d", r.SlotId)
	}
	if r.ConversionRate != 0.5 {
		t.Errorf("expected conversion_rate 0.5, got %f", r.ConversionRate)
	}
}

func TestPromoteAbTestWinnerSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, openapi.AbTest{
			Id:      intPtr(1),
			Name:    "test1",
			Project: 1,
			Slots:   []int{1},
			Status:  "completed",
		})
	})
	defer server.Close()

	result, err := client.PromoteAbTestWinner(1, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Name != "test1" {
		t.Errorf("expected name 'test1', got %s", result.Name)
	}
	if *result.Id != 1 {
		t.Errorf("expected id 1, got %d", *result.Id)
	}
}
