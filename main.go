package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Main(input string) string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var tm = map[string][]float64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p := scanner.Text()
		i := strings.Index(p, ";")
		city := p[:i]
		temp, _ := strconv.ParseFloat(p[i+1:], 64)

		if _, ok := tm[city]; ok {
			tm[city] = append(tm[city], temp)
		} else {
			tm[city] = []float64{temp}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var cities []string
	results := map[string]string{}

	for city, temps := range tm {
		sort.Float64s(temps)
		var avg float64
		for _, temp := range temps {
			avg += temp
		}
		avg = avg / float64(len(temps))
		avg = math.Ceil(avg*10) / 10
		cities = append(cities, city)
		results[city] = fmt.Sprintf("%s=%.1f/%.1f/%.1f", city, temps[0], avg, temps[len(temps)-1])
	}
	sort.Strings(cities)
	var sb strings.Builder

	sb.WriteString("{")
	for i, city := range cities {
		sb.WriteString(results[city])
		if i != len(cities)-1 {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("}\n")

	return sb.String()
}
