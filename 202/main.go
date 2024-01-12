package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isHappy(13))
}

var m = make(map[int]int)

func isHappy(n int) bool {

	log := math.Log10(float64(n))
	if log == math.Trunc(log) {
		return true
	}

	s := strconv.Itoa(n)
	ss := strings.Split(s, "")
	sum := 0

	for _, snum := range ss {
		num, _ := strconv.Atoi(snum)
		sum += num * num
	}

	if _, ok := m[sum]; ok {
		return false
	}

	m[sum] = sum

	return isHappy(sum)

}

// log := math.Log10(float64(num))
// if log == math.Trunc(log) {
// 	return true
// }
