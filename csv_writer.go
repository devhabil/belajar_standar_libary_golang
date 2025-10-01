package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Muhammad", "Habil", "Arifin"})
	_ = writer.Write([]string{"Eko", "Kurniawan", "Khanedy"})
	_ = writer.Write([]string{"Muhammad", "Kules", "Saputra"})

	writer.Flush()
}
