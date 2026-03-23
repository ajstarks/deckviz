package main

import (
	"fmt"
	"math"
)

func circle(x, y, size float64, color string, op float64) {
	//fmt.Printf("circle %.3f %.3f %.3f %q %v\n", x, y, size, color, op)
	fmt.Printf("<ellipse xp=\"%.3f\" yp=\"%.3f\" wp=\"%.3f\" hr=\"100\" color=%q opacity=\"%.3f\"/>\n", x, y, size, color, op)
}

func dottedline(x1, y1, x2, y2, size float64, count int, color string, op float64) {
	n := float64(count)
	// Undefined slope (vertical lines)
	if x1 == x2 {
		interval := math.Abs(y2-y1) / n
		if y1 < y2 {
			y := y1
			for range count {
				circle(x1, y, size, color, op)
				y += interval
			}
		} else {
			y := y2
			for range count {
				circle(x1, y, size, color, op)
				y += interval
			}
		}
		return
	}

	x := x1
	if x2 < x1 {
		x = x2
	}
	interval := math.Abs(x2-x1) / n
	m := (y2 - y1) / (x2 - x1) // slope
	for range count {
		y := m*(x-x1) + y1
		circle(x, y, size, color, op)
		x += interval
	}
}
func polar(x, y, r, theta float64) (float64, float64) {
	return (r * math.Cos(theta)) + x, (r * math.Sin(theta)) + y
}

func main() {
	r := 50.0
	nd := 10
	ds := 2.5
	fmt.Println("<deck>\n<slide>")
	for t := 0.0; t < 360; t += 9 {
		a := t * (math.Pi / 180)
		x, y := polar(50, 50, r, a)
		dottedline(50, 50, x, y, ds, nd, "blue", 50)
	}
	dottedline(50, 50, 50, 100, ds*2, nd, "red", 50)
	dottedline(50, 50, 50, 0, ds*2, nd, "red", 50)
	fmt.Println("</slide></deck>")

}
