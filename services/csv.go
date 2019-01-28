package services

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

// CreateCSVWriterFile asd
func CreateCSVWriterFile(fileName string) *csv.Writer {
	err := os.Chdir("../csv-files")
	if err != nil {
		log.Println(err)
	}
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	w := csv.NewWriter(file)
	return w
}

// WriteColumnHeaders sizes of columns
func WriteColumnHeaders(columns []string, csvWriter *csv.Writer) {
	// for idx, c := range columns {
	// 	columns = append(columns, c)
	// 	fmt.Println(idx, columns)
	// }

	err := csvWriter.Write(columns)
	if err != nil {
		log.Fatalf("Error writing columns %s", err)
	}
}

const (
	unixDate = "Mon Jan _2 15:04:05 MST 2006"
)

// ParseUnixAsDate asd
func ParseUnixAsDate(val int64) string {
	if val == 0 {
		return "%"
	}
	date := time.Unix((val / 1000), 0)
	fmt.Println(date)
	return date.String()
}
