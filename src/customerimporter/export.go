package customerimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func lookup(record []string) {
	fmt.Println(record[0])
}

// Start function
func Start() {
	count := 0

	data, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := csv.NewReader(data)
	t0 := time.Now()
	for {
		record, err := reader.Read()
		go lookup(record)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		count++
		//fmt.Println(count)
	}
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
}
