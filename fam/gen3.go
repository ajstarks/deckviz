package main

import (
	"encoding/xml"
	"fmt"
	"io"
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

func brace(w io.Writer, x, y, height, bw, bh, strokeWidth float64, color string) {
	fmt.Fprintf(w, "lbrace %.2f %.2f %.2f %.2f %.2f %.2f\n", x, y, height, bw, bh, strokeWidth)
}

//func line(w io.Writer, x1, y1, x2, y2, strokeWidth float64, color string) {
//	fmt.Fprintf(w, "line %v %v %v %v %v %q\n", x1, y1, x2, y2, strokeWidth, color)
//}

func image(w io.Writer, s string, x, y float64, width, height int, scale float64) {
	fmt.Fprintf(w, "image %q %v %v %v %v %v\n", s, x, y, width, height, scale)
}

const (
	midy              = 50.0
	familySize        = 5.0
	familyHeight      = 95.0
	parentX           = 3.0
	parentY           = 60.0
	parentBraceX      = 30.0
	parentSize        = familySize * 0.6
	parentSpacing     = 5.0
	parentBraceHeight = 95.0
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
	yearcolor         = "rgb(100,100,100)"
	grandcolor        = "maroon"
	ggcolor           = "steelblue"
	linecolor         = yearcolor
)

func timestamp(w io.Writer) {
	now := time.Now().Format("Jan _2, 2006 3:04 PM")
	etext(w, 95, 2, 1, now, "sans", yearcolor)
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

func ggstring(c []child) string {
	s := ""
	l := len(c)
	for i := 0; i < l-1; i++ {
		s = s + c[i].Name + " (" + c[i].Birth + "), "
	}
	s = s + c[l-1].Name + " (" + c[l-1].Birth + ")"
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
	etext(w, footerX, 15, footerSize, tnc, "sans", "")
	etext(w, footerX, 12, footerSize, tng, "sans", "")
	etext(w, footerX, 9, footerSize, tngg, "sans", "")
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
	timestamp(w)
	drawParent(w, data)
	drawChildren(w, data)
}

func process(w io.Writer, r io.Reader) error {
	data, err := readData(r)
	if err != nil {
		return err
	}
	beginSlide(w)
	draw(w, data)
	endSlide(w)
	return err
}

func main() {
	beginDeck(os.Stdout)
	for _, f := range os.Args[1:] {
		r, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		err = process(os.Stdout, r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
	}
	endDeck(os.Stdout)
}
