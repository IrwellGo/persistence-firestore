package persistence_firestore

import (
	"errors"
	"testing"
)

type TestObject struct {
	name string
}

func TestFakeFuncForUnitTesting(t *testing.T) {
	// Assign

	// Act
	pass := FakeFuncForUnitTesting(true)

	// Assert
	if pass != true {
		t.Error(errors.New("it failed"))
	}
}
