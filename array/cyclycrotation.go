package array

//CyclycRotation Function
func CyclycRotation(A []int, K int) []int {
	var result []int
	modulus := K % len(A)
	if len(A) == 0 || len(A) == 1 {
		result = A
	} else {
		sawal := A[0 : len(A)-modulus]

		for index := 1; index <= modulus; index++ {
			s := A[len(A)-index : len(A)]
			result = append(s, sawal...)
		}
	}

	return result
}
