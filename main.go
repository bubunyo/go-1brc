package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync/atomic"
)

func Main() {
	file, err := os.Open("./weather_stations.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var count int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		atomic.AddInt64(&count, 1)
	}

	fmt.Println(">>", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
