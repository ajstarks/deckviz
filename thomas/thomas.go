// Iris, Tulips, Jonquils, and Crocuses, 1969, Alma Thomas
// Acrylic on canvas
// 60 × 50 in
// 152.4 × 127 cm
// RecreationThursday 2021-10-28
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type brushStack struct {
	n         int     // number of items
	intensity int     // how many chips/item
	height    float64 // height in percent
	color     string  // the chip color
	opacity   float64 // chip  opacity
}

type brushStacks []brushStack

const (
	red1       = "#9b524e"
	red2       = "#8a2a1b"
	blue1      = "#204271"
	blue2      = "#68a8b2"
	blue3      = "#5c8bb9"
	blue4      = "#9dcbd3"
	pink       = "#c5abb0"
	yellow1    = "#e7d367"
	yellow2    = "#cba842"
	yellow3    = "#d3d9b3"
	orange     = "#b36735"
	violet1    = "#4b4b7a"
	violet2    = "#904561"
	violet3    = "#908eba"
	green1     = "#376769"
	green2     = "#c2cfbb"
	green3     = "#759899"
	bluegreen  = "#62a5ab"
	defop      = 100.0
	nleft      = 5.0
	nbottom    = 2.0
	nright     = 95.0
	nsize      = 1.0
	ellipsefmt = `<ellipse xp="%.3f" yp="%.3f" wp="%.3f" hp="%.3f" color=%q opacity="%.3f"/>`
	polyfmt    = `<polygon xc="%.3f %.3f %.3f %.3f" yc="%.3f %.3f %.3f %.3f" color=%q opacity="%.3f"/>`
	textfmt    = `<text xp="%.3f" yp="%.3f" sp="%.3f" align=%q>%s</text>`
	notefmt    = `brush=%s n=%d w=%.2f h=%.2f d=%.2f x=%g y=%g o=%.2f`
)

