package array

//OddOccuranceArray Function
func OddOccuranceArray(A []int) int {
	var result int
	dupMap := dupCount(A)
	for i, value := range dupMap {
		if value%2 != 0 {
			result = i
		}
	}
	return result
}

func dupCount(list []int) map[int]int {
	dupFrequency := make(map[int]int)
	for _, item := range list {
		_, exist := dupFrequency[item]
		if exist {
			dupFrequency[item]++
		} else {
			dupFrequency[item] = 1

		}
	}
	return dupFrequency

}
