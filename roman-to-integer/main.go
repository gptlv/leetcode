package main

func main() {

}

func romanToInt(s string) int {
	m := make(map[uint8]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	res := 0
	ptr := 0

	for ptr < len(s) {
		condition := ptr != len(s)-1 && ((s[ptr] == 'I' && (s[ptr+1] == 'V' || s[ptr+1] == 'X')) || (s[ptr] == 'X' && (s[ptr+1] == 'L' || s[ptr+1] == 'C')) || (s[ptr] == 'C' && (s[ptr+1] == 'D' || s[ptr+1] == 'M')))
		if condition {
			res += m[s[ptr+1]] - m[s[ptr]]
			ptr += 2
		} else {
			res += m[s[ptr]]
			ptr++
		}
	}

	return res
}
