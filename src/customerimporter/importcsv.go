package customerimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var count = 0

func filterDomain(record []string, domain string) {
	if len(record) > 0 && domain == record[2] {
		fmt.Println(record)
	}
	count++
}

// Start function
func Start(filename string, domain string) {
	csvData, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvData.Close()
	reader := csv.NewReader(csvData)
	t0 := time.Now()
	for {
		record, err := reader.Read()
		go filterDomain(record, domain)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	t1 := time.Now()
	fmt.Println("Execution time    : ", t1.Sub(t0))
	fmt.Println("Number of records : ", count)
}
