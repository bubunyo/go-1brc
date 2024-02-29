package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const width = 3000

type line struct {
	l   string
	i   int
	idx uint32
}

type cityIdx struct {
	c   string
	idx uint32
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

	lines := make(chan line, 10000)
	citiesIn := make(chan cityIdx, 10000)

	var wg sync.WaitGroup

	var cities []string
	cityIdxhash := map[string]uint32{}
	results := [width]map[string]*cityInfo{}
	chanMap := [width]chan line{}

	for i := 0; i < width; i++ {
		chanMap[i] = make(chan line)
		results[i] = map[string]*cityInfo{}
	}

	// handle cities
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := range citiesIn {
			cities = append(cities, c.c)
			cityIdxhash[c.c] = c.idx
		}
	}()

	var mwg sync.WaitGroup

	mwg.Add(width)
	for i := 0; i < width; i++ {
		go func(i int) {
			defer mwg.Done()
			for l := range chanMap[i] {
				city := l.l[:l.i]
				temp, _ := strconv.ParseFloat(l.l[l.i+1:], 64)

				if c, ok := results[l.idx][city]; ok {
					if temp > c.max {
						c.max = temp
					}
					if temp < c.min {
						c.min = temp
					}
					c.acc += temp
					c.count += 1
				} else {
					results[l.idx][city] = &cityInfo{
						max:   temp,
						min:   temp,
						acc:   temp,
						count: 1,
					}
					citiesIn <- cityIdx{city, l.idx}
				}
			}
		}(i)
	}

	go func() {
		mwg.Wait()
		close(citiesIn)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for l := range lines {
			city := l.l[:l.i]
			l.idx = hash(city) % width
			chanMap[l.idx] <- l
		}
		for i := range chanMap {
			close(chanMap[i])
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		i := strings.Index(l, ";")
		lines <- line{l: l, i: i}
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
		idx := cityIdxhash[city]
		c := results[idx][city]
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

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
