package customerimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Start function
func Start() {
	count := 0
	data, err := os.Open("big.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := csv.NewReader(data)
	t0 := time.Now()
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		count++
		fmt.Println(count, record)
	}
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
}
