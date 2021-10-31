// thomas: a generative art piece derived from
// Iris, Tulips, Jonquils, and Crocuses, 1969, Alma Thomas
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// randpoly makes a four-sided polygon bounded to the rectangle
// defined by (x, y) at its center with dimensions (w,h).
// wd and hd define the depth from the left/right, top/bottom respectively
// The effect is a brush stroke when the coordinates of the polygon are
// defined randomly within these constraints.
func randpoly(x, y, w, h, wd, hd float64, color string, opacity float64) {
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

	fmt.Printf("polygon \"%.2f %.2f %.2f %.2f\" \"%.2f %.2f %.2f %.2f\" \"%s\" %.2f\n",
		xp[0], xp[1], xp[2], xp[3], yp[0], yp[1], yp[2], yp[3], color, opacity)
}

func tower(x, y, w, h, wd, hd float64, n, nb int, color string, opacity float64) {
	yp := y
	for i := 0; i < n; i++ {
		for j := 0; j < nb; j++ {
			randpoly(x, yp, w, h, wd, hd, color, opacity)
		}
		yp += h
	}
}

func grid(x1, x2, y1, y2, w, h, wd, hd float64, nb int, palette []string) {
	rand.Seed(int64(time.Now().Unix()))
	for x := x1; x <= x2; x += w {
		for y := y1; y <= y2; y += h {
			for n := 0; n < nb; n++ {
				c := rand.Intn(len(palette))
				randpoly(x, y, w, h, wd, hd, palette[c], 10)
			}
		}
	}
}

func main() {
	bgcolor := "white"
	fmt.Printf("deck\nslide \"%s\"\n", bgcolor)
	w := 4.0
	h := w * 1.6
	wd := w * 0.05
	hd := h * 0.05
	grid(35, 70, 10, 90, w, h, wd, hd, 10, []string{"red", "green", "blue", "pink", "violet", "yellow", "orange"})

	tower(10, 10, 1, 3, wd, hd, 10, 3, "red", 100)
	tower(10, 50, 1, 3, wd, hd, 10, 3, "blue", 100)
	fmt.Println("eslide\nedeck")
}