var alldata = []brushStacks{
	{
		{n: 3, color: green2, opacity: defop},
		{n: 33, color: blue1, opacity: defop},
	},
	{
		{n: 8, color: green2, opacity: defop},
		{n: 12, color: green1, opacity: defop},
		{n: 16, color: blue3, opacity: defop},
	},
	{
		{n: 10, color: green2, opacity: defop},
		{n: 26, color: blue1, opacity: defop},
	},
	{
		{n: 8, color: green2, opacity: defop},
		{n: 14, color: blue1, opacity: defop},
		{n: 7, color: blue3, opacity: defop},
		{n: 7, color: blue1, opacity: defop},
	},
	{
		{n: 9, color: green2, opacity: defop},
		{n: 27, color: blue1, opacity: defop},
	},
	{
		{n: 7, color: green2, opacity: defop},
		{n: 29, color: violet1, opacity: defop},
	},
	{
		{n: 21, color: yellow1, opacity: defop},
		{n: 15, color: blue1, opacity: defop},
	},
	{
		{n: 18, color: yellow1, opacity: defop},
		{n: 18, color: green2, opacity: defop},
	},
	{
		{n: 36, color: yellow2, opacity: defop},
	},
	{
		{n: 36, color: orange, opacity: defop},
	},
	{
		{n: 36, color: red2, opacity: defop},
	},
	{
		{n: 36, color: violet1, opacity: defop},
	},
	{
		{n: 3, color: yellow1, opacity: defop},
		{n: 5, color: green1, opacity: defop},
		{n: 7, color: violet1, opacity: defop},
		{n: 10, color: blue1, opacity: defop},
		{n: 11, color: violet1, opacity: defop},
	},
	{
		{n: 4, color: yellow1, opacity: defop},
		{n: 6, color: green1, opacity: defop},
		{n: 21, color: blue1, opacity: defop},
		{n: 5, color: green1, opacity: defop},
	},
	{
		{n: 32, color: blue1, opacity: defop},
		{n: 4, color: blue2, opacity: defop},
	},
	{
		{n: 33, color: blue1, opacity: defop},
		{n: 3, color: blue2, opacity: defop},
	},
	{
		{n: 9, color: red1, opacity: defop},
		{n: 27, color: blue1, opacity: defop},
	},
	{
		{n: 10, color: yellow1, opacity: defop},
		{n: 26, color: blue1, opacity: defop},
	},
	{
		{n: 4, color: yellow1, opacity: defop},
		{n: 32, color: blue1, opacity: defop},
	},
	{
		{n: 5, color: yellow1, opacity: defop},
		{n: 31, color: green1, opacity: defop},
	},
	{
		{n: 36, color: yellow1, opacity: defop},
	},
	{
		{n: 36, color: red1, opacity: defop},
	},
	{
		{n: 36, color: red1, opacity: defop},
	},
	{
		{n: 36, color: blue2, opacity: defop},
	},
	{
		{n: 36, color: red1, opacity: defop},
	},
	{
		{n: 36, color: red1, opacity: defop},
	},
	{
		{n: 36, color: orange, opacity: defop},
	},
	{
		{n: 36, color: green1, opacity: defop},
	},
	{
		{n: 36, color: blue1, opacity: defop},
	},
	{
		{n: 36, color: blue1, opacity: defop},
	},
	{
		{n: 36, color: blue1, opacity: defop},
	},
	{
		{n: 28, color: violet2, opacity: defop},
		{n: 8, color: blue1, opacity: defop},
	},
	{
		{n: 11, color: orange, opacity: defop},
		{n: 25, color: blue1, opacity: defop},
	},
	{
		{n: 14, color: green2, opacity: defop},
		{n: 22, color: blue1, opacity: defop},
	},
	{
		{n: 36, color: green2, opacity: defop},
	},
	{
		{n: 36, color: yellow1, opacity: defop},
	},
	{
		{n: 36, color: red1, opacity: defop},
	},
	{
		{n: 36, color: pink, opacity: defop},
	},
	{
		{n: 36, color: blue1, opacity: defop},
	},
	{
		{n: 36, color: blue1, opacity: defop},
	},
	{
		{n: 36, color: blue4, opacity: defop},
	},
	{
		{n: 36, color: yellow3, opacity: defop},
	},
	{
		{n: 36, color: blue1, opacity: defop},
	},
	{
		{n: 6, color: red1, opacity: defop},
		{n: 30, color: blue1, opacity: defop},
	},
	{
		{n: 15, color: red1, opacity: defop},
		{n: 21, color: blue1, opacity: defop},
	},
	{
		{n: 17, color: red1, opacity: defop},
		{n: 19, color: violet3, opacity: defop},
	},
	{
		{n: 18, color: red1, opacity: defop},
		{n: 18, color: bluegreen, opacity: defop},
	},
	{
		{n: 15, color: pink, opacity: defop},
		{n: 21, color: blue1, opacity: defop},
	},
	{
		{n: 16, color: pink, opacity: defop},
		{n: 20, color: green3, opacity: defop},
	},
}

// chip makes a four-sided polygon bounded to the rectangle
// defined by (x, y) at its center with dimensions (w,h).
// wd defines the depth inside the left and right sides,
// hd defines the depth inside the top and bottom
// The effect is a brush Stack when the coordinates of the polygon are
// defined randomly within these constraints.
func chip(x, y, w, h, wd, hd float64, color string, opacity float64) {
	xp := make([]float64, 4)
	yp := make([]float64, 4)

	l := x - (w / 2)
	r := x + (w / 2)
	t := y + (h / 2)
	b := y - (h / 2)

	xp[0] = l + (w * rand.Float64())
	yp[0] = t - (hd * rand.Float64())

	xp[1] = r - (wd * rand.Float64())
	yp[1] = b + (h * rand.Float64())

	xp[2] = l + (w * rand.Float64())
	yp[2] = b + (hd * rand.Float64())

	xp[3] = l + (wd * rand.Float64())
	yp[3] = b + (h * rand.Float64())

	fmt.Printf(polyfmt, xp[0], xp[1], xp[2], xp[3], yp[0], yp[1], yp[2], yp[3], color, opacity)
}

// blob makes n number of ellipses bounded by the rectangle
// centered at (x,y) with dimensions (w,h)
// the width and height of the ellipses are determined as
// some fraction of (w,h) determined by (wh, hd)
func blob(x, y, w, h, wd, hd float64, n int, color string, opacity float64) {
	l := x - (w / 2)
	b := y - (h / 2)
	for i := 0; i < n; i++ {
		xp, yp := l+(w*rand.Float64()), b+(h*rand.Float64())
		ew, eh := w*wd, h*hd
		fmt.Printf(ellipsefmt, xp, yp, ew, eh, color, opacity)
	}
}

