package main

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	total, start, tank := 0, 0, 0

	for i := 0; i < length; i++ {
		tank += gas[i] - cost[i]
		total += gas[i] - cost[i]

		if tank <= 0 {
			start = (i + 1) % len(gas)
			tank = 0
		}
	}

	if total < 0 {
		return -1
	}

	return start
}
