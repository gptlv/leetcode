package main

func main() {

}

func dailyTemperatures(temperatures []int) []int {
	t := temperatures
	s := stack{}
	res := make([]int, len(t))

	for i := 0; i < len(t); i++ {
		for len(s) > 0 && t[i] > t[s[len(s)-1]] {
			idx := s.pop()
			res[idx] = i - idx
		}
		s.push(i)
	}

	return res
}

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() int {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}
