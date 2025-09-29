package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`h([a-z])l`)

	fmt.Println(regex.MatchString("hil"))
	fmt.Println(regex.MatchString("Hel"))
	fmt.Println(regex.MatchString("Halil"))

	fmt.Println(regex.FindAllString("hal hil hul hel hol hAl", 10))
}
