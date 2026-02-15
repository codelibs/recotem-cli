package api

import (
	"net/http"
	"testing"
)

func TestGetUsersSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"count":    1,
			"next":     nil,
			"previous": nil,
			"results": []map[string]any{
				{
					"id":        1,
					"username":  "admin",
					"email":     "admin@test.com",
					"is_active": true,
					"is_staff":  true,
				},
			},
		})
	})
	defer server.Close()

	users, err := client.GetUsers(nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if users == nil {
		t.Fatal("expected non-nil paginated user list")
	}
	if users.Count == nil || *users.Count != 1 {
		t.Errorf("expected count 1, got %v", users.Count)
	}
	if users.Results == nil || len(*users.Results) != 1 {
		t.Fatal("expected 1 user in results")
	}
	user := (*users.Results)[0]
	if user.Username != "admin" {
		t.Errorf("expected username 'admin', got %q", user.Username)
	}
	if !user.IsActive {
		t.Error("expected is_active to be true")
	}
}

func TestGetUsersWithPagination(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"count":    2,
			"next":     "http://example.com/api/v1/user/?page=2",
			"previous": nil,
			"results": []map[string]any{
				{
					"id":        1,
					"username":  "user1",
					"email":     "user1@test.com",
					"is_active": true,
					"is_staff":  false,
				},
			},
		})
	})
	defer server.Close()

	page := intPtr(1)
	pageSize := intPtr(1)
	users, err := client.GetUsers(page, pageSize)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if users == nil {
		t.Fatal("expected non-nil paginated user list")
	}
	if users.Count == nil || *users.Count != 2 {
		t.Errorf("expected count 2, got %v", users.Count)
	}
	if users.Next == nil {
		t.Error("expected non-nil next page URL")
	}
}

func TestCreateUserSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":        2,
			"username":  "newuser",
			"email":     "newuser@test.com",
			"is_active": true,
			"is_staff":  false,
		})
	})
	defer server.Close()

	user, err := client.CreateUser("newuser", "password123", "newuser@test.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user == nil {
		t.Fatal("expected non-nil user")
	}
	if user.Username != "newuser" {
		t.Errorf("expected username 'newuser', got %q", user.Username)
	}
	if user.Id == nil || *user.Id != 2 {
		t.Errorf("expected id 2, got %v", user.Id)
	}
	if string(user.Email) != "newuser@test.com" {
		t.Errorf("expected email 'newuser@test.com', got %q", user.Email)
	}
}

func TestCreateUserError(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusBadRequest, map[string]any{
			"username": []string{"A user with that username already exists."},
		})
	})
	defer server.Close()

	user, err := client.CreateUser("existing", "password123", "existing@test.com")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if user != nil {
		t.Errorf("expected nil user on error, got %v", user)
	}
}

func TestGetUserSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":        1,
			"username":  "admin",
			"email":     "admin@test.com",
			"is_active": true,
			"is_staff":  true,
		})
	})
	defer server.Close()

	user, err := client.GetUser(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user == nil {
		t.Fatal("expected non-nil user")
	}
	if user.Username != "admin" {
		t.Errorf("expected username 'admin', got %q", user.Username)
	}
	if user.Id == nil || *user.Id != 1 {
		t.Errorf("expected id 1, got %v", user.Id)
	}
}

func TestGetUserNotFound(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusNotFound, map[string]any{
			"detail": "Not found.",
		})
	})
	defer server.Close()

	user, err := client.GetUser(999)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if user != nil {
		t.Errorf("expected nil user on error, got %v", user)
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteUser(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDeleteUserNotFound(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusNotFound, map[string]any{
			"detail": "Not found.",
		})
	})
	defer server.Close()

	err := client.DeleteUser(999)
	if err == nil {
		t.Fatal("expected error for not found, got nil")
	}
}

func TestDeactivateUserSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":        1,
			"username":  "admin",
			"email":     "admin@test.com",
			"is_active": false,
			"is_staff":  true,
		})
	})
	defer server.Close()

	err := client.DeactivateUser(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestActivateUserSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":        1,
			"username":  "admin",
			"email":     "admin@test.com",
			"is_active": true,
			"is_staff":  true,
		})
	})
	defer server.Close()

	err := client.ActivateUser(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestResetUserPasswordSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"status": "password updated",
		})
	})
	defer server.Close()

	err := client.ResetUserPassword(1, "new-secure-password")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestResetUserPasswordError(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusBadRequest, map[string]any{
			"new_password": []string{"This password is too common."},
		})
	})
	defer server.Close()

	err := client.ResetUserPassword(1, "password")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":        1,
			"username":  "admin",
			"email":     "updated@test.com",
			"is_active": true,
			"is_staff":  true,
		})
	})
	defer server.Close()

	email := stringPtr("updated@test.com")
	user, err := client.UpdateUser(1, email, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user == nil {
		t.Fatal("expected non-nil user")
	}
	if string(user.Email) != "updated@test.com" {
		t.Errorf("expected email 'updated@test.com', got %q", user.Email)
	}
}
