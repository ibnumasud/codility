package timecomplexity

//FrogJump Fucntion
func FrogJump(x int, y int, d int) int {
	var result int
	distance := y - x
	result = distance / d
	if distance%d != 0 {
		result++
	}
	return result
}

//PermMissElement function
func PermMissElement(A []int) int {
	total := (len(A) + 1) * (len(A) + 2) / 2
	for _, a := range A {
		total -= a
	}
	return total
}
