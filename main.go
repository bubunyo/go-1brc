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
	"sync"
)

type line struct {
	l string
	i int
}

type cityInfo struct {
	min, acc, max float64
	count         int64
}

func Main(input string) string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make(chan line, 100000)

	var wg sync.WaitGroup

	wg.Add(1)

	var cities []string
	results := map[string]*cityInfo{}

	go func() {
		for l := range lines {
			city := l.l[:l.i]
			temp, _ := strconv.ParseFloat(l.l[l.i+1:], 64)

			if c, ok := results[city]; ok {
				if temp > c.max {
					c.max = temp
				}
				if temp < c.min {
					c.min = temp
				}
				c.acc += temp
				c.count += 1
			} else {
				results[city] = &cityInfo{
					max:   temp,
					min:   temp,
					acc:   temp,
					count: 1,
				}
				cities = append(cities, city)
			}
		}
		wg.Done()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p := scanner.Text()
		i := strings.Index(p, ";")
		lines <- line{p, i}
	}

	close(lines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	sort.Strings(cities)
	var sb strings.Builder

	sb.WriteString("{")
	for i, city := range cities {
		c := results[city]
		avg := c.acc / float64(c.count)
		avg = math.Ceil(avg*10) / 10
		sb.WriteString(fmt.Sprintf("%s=%.1f/%.1f/%.1f", city, c.min, avg, c.max))
		if i != len(cities)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("}\n")

	return sb.String()
}
