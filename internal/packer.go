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

	// if it fits exactly to the biggest pack size, return the number of packs required
	biggest := s.Biggest()

	n := itemsOrdered / biggest
	rest := itemsOrdered - n*biggest
	if rest == 0 {
		packs[s.Biggest()] = n

		return packs
	}

	// *****
	// END
	// *****

	packs = map[int]int{biggest: n}

	// first possible scenario is to fit the rest items to the biggest pack size
	p := []Possible{{
		extra:         biggest - rest,
		amountOfPacks: 1,
		packs:         map[int]int{biggest: 1},
	},
	}

	tmpPacks := make(map[int]int)

	// calculate all possible pack combinations
	calc(rest, tmpPacks, s.PacksAvailable[1:], &p)

	minExtra, minAmountOfPacks := biggest-rest, 1

	// find the best possible combination
	for _, v := range p {
		if v.extra <= minExtra && v.amountOfPacks <= minAmountOfPacks {
			minExtra = v.extra
			minAmountOfPacks = v.amountOfPacks
		}
	}

	for _, v := range p {
		if v.extra == minExtra && v.amountOfPacks == minAmountOfPacks {
			for k, v1 := range v.packs {
				if k == biggest {
					packs[k]++ // if the biggest pack size is already in the map, add one more
				} else {
					packs[k] = v1 // add the rest pack sizes
				}
			}
		}
	}

	// remove 0 values
	for k, v := range packs {
		if v == 0 {
			delete(packs, k)
		}
	}

	return packs
}

// Possible is a struct to hold possible pack combinations.
type Possible struct {
	packs         map[int]int
	extra         int
	amountOfPacks int
}

// calc is a recursive function to calculate all possible pack combinations.
func calc(rest int, packs map[int]int, availablePackSizes []int, p *[]Possible) {
	amountOfPacks := rest / availablePackSizes[0]
	extra := rest - amountOfPacks*availablePackSizes[0]

	packs[availablePackSizes[0]] = amountOfPacks

	// exact match, perfect, here we are
	if extra == 0 {
		*p = append(*p, Possible{
			packs:         packs,
			extra:         0,
			amountOfPacks: amountOfPacks,
		})

		return
	}

	// we are on the smallest pack size, and we have extra items
	isSmallest := len(availablePackSizes) == 1
	extraLessThanSmallestPack := availablePackSizes[0]-extra < availablePackSizes[len(availablePackSizes)-1]

	// IF:
	// 1. we are on the smallest pack size, and we have extra items
	// 2. the extra items are less than the smallest pack size
	if isSmallest || extraLessThanSmallestPack {
		packs[availablePackSizes[0]]++ // add one more pack to fit the rest items

		*p = append(*p, Possible{
			packs:         packs,
			extra:         availablePackSizes[0] - extra,
			amountOfPacks: amountOfPacks + 1,
		})

		return
	}

	// the extra items are more than the smallest pack size, continue recursion
	calc(extra, packs, availablePackSizes[1:], p)

	return
}
