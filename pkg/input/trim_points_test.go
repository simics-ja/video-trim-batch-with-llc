package input

import (
	"fmt"
	"testing"
)

func TestTrimPointsReader(t *testing.T) {
	filePath := "./sample_file/time_in_second.mp4.csv"

	csv := NewCsv(filePath)
	records := csv.GetTrimPoints()
	defer csv.Close()

	expected := [][]string{
		{"3.04", "6.08"},
		{"9.12", "12.16"},
	}

	for i, record := range records {
		fmt.Printf("%v comp %v\n", record[0], expected[i][0])
		if record[0] != expected[i][0] {
			t.Errorf("records[%v][0] should be %v", i, expected[i][0])
		}
		fmt.Printf("%v comp %v\n", record[1], expected[i][1])
		if record[1] != expected[i][1] {
			t.Errorf("records[%v][1] should be %v", i, expected[i][1])
		}
	}
}
