package internal

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	store := &Store{PacksAvailable: []int{250, 1000, 2000, 5000, 500}}

	tests := []struct {
		name          string
		itemsOrdered  int
		expectedPacks map[int]int
	}{
		{
			name:         "1 Item Ordered",
			itemsOrdered: 1,
			expectedPacks: map[int]int{
				250: 1,
			},
		},
		{
			name:         "250 Items Ordered",
			itemsOrdered: 250,
			expectedPacks: map[int]int{
				250: 1,
			},
		},
		{
			name:         "251 Items Ordered",
			itemsOrdered: 251,
			expectedPacks: map[int]int{
				500: 1,
			},
		},
		{
			name:         "501 Items Ordered",
			itemsOrdered: 501,
			expectedPacks: map[int]int{
				250: 1,
				500: 1,
			},
		},
		{
			name:         "12001 Items Ordered",
			itemsOrdered: 12001,
			expectedPacks: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},

		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packs := store.CalculatePacks(tt.itemsOrdered)
			if !reflect.DeepEqual(packs, tt.expectedPacks) {
				t.Errorf("Expected %v, but got %v", tt.expectedPacks, packs)
			}
		})
	}
}
