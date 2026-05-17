package carfleet

import (
	"sort"
)

func CarFleet(target int, position, speed []int) int {
	cars := make([][2]int, len(position))
	for i := range position {
		cars[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] < cars[j][0]
	})

	fleets := len(cars)

	currentCarTimeToDestination := float32(
		target-cars[len(cars)-1][0],
	) / float32(
		cars[len(cars)-1][1],
	)

	for i := len(cars) - 1; i > 0; i-- {
		previousCarTimeToDestination := float32(target-cars[i-1][0]) / float32(cars[i-1][1])
		if previousCarTimeToDestination <= currentCarTimeToDestination {
			fleets--
		} else {
			currentCarTimeToDestination = previousCarTimeToDestination
		}
	}

	return fleets
}
