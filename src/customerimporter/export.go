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

func lookup(record []string, domain string) {
	if len(record) > 0 && domain == record[2] {
		fmt.Println(record)
	}
	count++
}

// Start function
func Start() {
	domain := "gmail.com"

	data, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := csv.NewReader(data)
	t0 := time.Now()
	for {
		record, err := reader.Read()
		go lookup(record, domain)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
	fmt.Println("Number of records ", count)
}
