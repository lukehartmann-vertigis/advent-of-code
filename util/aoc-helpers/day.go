package aoc_helpers

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Djosar/advent-of-code/util/collections"
)

const AOCHost = "https://adventofcode.com"

type TaskFunc = func(d *AOCDay) (any, error)

type AOCDay struct {
	Id        int
	TestInput string
	TestLines []string
	Input     string
	Lines     []string
	Matrix    *collections.Matrix[string]
	inputUrl  string
	filePath  string
	tasks     []TaskFunc
}

func NewDay(id int, year int, testInput string) *AOCDay {
	filePath := fmt.Sprintf("./%d/%d/input.txt", year, id)
	inputUrl := fmt.Sprintf("%s/%d/day/%d/input", AOCHost, year, id)
	testLines := strings.Split(testInput, "\n")

	day := &AOCDay{
		Id:        id,
		TestInput: testInput,
		TestLines: testLines,
		Lines:     []string{},
		inputUrl:  inputUrl,
		filePath:  filePath,
		tasks:     []TaskFunc{},
	}

	if err := day.ensureInputFile(); err != nil {
		fmt.Println(err)
	}

	if err := day.loadInput(); err != nil {
		fmt.Println(err)
	}

	return day
}

func (d *AOCDay) PushTask(t TaskFunc) {
	d.tasks = append(d.tasks, t)
}

func (d *AOCDay) Run() error {
	fmt.Println("==========")

	for idx, task := range d.tasks {
		res, err := task(d)
		if err != nil {
			return err
		}

		fmt.Printf("| Task %d: %v\n", idx, res)
	}

	fmt.Println("==========")
	fmt.Println()

	return nil
}

func (d *AOCDay) ensureInputFile() error {
	if _, err := os.Stat(d.filePath); err == nil {
		// fmt.Printf("Input file already exists: '%s'\n", d.filePath)
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return fmt.Errorf("AOC_SESSION env var not set")
	}

	// Ensure the directory
	if err := os.MkdirAll(filepath.Dir(d.filePath), 0o755); err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(d.filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Download the file content
	fmt.Println(d.inputUrl)
	req, err := http.NewRequest("GET", d.inputUrl, nil)
	if err != nil {
		return err
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error trying to fetch input file, error code %d", res.StatusCode)
	}

	// Copy response body into created file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (d *AOCDay) loadInput() error {
	file, err := os.Open(d.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCap = 1024 * 1024 // 1MB
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, maxCap)

	var inputBuffer strings.Builder
	var lines = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		inputBuffer.WriteString(scanner.Text() + "\n")
	}
	input := inputBuffer.String()
	d.Input = inputBuffer.String()

	matrixWidth := len(lines[0])
	d.Matrix = collections.NewMatrix(strings.Split(input, ""), matrixWidth)

	if err := scanner.Err(); err != nil {
		return err
	}

	d.Lines = lines

	return nil
}
