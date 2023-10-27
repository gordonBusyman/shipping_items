package internal

// CalculatePacks calculates the number of packs required for a given number of items
func (s *Store) CalculatePacks(itemsOrdered int) map[int]int {
	// Ensure pack sizes are sorted in descending order
	s.Sort()

	packs := make(map[int]int)
	remainingItems := itemsOrdered

	// *****
	// BEGIN: Catching easy cases to avoid unnecessary calculations
	// *****

	// If there are no packs available, return an empty map
	if remainingItems == 0 {
		return packs
	}

	// If the number of items ordered is less than the smallest pack size, return the smallest pack size
	if remainingItems < s.Smallest() {
		packs[s.Smallest()]++

		return packs
	}

	// If there is only one pack size available, return the number of packs required
	if len(packs) == 1 {
		if numPacks := remainingItems / packs[0]; numPacks > 0 {
			packs[packs[0]] = numPacks
			if remainingItems > 0 {
				packs[packs[0]]++
			}

			return packs
		}
	}
	// *****
	// END
	// *****

	for _, size := range s.PacksAvailable {
		// Calculate the number of packs of the current size
		if numPacks := remainingItems / size; numPacks > 0 {
			packs[size] = numPacks
			remainingItems -= numPacks * size
		}

		if remainingItems == 0 {
			return packs
		}
	}

	// If there are remaining items after using all available pack sizes
	if remainingItems > 0 {
		if _, ok := packs[s.Smallest()]; ok && 2*s.Smallest() >= s.PacksAvailable[len(s.PacksAvailable)-2] {
			packs[s.PacksAvailable[len(s.PacksAvailable)-2]]++
			packs[s.Smallest()]--
		} else {
			packs[s.Smallest()]++
		}
	}

	for k, v := range packs {
		if v == 0 {
			delete(packs, k)
		}
	}

	return packs
}
