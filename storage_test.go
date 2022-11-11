package main

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStorage_FindUser(t *testing.T) {
	tests := map[string]struct {
		users        []User // The given underlying user storage.
		id           int    // The desired user ID.
		expectedUser User   // The expected User instance.
		expectedErr  error  // The error that is expected to be returned.
	}{
		// In case the user exists, FindUser is expected to return the User instance with
		// the given ID (in this case, user #2). The call should return no error.
		"user exists": {
			users: []User{
				{ID: 1, Name: "User #1"},
				{ID: 2, Name: "User #2"},
				{ID: 3, Name: "User #3"},
			},
			id:           2,
			expectedUser: User{ID: 2, Name: "User #2"},
			expectedErr:  nil,
		},
		// In case the user doesn't exist, FindUser is expected to return an error.
		"user doesn't exist": {
			users: []User{
				{ID: 1, Name: "User #1"},
			},
			id:           2,
			expectedUser: User{},
			expectedErr:  ErrUserNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Create the storage using the underlying storage provided by the test case.
			storage := NewStorage(test.users)

			// Perform the call using the given arguments and store the results.
			user, err := storage.FindUser(test.id)

			// Use errors.Is to determine if the expected error was returned.
			// This also works with nil errors.
			if !errors.Is(err, test.expectedErr) {
				t.Fatalf("error expectancy doesn't match: expectedErr is %v, but got %v", test.expectedErr, err)
			}

			// Compare the expected and the actually retrieved user instances and store
			// the potential diff. If there is a difference, fail and print the diff.
			if diff := cmp.Diff(test.expectedUser, user); diff != "" {
				t.Errorf("user expectancy doesn't match: %v", diff)
			}
		})
	}
}
