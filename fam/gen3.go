package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"
)

type Family struct {
	Name   string `xml:"name,attr"`
	Parent parent `xml:"parent"`
}

type parent struct {
	Husband  string  `xml:"husband,attr"`
	Wife     string  `xml:"wife,attr"`
	Married  string  `xml:"married,attr"`
	Image    string  `xml:"image,attr"`
	Children []child `xml:"child"`
}

type child struct {
	Generation int     `xml:"gen,attr"`
	Name       string  `xml:"name,attr"`
	Birth      string  `xml:"birth,attr"`
	Death      string  `xml:"death,attr"`
	Gender     string  `xml:"gender,attr"`
	Image      string  `xml:"image,attr"`
	Grands     []child `xml:"child"`
}

func readData(r io.Reader) (Family, error) {
	var data Family
	err := xml.NewDecoder(r).Decode(&data)
	return data, err
}

func beginDeck(w io.Writer) {
	fmt.Fprintln(w, "deck")
}

func beginSlide(w io.Writer) {
	fmt.Fprintln(w, "slide")
}

func endSlide(w io.Writer) {
	fmt.Fprintln(w, "eslide")
}

func endDeck(w io.Writer) {
	fmt.Fprintln(w, "edeck")
}

func text(w io.Writer, x, y, fs float64, s, font, color string) {
	fmt.Fprintf(w, "text \"%s\" %.2f %.2f %.2f %q %q\n", s, x, y, fs, font, color)
}

func etext(w io.Writer, x, y, fs float64, s, font, color string) {
	fmt.Fprintf(w, "etext \"%s\" %.2f %.2f %.2f %q %q\n", s, x, y, fs, font, color)
}

func ctext(w io.Writer, x, y, fs float64, s, font, color string) {
	fmt.Fprintf(w, "ctext \"%s\" %.2f %.2f %.2f %q %q\n", s, x, y, fs, font, color)
}

func rtext(w io.Writer, x, y, fs, angle float64, s, font, color string) {
	fmt.Fprintf(w, "rtext \"%s\" %.2f %.2f %.3f %.2f %q %q\n", s, x, y, fs, angle, font, color)
}

func brace(w io.Writer, x, y, height, bw, bh, strokeWidth float64, color string) {
	fmt.Fprintf(w, "lbrace %.2f %.2f %.2f %.2f %.2f %.2f\n", x, y, height, bw, bh, strokeWidth)
}

func circle(w io.Writer, x, y, r float64, color string, opacity float64) {
	fmt.Fprintf(w, "circle %.2f %.2f %.2f %q %.2f\n", x, y, r, color, opacity)
}

func line(w io.Writer, x1, y1, x2, y2, strokeWidth float64, color string, opacity float64) {
	fmt.Fprintf(w, "line %v %v %v %v %v %q %.2f\n", x1, y1, x2, y2, strokeWidth, color, opacity)
}

func image(w io.Writer, s string, x, y float64, width, height int, scale float64) {
	fmt.Fprintf(w, "image %q %v %v %v %v %v\n", s, x, y, width, height, scale)
}

const (
	canvasWidth       = 792
	canvasHeight      = 612
	midx              = 50.0
	midy              = 50.0
	familySize        = 5.0
	familyHeight      = 95.0
	parentX           = 3.0
	parentY           = 60.0
	parentBraceX      = 30.0
	parentSize        = familySize * 0.6
	parentSpacing     = 5.0
	parentBraceHeight = 95.0
	parentFooterY     = 5.0
	husbandFooterX    = midx - 25
	wifeFooterX       = midx + 25
	parentStroke      = 0.2
	footerX           = parentBraceX - 5
	footerSize        = 1.5
	childTop          = 92.0
	childSize         = 2.0
	childX            = parentBraceX + 15.0
	childDateX        = childX + 7
	childBraceX       = childX + 9
	childStroke       = parentStroke / 2
	childImageX       = parentBraceX + 4
	grandX            = childX + 10.5
	grandSize         = 1.2
	grandSpacing      = grandSize * 1.8
	ggrandSize        = grandSize * 0.75
	treeMinX          = 7.0
	treeMaxX          = 100 - treeMinX
	minFanAngle       = 30.0
	maxFanAngle       = 150.0
	fanRadius         = 4.0
	trunkStroke       = 0.2
	lineOpacity       = 25.0
	branchStroke      = trunkStroke / 2
	timeX             = 98.0
	timeY             = 2.0
	timeSize          = 0.6
	yearcolor         = "rgb(100,100,100)"
	grandcolor        = "maroon"
	ggcolor           = "steelblue"
	linecolor         = yearcolor
)

