package api

import (
	"net/http"
	"testing"

	"recotem.org/cli/recotem/pkg/openapi"
)

func TestGetConversionEventsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, openapi.PaginatedConversionEventList{
			Count: intPtr(1),
			Results: &[]openapi.ConversionEvent{
				{
					Id:        intPtr(1),
					AbTest:    1,
					UserId:    "u1",
					EventType: "click",
					Slot:      1,
				},
			},
		})
	})
	defer server.Close()

	result, err := client.GetConversionEvents(nil, nil, nil)
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
	event := (*result.Results)[0]
	if event.UserId != "u1" {
		t.Errorf("expected user_id 'u1', got %s", event.UserId)
	}
	if event.EventType != "click" {
		t.Errorf("expected event_type 'click', got %s", event.EventType)
	}
	if event.AbTest != 1 {
		t.Errorf("expected ab_test 1, got %d", event.AbTest)
	}
}

func TestCreateConversionEventSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, openapi.ConversionEvent{
			Id:        intPtr(1),
			AbTest:    1,
			UserId:    "u1",
			EventType: "click",
			Slot:      1,
		})
	})
	defer server.Close()

	result, err := client.CreateConversionEvent(1, "u1", nil, "click", 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if *result.Id != 1 {
		t.Errorf("expected id 1, got %d", *result.Id)
	}
	if result.UserId != "u1" {
		t.Errorf("expected user_id 'u1', got %s", result.UserId)
	}
	if result.EventType != "click" {
		t.Errorf("expected event_type 'click', got %s", result.EventType)
	}
}

func TestGetConversionEventSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, openapi.ConversionEvent{
			Id:        intPtr(1),
			AbTest:    1,
			UserId:    "u1",
			EventType: "click",
			Slot:      1,
		})
	})
	defer server.Close()

	result, err := client.GetConversionEvent(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if *result.Id != 1 {
		t.Errorf("expected id 1, got %d", *result.Id)
	}
	if result.AbTest != 1 {
		t.Errorf("expected ab_test 1, got %d", result.AbTest)
	}
}

func TestBatchCreateConversionEventsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, []openapi.ConversionEvent{
			{
				Id:        intPtr(1),
				AbTest:    1,
				UserId:    "u1",
				EventType: "click",
				Slot:      1,
			},
		})
	})
	defer server.Close()

	events := []openapi.ConversionEventCreate{
		{
			AbTest:    1,
			UserId:    "u1",
			EventType: "click",
			Slot:      1,
		},
	}

	result, err := client.BatchCreateConversionEvents(events)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if len(result) == 0 {
		t.Fatal("expected non-empty response body")
	}
}
