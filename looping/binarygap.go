package looping

import (
	"strconv"
	"strings"
)

//BinaryGap fungction
func BinaryGap(N int) int {
	gap := 0
	n := strconv.FormatInt(int64(N), 2)
	s := strings.Split(n, "1")
	for i := 1; i < len(s); i++ {
		if len(s[i]) > gap {
			gap = len(s[i])
		}
	}
	return gap
}
