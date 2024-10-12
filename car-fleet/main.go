package main

import (
	"sort"
)

func main() {

}

type Fleet struct {
	position int
	speed    int
}

func (f Fleet) GetDelta(target int) float64 {
	return float64((target - f.position)) / float64(f.speed)
}

type Fleets []Fleet

func (f Fleets) Less(i, j int) bool {
	return f[i].position < f[j].position
}

func (f Fleets) Len() int {
	return len(f)
}

func (f Fleets) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func carFleet(target int, position []int, speed []int) int {
	fleets := Fleets{}
	stack := []float64{}

	for i := range position {
		fleets = append(fleets, Fleet{position[i], speed[i]})
	}

	sort.Sort(sort.Reverse(fleets))

	for _, fleet := range fleets {
		arriving := fleet.GetDelta(target)

		if len(stack) == 0 || arriving > stack[len(stack)-1] {
			stack = append(stack, arriving)
		}
	}

	return len(stack)
}
