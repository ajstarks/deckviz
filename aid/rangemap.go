package main

import (
	"bufio"
	"io"
	"flag"
	"os"
	"fmt"
	"strconv"
	"math"
)

const (
	maxfloat = math.MaxFloat64
	minfloat = -math.MaxFloat64
)

func readvalues(r io.Reader) ([]float64, error) {
	var data []float64
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		v, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		data = append(data, v)
	}
	return data, scanner.Err()
}

func extrema(data []float64) (float64, float64) {
	max := minfloat
	min := maxfloat
	for _, d := range data {
		if d > max {
			max = d
		}
		if d < min {
			min = d
		}
	}
	return min, max
}
		
func printmap(w io.Writer, low, high float64, data []float64) {
	min, max := extrema(data)
	for _, d := range data {
		fmt.Fprintf(w, "%.3f -> %.3f\n", d, vmap(d, min, max, low, high))
	}
}

func vmap(value, low1, high1, low2, high2 float64) float64 {
    return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

func main() {
	low := flag.Float64("min", 0, "low value")
	high := flag.Float64("max", 100, "high value")
	flag.Parse()
	data, err := readvalues(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	printmap(os.Stdout, *low, *high, data)
}
	
	
	