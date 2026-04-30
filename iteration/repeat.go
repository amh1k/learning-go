package iteration

import "strings"
const repeatedCount = 5

func Repeat(character string, cnt int) string {
	var repeated strings.Builder

	for i:= 0; i  < cnt; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()

}