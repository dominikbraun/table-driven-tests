package main

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStorage_FindUser(t *testing.T) {
	tests := map[string]struct {
		users        map[string]User // The given underlying user storage.
		id           int             // The desired user ID.
		expectedUser User            // The expected User instance.
		expectedErr  error           // The error that is expected to be returned.
	}{
		// In case the user exists, FindUser is expected to return the User instance with
		// the given ID (in this case, user #2). The call should return no error.
		"user exists": {
			users: map[string]User{
				"User #1": {ID: 1, Name: "User #1"},
				"User #2": {ID: 2, Name: "User #2"},
				"User #3": {ID: 3, Name: "User #3"},
			},
			id:           2,
			expectedUser: User{ID: 2, Name: "User #2"},
			expectedErr:  nil,
		},
		// In case the user doesn't exist, FindUser is expected to return an error.
		"user doesn't exist": {
			users: map[string]User{
				"User #1": {ID: 1, Name: "User #1"},
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

func TestStorage_AddUser(t *testing.T) {
	// t.Parallel() signals that this test is to be run in parallel with (and only with) other parallel tests.
	t.Parallel()

	tests := map[string]struct {
		users       map[string]User // The given underlying user storage.
		id          int             // The user id to add.
		name        string          // The user name to add.
		expectedErr error           // The error that is expected to be returned.
	}{
		// In case the user doesn't exist, AddUser is expected to return no error.
		"user doesn't exist": {
			users: map[string]User{
				"User #1": {ID: 1, Name: "User #1"},
				"User #2": {ID: 2, Name: "User #2"},
				"User #3": {ID: 3, Name: "User #3"},
			},
			id:          4,
			name:        "User #4",
			expectedErr: nil,
		},
		// In case of an already existing user, AddUser is expected to return an error.
		"user exists": {
			users: map[string]User{
				"User #1": {ID: 1, Name: "User #1"},
			},
			id:          1,
			name:        "User #1",
			expectedErr: ErrUserAlreadyExists,
		},
	}

	for name, test := range tests {
		// Assign a new local variable because all tests will run in parallel.
		// Otherwise this can lead to unexpected test results, because the test value gets overwritten by the
		// next for loop iteration while a previous tests is already running. This is called "race condition".
		// If you comment this line out, you may find that running `go test -cover ./...` doesn't yield 100% anymore.
		test := test

		t.Run(name, func(t *testing.T) {
			// Call t.Parallel() in every subtest, so all subtests can run in parallel.
			t.Parallel()

			// Create the storage using the underlying storage provided by the test case.
			storage := NewStorage(test.users)

			// Perform the call using the given arguments.
			err := storage.AddUser(test.id, test.name)

			// Use errors.Is to determine if the expected error was returned.
			// This also works with nil errors.
			if !errors.Is(err, test.expectedErr) {
				t.Fatalf("error expectancy doesn't match: expectedErr is %v, but got %v", test.expectedErr, err)
			}
		})
	}
}
