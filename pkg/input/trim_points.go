package input

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Csv struct {
	file *os.File
}

func NewCsv(filePath string) *Csv {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("failed to open file: %v", err)
	}
	return &Csv{f}
}

func (c *Csv) GetTrimPoints() [][]string {
	reader := csv.NewReader(c.file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("failed to read csv:", err)
	}

	return records
}

func (c *Csv) Close() {
	c.file.Close()
}
