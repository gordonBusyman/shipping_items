package internal

import (
	"os"
	"testing"
)

func TestAvailablePacks(t *testing.T) {
	store := &Store{DBName: "test.json"}

	// Write some test data to the file
	err := store.WritePacks([]int{1, 2, 3})
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	// Test AvailablePacks function
	packs, err := store.AvailablePacks()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expected := []int{1, 2, 3}
	for i, pack := range packs {
		if pack != expected[i] {
			t.Errorf("Expected '%d', but got '%d'", expected[i], pack)
		}
	}

	// Clean up after test
	os.Remove("test.json")
}

func TestWritePacks(t *testing.T) {
	store := &Store{DBName: "test.json"}

	// Test WritePacks function
	err := store.WritePacks([]int{1, 2, 3})
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	// Clean up after test
	os.Remove("test.json")
}
