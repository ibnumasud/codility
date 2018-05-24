package timecomplexity

import (
	"fmt"
	"math"
)

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

//TapeEquilibirium function
func TapeEquilibirium(A []int) int {
	var result int
	head := A[0]
	tail := 0
	for i := 1; i <= len(A)-1; i++ {
		tail += A[i]
	}
	result = int(math.Abs(float64(head - tail)))
	fmt.Println(result)

	for i := 1; i <= len(A)-1; i++ {
		head += A[i]
		tail -= A[i]
		if result > int(math.Abs(float64(head-tail))) {
			result = int(math.Abs(float64(head - tail)))
		}

	}

	return result
}
