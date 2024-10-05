package main

func main() {

}

func isIsomorphic(s string, t string) bool {
	return helper(s, t) && helper(t, s)
}

func helper(s string, t string) bool {
	m := make(map[uint8]uint8)

	for i := 0; i < len(s); i++ {
		m[s[i]] = t[i]
	}

	for i := 0; i < len(s); i++ {
		if m[s[i]] != t[i] {
			return false
		}
	}

	return true
}

// func isIsomorphic(s string, t string) bool {
// 	sm := make(map[uint8][5 * 10e4]int)
// 	tm := make(map[uint8][5 * 10e4]int)

// 	for i := 0; i < len(s); i++ {
// 		sCur := sm[s[i]]
// 		sCur[i] = 1
// 		sm[s[i]] = sCur

// 		tCur := tm[t[i]]
// 		tCur[i] = 1
// 		tm[t[i]] = tCur
// 	}

// 	for i := 0; i < len(s); i++ {
// 		if sm[s[i]] != tm[t[i]] {
// 			return false
// 		}
// 	}

// 	return true
// }
