package users_test

import (
	"testing"

	"modular-monolith/internal/auth/domain/users"
)

func TestUsersRepository_GetAll(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		want    []users.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var u users.UserRepository
			got, gotCount, gotErr := u.GetAll(t.Context(), 20, 0)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAll() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAll() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetAll() = %v, %v, want %v", got, gotCount, tt.want)
			}
		})
	}
}
