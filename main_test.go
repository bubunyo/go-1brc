package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_WithData(t *testing.T) {
	res := Main("./test_cases/measurements-1.txt")
	out := readFile("./test_cases/measurements-1.out")
	assert.Equal(t, out, res)
}

func TestMain_OutputFile(t *testing.T) {
	f := "measurements-10000-unique-keys.txt"
	res := Main("./test_cases/" + f)
	outputToFile(f, res)
}

func TestMain_WithData_NoOut(t *testing.T) {
	_ = Main("./data/measurements-1000000000.txt")
}

func TestMain(t *testing.T) {
	inputFiles := find("./test_cases", ".txt")
	for _, test := range inputFiles {
		t.Run(test, func(t *testing.T) {
			assert.Equal(t, readFile(test+".out"), Main(test+".txt"))
		})
	}
}

func TestMain_1br(t *testing.T) {
	fmt.Println("Res:", Main("./data/measurements-1000000000.txt"))
}

func outputToFile(name, out string) {
	name = "./out/" + name + ".out"
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("File does not exists or cannot be created. file=%s", name)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	fmt.Fprintf(w, "%v\n", out)
	w.Flush()
}

func readFile(input string) string {
	fileContent, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return string(fileContent)
}

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s[:len(s)-len(ext)])
		}
		return nil
	})
	return a
}
