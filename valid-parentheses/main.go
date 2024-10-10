package main

func main() {

}

func isValid(s string) bool {
	q := []byte{}

	m := make(map[byte]byte)

	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'

	for i := 0; i < len(s); i++ {
		if isClosing(s[i]) {
			if len(q) == 0 {
				return false
			}

			last := q[len(q)-1]
			if m[last] != s[i] {
				return false
			}
			q = q[:len(q)-1]

		} else {
			q = append(q, s[i])
		}
	}

	return len(q) == 0
}

func isClosing(bracket byte) bool {
	return bracket == ')' || bracket == ']' || bracket == '}'
}
