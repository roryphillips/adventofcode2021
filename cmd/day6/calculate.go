package main

func calculatePopulation(ages []int, days int, spawnTimer int, initialPause int) uint64 {
	maxAge := spawnTimer + initialPause
	ageCounts := make([]uint64, maxAge)
	for _, age := range ages {
		ageCounts[age] += 1
	}

	// Perform the following mutation logic for each day
	for i := 0; i < days; i++ {
		// Because we start at the oldest age and rewrite in place, we default to no carry
		var carry uint64 = 0
		// Starting from our youngest possible fish
		for age := maxAge - 1; age >= 0; age -= 1 {
			// Get the current population of this age bracket
			population := ageCounts[age]
			// Set the new population for this bracket to be the carryover from the previous population
			ageCounts[age] = carry
			// If the age is 0, we're spawning new fish and resetting the timer
			if age == 0 {
				// Add a set of fish that require the initial pause
				ageCounts[maxAge-1] += population
				// Skip the initial pause for our maturing fish
				// Assume there may be fish already occupying this space, so add to the bracket don't overwrite
				ageCounts[maxAge-initialPause-1] += population
			} else {
				// If the age is not 0, we need to carryover the population into the younger age bracket
				carry = population
			}
		}
	}

	var sum uint64 = 0
	for _, age := range ageCounts {
		sum += age
	}
	return sum
}
