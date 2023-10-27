package internal

import "sort"

// Store is the store struct
type Store struct {
	PacksAvailable []int
	DBName         string
}

// NewStore returns a new Store
func NewStore(db string) *Store {
	return &Store{DBName: db}
}

// Smallest returns the smallest pack size available
func (s *Store) Smallest() int {
	return s.PacksAvailable[len(s.PacksAvailable)-1]
}

// Sort sorts the packs available in descending order
func (s *Store) Sort() {
	sort.Sort(sort.Reverse(sort.IntSlice(s.PacksAvailable)))
}
