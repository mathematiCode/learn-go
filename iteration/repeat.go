package iteration

import "strings"

func Repeat(char string, numTimes int) string {
	var result strings.Builder
	for i := 1; i <= numTimes; i++ {
		result.WriteString(char)
	}
	return result.String()
}
