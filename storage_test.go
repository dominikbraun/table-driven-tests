package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStorage_FindUser(t *testing.T) {
	tests := map[string]struct {
		users        []User // The given underlying user storage.
		id           int    // The desired user ID.
		expectedUser User   // The expected User instance.
		shouldFail   bool   // Indicates whether the function should return an error.
	}{
		// In case the user exists, FindUser is expected to return the User instance with
		// the given ID (in this case, user #2). The call should not fail.
		"user exists": {
			users: []User{
				{ID: 1, Name: "User #1"},
				{ID: 2, Name: "User #2"},
				{ID: 3, Name: "User #3"},
			},
			id:           2,
			expectedUser: User{ID: 2, Name: "User #2"},
			shouldFail:   false,
		},
		// In case the user doesn't exist, FindUser is expected to return an error.
		"user doesn't exist": {
			users: []User{
				{ID: 1, Name: "User #1"},
			},
			id:           2,
			expectedUser: User{},
			shouldFail:   true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Create the storage using the underlying storage provided by the test case.
			storage := NewStorage(test.users)

			// Perform the call using the given arguments and store the results.
			user, err := storage.FindUser(test.id)

			// Check whether the expression (err != nil) matches the expectation from the
			// test case. That way, the test fails if the test case expects an error and
			// no error is returned, or if the test case doesn't expect an error and an
			// error is returned.
			if test.shouldFail != (err != nil) {
				t.Fatalf("error expectancy doesn't match: shouldFail is %v, but error != nil is %v", test.shouldFail, err != nil)
			}

			// Compare the expected and the actually retrieved user instances and store
			// the potential diff. If there is a difference, fail and print the diff.
			if diff := cmp.Diff(test.expectedUser, user); diff != "" {
				t.Errorf("user expectancy doesn't match: %v", diff)
			}
		})
	}
}
