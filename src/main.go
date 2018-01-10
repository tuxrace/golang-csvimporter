package main

import (
	"os"

	"./customerimporter"
)

func main() {
	filename := os.Args[1]
	domain := os.Args[2]
	customerimporter.Start(filename, domain)
}
