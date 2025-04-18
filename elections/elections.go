// elections: show election results on a state grid,
// using deck markup.
// elections [opts] file... | pdfdeck ... -symbol stateface ...
package main

import (
	"bufio"
	"flag"
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

// command line options
type options struct {
	top, left, rowsize, colsize float64
	bgcolor, textcolor, shape   string
}

// character map for the Stateface font
var statemap = map[string]string{
	"AL": "B",
	"AK": "A",
	"AZ": "D",
	"AR": "C",
	"CA": "E",
	"CO": "F",
	"CT": "G",
	"DE": "H",
	"FL": "I",
	"GA": "J",
	"HI": "K",
	"ID": "M",
	"IL": "N",
	"IN": "O",
	"IA": "L",
	"KS": "P",
	"KY": "Q",
	"LA": "R",
	"ME": "U",
	"MD": "T",
	"MA": "S",
	"MI": "V",
	"MN": "W",
	"MS": "Y",
	"MO": "X",
	"MT": "Z",
	"NE": "c",
	"NV": "g",
	"NH": "d",
	"NJ": "e",
	"NM": "f",
	"NY": "h",
	"NC": "a",
	"ND": "b",
	"OH": "i",
	"OK": "j",
	"OR": "k",
	"PA": "l",
	"RI": "m",
	"SC": "n",
	"SD": "o",
	"TN": "p",
	"TX": "q",
	"UT": "r",
	"VT": "t",
	"VA": "s",
	"WA": "u",
	"WV": "w",
	"WI": "v",
	"WY": "x",
	"DC": "y",
}
var partyColors = map[string]string{"r": "red", "d": "blue", "i": "gray", "w": "peru", "dr": "purple", "f": "green"}
var width, height int

// maprange maps one range into another
func maprange(value, low1, high1, low2, high2 float64) float64 {
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
	var min, max int
	min, max = math.MaxInt32, -math.MaxInt32
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
		// name,col,row,party,population
		d.name = fields[0]
		d.col = atoi(fields[1])
		d.row = atoi(fields[2])
		d.party = fields[3]
		d.population = atoi(fields[4])
		data = append(data, d)
		// compute min, max
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
func process(opts options, data []egrid, min, max int, title string) {
	beginPage(opts.bgcolor, opts.textcolor)
	fmin, fmax := float64(min), float64(max)
	amin, amax := area(fmin), area(fmax)
	sumpop := 0
	for _, d := range data {
		sumpop += d.population
		x := opts.left + (float64(d.row) * opts.colsize)
		y := opts.top - (float64(d.col) * opts.rowsize)
		fpop := float64(d.population)
		apop := area(fpop)

		// defaults
		txcolor := "white"
		txsize := 1.2
		font := "sans"
		name := d.name

		switch opts.shape {
		case "c": // circle
			r := maprange(apop, amin, amax, 2, opts.colsize)
			circle(x, y, r, partyColors[d.party])
		case "h": // hexagom
			r := maprange(apop, amin, amax, 2, opts.colsize)
			hexagon(x, y, r/2, partyColors[d.party])
		case "s": // square
			r := maprange(fpop, fmin, fmax, 2, opts.colsize)
			square(x, y, r, partyColors[d.party])
		case "l": // lines
			r := maprange(apop, amin, amax, 2, opts.colsize)
			polylines(x, y, r/2, 0.2, partyColors[d.party])
			txcolor = partyColors[d.party]
		case "p": // plain text
			txcolor = partyColors[d.party]
			txsize = maprange(fpop, fmin, fmax, 2, opts.colsize*0.75)
		case "g": // geographic
			txcolor = partyColors[d.party]
			name = statemap[d.name]
			font = "symbol"
			txsize = maprange(fpop, fmin, fmax, 2, opts.colsize)
		default:
			r := maprange(apop, amin, amax, 2, opts.colsize)
			circle(x, y, r, partyColors[d.party])
		}
		ctext(x, y-0.5, txsize, name, font, txcolor)
	}
	showtitle(title, sumpop, opts.top+15)
	endPage()
}

// partycand defines the party and candidate
func partycand(s, def string) (string, string) {
	var party, cand string
	f := strings.Split(s, ":")
	if len(f) > 1 {
		party = f[1]
		cand = f[0]
	} else {
		party = def
		cand = s
	}
	return party, cand
}

func million(n int) string {
	s := strconv.FormatInt(int64(n), 10)
	p := len(s)
	return s[0:p-6] + "," + s[p-6:p-3] + "," + s[p-3:p]
}

// showtitle shows the title and subhead
func showtitle(s string, pop int, top float64) {
	fields := strings.Fields(s) // year, democratic, republican, third-party (optional)

	if len(fields) < 2 {
		return
	}
	suby := top - 7
	ctext(50, top, 4, fields[0]+" US Presidential Election", "sans", "")
	ctext(90, 5, 1.5, "Population: "+million(pop), "sans", "")

	var party string
	var cand string
	if len(fields) > 1 {
		party, cand = partycand(fields[1], "d")
		legend(20, suby, 2.0, cand, partyColors[party])
	}
	if len(fields) > 2 {
		party, cand = partycand(fields[2], "r")
		legend(80, suby, 2.0, cand, partyColors[party])
	}
	if len(fields) > 3 {
		party, cand = partycand(fields[3], "i")
		legend(50, suby, 2.0, cand, partyColors[party])
	}
}

// circle makes a circle
func circle(x, y, r float64, color string) {
	fmt.Printf("<ellipse xp=\"%g\" yp=\"%g\" wp=\"%.4g\" hr=\"100\" color=\"%s\"/>\n", x, y, r, color)
}

// square makes a square centered at (x,y), size r
func square(x, y, r float64, color string) {
	fmt.Printf("<rect xp=\"%g\" yp=\"%g\" wp=\"%.4g\" hr=\"100\" color=\"%s\"/>\n", x, y, r, color)
}

// ctext makes centered text
func ctext(x, y, ts float64, s string, font string, color string) {
	fmt.Printf("<text align=\"c\" xp=\"%g\" yp=\"%g\" sp=\"%g\" font=\"%s\" color=\"%s\">%s</text>\n", x, y, ts, font, color, s)
}

// ltext makes left-aligned text
func ltext(x, y, ts float64, s string, color string) {
	fmt.Printf("<text xp=\"%g\" yp=\"%g\" sp=\"%g\" font=\"sans\" color=\"%s\">%s</text>\n", x, y, ts, color, s)
}

// legend makes the subtitle
func legend(x, y, ts float64, s string, color string) {
	ltext(x, y, ts, s, "")
	circle(x-ts, y+ts/3, ts/2, color)
}

// pangles computes the points of a polygon based on a series of angles
func pangles(cx, cy, r float64, angles []float64) ([]float64, []float64) {
	px := make([]float64, len(angles))
	py := make([]float64, len(angles))
	aspect := float64(width) / float64(height)
	for i, a := range angles {
		t := a * (math.Pi / 180)
		px[i] = cx + (r * math.Cos(t))
		py[i] = cy + ((r * aspect) * math.Sin(t))
	}
	return px, py
}

// hexagon makes a hexagon centered at (x,y), inscribes in a circle of radius r
func hexagon(cx, cy, r float64, color string) {
	// construct a polygon with points at these angles
	px, py := pangles(cx, cy, r, []float64{30, 90, 150, 210, 270, 330})
	// make the deck markup
	fmt.Printf("<polygon xc=\"")
	end := len(px) - 1
	for i := 0; i < end; i++ {
		fmt.Printf("%.3f ", px[i])
	}
	fmt.Printf("%.3f\" yc=\"", px[end])
	for i := 0; i < end; i++ {
		fmt.Printf("%.3f ", py[i])
	}
	fmt.Printf("%.3f\" color=\"%s\"/>\n", py[end], color)
}

func polylines(cx, cy, r, lw float64, color string) {
	// construct a polygon with points at these angles
	px, py := pangles(cx, cy, r, []float64{30, 90, 150, 210, 270, 330})
	attr := fmt.Sprintf("sp=\"%g\" color=\"%s\"/>", lw, color)
	lx := len(px) - 1
	for i := 0; i < lx; i++ {
		fmt.Printf("<line xp1=\"%g\" yp1=\"%g\" xp2=\"%g\" yp2=\"%g\" %s\n", px[i], py[i], px[i+1], py[i+1], attr)
	}
	fmt.Printf("<line xp1=\"%g\" yp1=\"%g\" xp2=\"%g\" yp2=\"%g\" %s\n", px[0], py[0], px[lx], py[lx], attr)
}

// beginPage starts a page
func beginPage(bgcolor, textcolor string) {
	fmt.Printf("<slide bg=\"%s\" fg=\"%s\">\n", bgcolor, textcolor)
}

// endPage ends a page
func endPage() {
	ctext(50, 5, 1.2, "The area of a circle denotes state population: source U.S. Census", "sans", "gray")
	fmt.Println("</slide>")
}

// beginDoc starts the document
func beginDoc() {
	fmt.Printf("<deck>\n<canvas width=\"%d\" height=\"%d\"/>\n", width, height)
}

// endDoc ends the document
func endDoc() {
	fmt.Println("</deck>")
}

func main() {
	var opts options
	flag.IntVar(&width, "width", 1200, "canvas width")
	flag.IntVar(&height, "height", 900, "canvas height")
	flag.Float64Var(&opts.top, "top", 75, "top")
	flag.Float64Var(&opts.left, "left", 25, "left")
	flag.Float64Var(&opts.rowsize, "rowsize", float64(width)*0.006, "rowsize")
	flag.Float64Var(&opts.colsize, "colsize", float64(height)*0.006, "colsize")
	flag.StringVar(&opts.bgcolor, "bgcolor", "black", "background color")
	flag.StringVar(&opts.textcolor, "textcolor", "white", "textcolor")
	flag.StringVar(&opts.shape, "shape", "c", "shape for states:\n\"c\": circle,\n\"h\": hexagon,\n\"s\": square\n\"l\": line\n\"g\": geographic\n\"p\": plain text")
	flag.Parse()

	beginDoc()
	for _, f := range flag.Args() { // for every file, make a page
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
		process(opts, data, min, max, title)
		r.Close()
	}
	endDoc()
}
