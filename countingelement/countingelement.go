package countingelement

//FrogRiverOne function
func FrogRiverOne(X int, A []int) int {
	temp := make(map[int]bool)
	todo := X
	for i := 0; i < len(A); i++ {
		internal := A[i] - 1
		// fmt.Println(internal)
		_, exist := temp[internal]
		if !exist {
			todo--
			temp[internal] = true
		}
		if todo == 0 {
			return i
		}
	}
	return -1
}

//MissingInteger function
func MissingInteger(A []int) int {
	minimum := 1
	temp := make(map[int]bool)
	for _, a := range A {
		if a > 0 {
			temp[a] = true
		}
	}

	for i := 0; i < len(temp); i++ {
		_, exist := temp[minimum]
		if exist {
			minimum++
		}
	}
	return minimum
}
