// elections: show election results on a state grid
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Data file structure
type egrid struct {
	name       string
	party      string
	row        int
	col        int
	population int
}

var partymap = map[string]string{"r": "red", "d": "blue", "i": "gray"}

// vmap maps one range into another
func vmap(value, low1, high1, low2, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

// area computes the area of a circle
func area(v float64) float64 {
	return math.Sqrt((v / math.Pi)) * 2
}

// atoi converts a string to an integer
func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

// readData reads election data into the data structure
func readData(r io.Reader) ([]egrid, int, int, string, error) {
	var d egrid
	var data []egrid
	title := ""
	scanner := bufio.NewScanner(r)
	min, max := math.MaxInt32, -math.MaxInt32
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 { // skip blank lines
			continue
		}
		if t[0] == '#' && len(t) > 2 { // get the title
			title = t[2:]
			continue
		}
		fields := strings.Split(t, "\t")
		if len(fields) < 5 { // skip incomplete records
			continue
		}
		d.name = fields[0]
		d.col = atoi(fields[1])
		d.row = atoi(fields[2])
		d.party = fields[3]
		d.population = atoi(fields[4])
		data = append(data, d)
		if d.population > max {
			max = d.population
		}
		if d.population < min {
			min = d.population
		}
	}
	return data, min, max, title, scanner.Err()
}

// process walks the data, making the visualization
func process(startx, starty, rowsize, colsize float64, data []egrid, min, max int, title string) {
	amin := area(float64(min))
	amax := area(float64(max))
	beginPage("black", "white")
	showtitle(title, starty+15)
	for _, d := range data {
		x := startx + (float64(d.row) * colsize)
		y := starty - (float64(d.col) * rowsize)
		r := vmap(area(float64(d.population)), amin, amax, 2, colsize)
		circle(x, y, r, partymap[d.party])
		ctext(x, y-0.5, 1.2, d.name, "white")
	}
	endPage()
}

// showtitle shows the title and subhead
func showtitle(s string, top float64) {
	fields := strings.Fields(s)
	if len(fields) != 3 {
		return
	}
	ctext(50, top, 3.6, fields[0]+" US Presidential Election", "")
	legend(20, top-5, 2.0, fields[1], partymap["d"])
	legend(80, top-5, 2.0, fields[2], partymap["r"])
}

// circle makes a circle
func circle(x, y, r float64, color string) {
	fmt.Printf("circle %g %g %.4g \"%s\"\n", x, y, r, color)
}

// ctext makes centered text
func ctext(x, y, ts float64, s string, color string) {
	fmt.Printf("ctext \"%s\" %g %g %g \"sans\" \"%s\"\n", s, x, y, ts, color)
}

// ltext makes left-aligned text
func ltext(x, y, ts float64, s string, color string) {
	fmt.Printf("text \"%s\" %g %g %g \"sans\" \"%s\"\n", s, x, y, ts, color)
}

// legend makes the subtitle
func legend(x, y, ts float64, s string, color string) {
	ltext(x, y, ts, s, "")
	circle(x-ts, y+ts/3, ts/2, color)
}

// beginPage starts a page
func beginPage(bgcolor, textcolor string) {
	fmt.Printf("slide \"%s\" \"%s\"\n", bgcolor, textcolor)
}

// endPage ends a page
func endPage() {
	ctext(50, 5, 1.5, "Area of circle denotes state population: source U.S. Census", "gray")
	fmt.Println("eslide")
}

// beginDoc starts the document
func beginDoc() {
	fmt.Println("deck")
}

// endDoc ends the document
func endDoc() {
	fmt.Println("edeck")
}

func main() {
	beginDoc()
	for _, f := range os.Args[1:] {
		r, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		data, min, max, title, err := readData(r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		process(10, 75, 9, 7, data, min, max, title)
	}
	endDoc()
}