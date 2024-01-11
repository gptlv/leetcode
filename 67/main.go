package main

import "strconv"

func main() {

}

func addBinary(a string, b string) string {
	da, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		return ""
	}

	db, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return ""
	}

	sum := da + db

	return strconv.FormatInt(int64(sum), 2)

}
