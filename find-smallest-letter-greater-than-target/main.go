package main

func main() {

}

func nextGreatestLetter(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	var left, right, middle int
	right = len(letters) - 1

	for left < right {
		middle = (left + right) / 2

		if letters[middle] > target {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return letters[left]
}
