package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Habil Arifin", "Habil"))
	fmt.Println(strings.Split("Muhammad Habil Arifin", " "))
	fmt.Println(strings.ToLower("Muhammad Habil Arifin"))
	fmt.Println(strings.ToUpper("Muhammad Habil Arifin"))
	fmt.Println(strings.Trim("    Muhammad Habil Arifin    ", " "))
	fmt.Println(strings.ReplaceAll("Muhammad Habil Arifin Habil Arifin", "Habil", "Arya"))
}