// ctower makes a column of brush strokes, using the "chip" brush stroke
func ctower(data brushStacks, x, y, w, h, wd, hd float64, nb int, opacity float64) {
	yp := y
	for i := 0; i < len(data); i++ {
		d := data[i]
		for j := 0; j < d.n; j++ {
			for n := 0; n < nb; n++ {
				chip(x, yp, w, h, wd, hd, d.color, opacity)
			}
			yp += h
		}
	}
}

// btower makes a colum of brush strokes using the "blob" brush stroke
func btower(data brushStacks, x, y, w, h, wd, hd float64, nb int, opacity float64) {
	yp := y
	for i := 0; i < len(data); i++ {
		d := data[i]
		for j := 0; j < d.n; j++ {
			blob(x, yp, w, h, wd, hd, nb, d.color, opacity)
			yp += h
		}
	}
}

// allctower makes a series of columns beginning at (x,y) using the chip brush stroke
// the data that defines the columns is in "data"
func allctower(data []brushStacks, x, y, w, h, wd, hd float64, nb int, op float64) {
	xp := x
	for _, f := range data {
		ctower(f, xp, y, w, h, wd, hd, nb, op)
		xp += w
	}
}

// allbtower makes a series of columns beginning at (x,y) using the blob brush stroke
// the data that defines the columns is in "data"
func allbtower(data []brushStacks, x, y, w, h, wd, hd float64, nb int, op float64) {
	xp := x
	for _, f := range data {
		btower(f, xp, y, w, h, wd, hd, nb, op)
		xp += w
	}
}

// grid makes a 2-D grid of brush strokes
func grid(x1, x2, y1, y2, w, h, wd, hd float64, nb int, palette []string) {
	for x := x1; x <= x2; x += w {
		for y := y1; y <= y2; y += h {
			for n := 0; n < nb; n++ {
				c := rand.Intn(len(palette))
				blob(x, y, w, h, wd, hd, 1, palette[c], 10)
			}
		}
	}
}

func main() {
	var w, h, dfactor, sx, sy, opacity float64
	var nc int
	var bgcolor, brushtype string
	var note bool
	flag.Float64Var(&w, "w", 1.5, "brush width")
	flag.Float64Var(&h, "h", 2.4, "brush height")
	flag.Float64Var(&dfactor, "d", 0.05, "depth factor")
	flag.Float64Var(&sx, "x", 15, "start x")
	flag.Float64Var(&sy, "y", 10, "start y")
	flag.Float64Var(&opacity, "o", defop, "opacity")
	flag.IntVar(&nc, "n", 2, "number if chips")
	flag.StringVar(&bgcolor, "bg", "white", "background color")
	flag.StringVar(&brushtype, "brush", "c", "brush type (c for chip, b for blob")
	flag.BoolVar(&note, "note", false, "show notes")
	flag.Parse()

	rand.Seed(int64(time.Now().Unix()))
	fmt.Printf("<deck>\n<slide bg=\"%s\">\n", bgcolor)
	wd := w * dfactor
	hd := h * dfactor

	if note {
		ntext := fmt.Sprintf(notefmt, brushtype, nc, w, h, dfactor, sx, sy, opacity)
		fmt.Printf(textfmt, nleft, nbottom, nsize, "l", ntext)
		fmt.Printf(textfmt, nright, nbottom, nsize, "e", time.Now().Format(time.RFC3339))
	}
	switch brushtype {
	case "c":
		allctower(alldata, sx, sy, w, h, wd, hd, nc, opacity)
	case "b":
		allbtower(alldata, sx, sy, w, h, wd, hd, nc, opacity)
	default:
		allctower(alldata, sx, sy, w, h, wd, hd, nc, opacity)
	}
	fmt.Println("</slide></deck>")

	// fmt.Printf("slide \"%s\"\n", bgcolor)
	// blob(50, 50, 10, 15, 1, 1, 10, red1, 60)
	// fmt.Println("eslide")

	// fmt.Printf("slide \"%s\"\n", bgcolor)
	// grid(20, 80, 20, 80, w, h, 1, 1, 10, []string{"red", "green", "blue", "pink", "violet", "yellow", "orange"})
	// fmt.Println("eslide")
}
