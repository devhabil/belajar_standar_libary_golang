package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Muhammad,Habil,arifin\n" +
		"Eko,Kurniawan,Khanedy\n" +
		"Joku,Jeki,Juju"

	reader := csv.NewReader(strings.NewReader((csvString)))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
