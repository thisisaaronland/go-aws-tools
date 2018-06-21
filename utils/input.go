package utils

import (
       "fmt"
       "strings"
)

func Readline(prompt string) string {

	var input string

	fmt.Print(fmt.Sprintf("%s ", prompt))
	fmt.Scanf("%s", &input)

	// go-sanitize strings here? (20180621/thisisaaronland)

	return strings.Trim(input, " ")
}