//vmap maps one interval to another
func vmap(value float64, low1 float64, high1 float64, low2 float64, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

func timestamp(w io.Writer) {
	now := time.Now().Format("Jan _2, 2006 3:04 PM")
	etext(w, timeX, timeY, timeSize, now, "sans", yearcolor)
}

func drawParent(w io.Writer, data Family) {
	if len(data.Parent.Image) > 0 {
		image(w, data.Parent.Image, 15, 90, 100, 0, 0)
	}
	py := parentY
	text(w, parentX, py, familySize, data.Name, "", "")
	py -= parentSpacing * 2
	text(w, parentX, py, parentSize, data.Parent.Husband, "", "")
	py -= parentSpacing
	text(w, parentX, py, parentSize, data.Parent.Wife, "", "")
	py -= parentSpacing
	text(w, parentX, py, parentSize, "(married "+data.Parent.Married+")", "sans", yearcolor)
	brace(w, parentBraceX, midy, parentBraceHeight, parentSize, parentSize, parentStroke, "")
}

func bminmax(children []child) (int, int) {
	min, max := 99999, 0
	for _, c := range children {
		birth, err := strconv.Atoi(c.Birth)
		if err != nil {
			return 0, 0
		}
		if birth < min {
			min = birth
		}
		if birth > max {
			max = birth
		}
	}
	return min, max
}

// polar to Cartesian coordinates, corrected for aspect ratio
func polar(cx, cy, r, theta, cw, ch float64) (float64, float64) {
	ry := r * (cw / ch)
	t := theta * (math.Pi / 180)
	return cx + (r * math.Cos(t)), cy + (ry * math.Sin(t))
}

func scale(w io.Writer, x, y float64, min, max, interval int) {
	fmin, fmax := float64(min), float64(max)
	line(w, treeMinX, y+3, treeMaxX, y+3, 0.1, yearcolor, 100)
	for v := min; v <= max; v += interval {
		label := strconv.Itoa(int(v))
		lx := vmap(float64(v), fmin, fmax, treeMinX, treeMaxX)
		ctext(w, lx, y, 1, label, "sans", yearcolor)
		line(w, lx, y+2, lx, y+3, 0.1, yearcolor, 100)
	}
}

func famtree(w io.Writer, data Family) {
	var gencolors = map[string]string{"m": "blue", "f": "#FF69B4"}
	ctext(w, midx, parentFooterY, familySize, data.Name, "sans", "")
	text(w, treeMinX, 100-parentFooterY, parentSize, data.Parent.Husband, "sans", "")
	etext(w, treeMaxX, 100-parentFooterY, parentSize, data.Parent.Wife, "sans", "")
	children := data.Parent.Children
	cy := 50.0
	cr := 1.5
	bmin, bmax := bminmax(children)
	posx := make([]float64, len(children))
	scale(w, treeMinX, cy-7, bmin, bmax, 1)
	var sep float64
	var fy float64
	sep = 10
	for i, c := range children {
		color := gencolors[c.Gender]
		birthdate, _ := strconv.Atoi(c.Birth)
		posx[i] = vmap(float64(birthdate), float64(bmin), float64(bmax), treeMinX, treeMaxX)
		cx := posx[i]
		if i > 0 {
			sep = posx[i] - posx[i-1]
		}
		circle(w, cx, cy, cr, color, 100)
		ctext(w, cx, cy-3, childSize*0.75, c.Name, "sans", "")
		fy = cy + (sep / 2) // trunkHeight
		if len(c.Grands) > 0 {
			div := (maxFanAngle - minFanAngle) / float64(len(c.Grands)-1)
			line(w, cx, cy, cx, fy, 1, color, lineOpacity)
			angle := maxFanAngle
			// for each grand child, render radiating from their parents
			for _, g := range c.Grands {
				gcolor := gencolors[g.Gender]
				ax, ay := polar(cx, fy, fanRadius, angle, canvasWidth, canvasHeight)
				line(w, cx, fy, ax, ay, 0.3, gcolor, 50)
				rtext(w, ax, ay, angle, grandSize*0.7, g.Name, "sans", grandcolor)

				if len(g.Grands) > 0 { // add great-grandchildren, following grandchildren, with a slight angle change
					ggadjust := float64(len(g.Name)) * (grandSize * 0.35)
					gx, gy := polar(cx, fy, fanRadius+ggadjust, angle, canvasWidth, canvasHeight)
					rtext(w, gx, gy, angle-minFanAngle, grandSize*0.5, ggstringName(g.Grands), "sans", ggcolor)
				}
				angle -= div
			}
		}
		circle(w, midx, 12, 2, gencolors["m"], lineOpacity)
		circle(w, midx, 12, 2, gencolors["f"], lineOpacity)
		line(w, midx, 12, cx, cy, cr/2, color, lineOpacity)
		counts(w, data)
	}
}

func ggstring(c []child) string {
	s := ""
	if len(c) < 1 {
		return s
	}
	l := len(c)
	for i := 0; i < l-1; i++ {
		s = s + c[i].Name + " (" + c[i].Birth + "), "
	}
	s = s + c[l-1].Name + " (" + c[l-1].Birth + ")"
	return s
}

func ggstringName(c []child) string {
	s := ""
	if len(c) < 1 {
		return s
	}
	l := len(c)
	for i := 0; i < l-1; i++ {
		s = s + c[i].Name + ", "
	}
	s = s + c[l-1].Name
	return s
}

func counts(w io.Writer, data Family) {
	nc := len(data.Parent.Children)
	ng := 0
	ngg := 0
	for _, c := range data.Parent.Children {
		for _, g := range c.Grands {
			ng++
			for i := 0; i < len(g.Grands); i++ { //  _, gg := range g.Grands {
				ngg++
			}
		}
	}
	tnc := "Children: " + strconv.Itoa(nc)
	tng := "Grand Children: " + strconv.Itoa(ng)
	tngg := "Great Grand Children: " + strconv.Itoa(ngg)
	etext(w, footerX, parentFooterY+6, footerSize, tnc, "sans", yearcolor)
	etext(w, footerX, parentFooterY+3, footerSize, tng, "sans", yearcolor)
	etext(w, footerX, parentFooterY, footerSize, tngg, "sans", yearcolor)
}

func drawChildren(w io.Writer, data Family) {
	children := data.Parent.Children
	childy := childTop
	//childSpacing := familyHeight / float64(len(children))
	for _, c := range children {
		etext(w, childX, childy, childSize, c.Name, "sans", "")
		etext(w, childDateX, childy, childSize, "("+c.Birth+")", "sans", yearcolor)

		lc := float64(len(c.Grands))
		gh := lc * grandSpacing
		gy := (childy + (gh / 2))

		if lc > 1 {
			brace(w, childBraceX, childy+(childSize*0.7), gh, 1, 1, childStroke, "")
		}
		if len(c.Image) > 0 {
			image(w, c.Image, childImageX, childy+childSize*1.5, 25, 0, 0)
		}
		for _, g := range c.Grands {
			text(w, grandX, gy, grandSize, g.Name, "sans", grandcolor)
			text(w, grandX+8, gy, grandSize, "("+g.Birth+")", "sans", yearcolor)
			if len(g.Grands) > 0 {
				text(w, grandX+13, gy, ggrandSize, ggstring(g.Grands), "sans", ggcolor)
			}
			gy -= grandSpacing
		}
		childy -= (gh + 4.5) // childSpacing
	}
	counts(w, data)
}

func draw(w io.Writer, data Family) {
	drawParent(w, data)
	drawChildren(w, data)
}

func process(w io.Writer, r io.Reader, style string) error {
	data, err := readData(r)
	if err != nil {
		return err
	}
	beginSlide(w)
	timestamp(w)
	switch style {
	case "tree":
		famtree(w, data)
	case "text":
		draw(w, data)
	}
	endSlide(w)
	return err
}

func main() {
	var style string
	flag.StringVar(&style, "style", "tree", "style (tree or text)")
	flag.Parse()

	beginDeck(os.Stdout)
	for _, f := range flag.Args() {
		r, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		err = process(os.Stdout, r, style)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
	}
	endDeck(os.Stdout)
}
