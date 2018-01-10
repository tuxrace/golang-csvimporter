package main

import (
	"os"

	"./customerimporter"
)

func main() {
	filename := os.Args[1]
	customerimporter.Start(filename)
}
