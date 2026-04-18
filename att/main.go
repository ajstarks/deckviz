package main

import (
	"fmt"
	"os"

	"github.com/ajstarks/gpdf"
)

const (
	title        = "APOLLO 11 TRANSLUNAR/TRANSEARTH TRAJECTORY CHART"
	cx           = 50
	cy           = 60
	cw           = 612
	ch           = 792
	ttop         = 97.0
	tts          = 2.5
	ats          = 1.0
	earthSize    = 7 // 10
	moonSize     = (earthSize * 0.27) / 2
	outerRadius  = 40
	outerRadius1 = outerRadius + 1    // 41
	outerRadius2 = outerRadius1 + 0.5 // 41.5
	outerRadius3 = outerRadius2 + 0.5 // 42.0
	outerRadius4 = outerRadius3 + 0.5 // 42.5
	outerRadius5 = outerRadius4 + 0.5 // 43
	outerRadius6 = outerRadius5 + 0.5 // 43.5
	textcolor    = "black"
	distcolor    = "gray"
	teicolor     = "green"
	tlicolor     = "blue"
	lineSpacing  = 1.25
)

type tabrow struct {
	colsize float64
	ls      float64
	text    string
	color   string
}

type moon struct {
	size       float64
	radius     float64
	angle      float64
	note       string
	color      string
	orbitcolor string
}

type locnote struct {
	note      string
	color     string
	align     string
	textsize  float64
	radius    float64
	angle     float64
	rotation  float64
	showpoint bool
}

var locations = []locnote{
	{note: "Launch", radius: 3.2, angle: 287, color: tlicolor},
	{note: "Earth Orbit", radius: 4, angle: 90, color: tlicolor, align: "c"},
	{note: "Translunar insertion", radius: 4, angle: 220, color: tlicolor, showpoint: true, align: "e"},
	{note: "TRANSLUNAR COAST", radius: 12, angle: 40, rotation: 70, color: tlicolor},
	{note: "Lunar orbit insertion", radius: 39.5, angle: 60, color: tlicolor},
	{note: "Lunar landing (Jul 20)", radius: 38.5, angle: 73, color: tlicolor, showpoint: true, align: "e"},
	{note: "Lunar liftoff (Jul 21)", radius: 41, angle: 85, color: tlicolor},
	{note: "Transearth Injection (Jul 22)", radius: 41.5, angle: 96, showpoint: true, color: teicolor, align: "e"},
	{note: "TRANSEARTH COAST", radius: 11, angle: 100, rotation: 80, color: teicolor},
	{note: "Entry Interface", radius: 4, angle: 160, color: teicolor, showpoint: true, align: "e"},
	{note: "Splashdown", radius: 3, angle: 200, color: teicolor, showpoint: true, align: "e"},
	{note: "12 hrs", radius: 13, angle: 30, color: tlicolor},
	{note: "1 day", radius: 22, angle: 47, color: tlicolor},
	{note: "1 day, 12 hrs", radius: 27, angle: 52, color: tlicolor},
	{note: "2 days", radius: 33, angle: 56, color: tlicolor},
	{note: "2 days, 12 hrs", radius: 37, angle: 58, color: tlicolor},
	{note: "6 days", radius: 35, angle: 94, color: teicolor, align: "e"},
	{note: "6 days, 12 hrs", radius: 28, angle: 95, color: teicolor, align: "e"},
	{note: "7 days", radius: 22, angle: 97, color: teicolor, align: "e"},
	{note: "7 days, 12 hrs", radius: 12, angle: 108, color: teicolor, align: "e"},
	{note: "8 days", radius: 5, angle: 145, color: teicolor, align: "e"},
}

var missionTable = [][]tabrow{
	{
		{colsize: 5, ls: lineSpacing, text: "Launch Dates"},
		{text: "Jul 16"},
		{text: "Jul 18"},
		{text: "Jul 21"},
	},
	{
		{colsize: 5, ls: lineSpacing, text: ""},
		{text: "Aug 14"},
		{text: "Aug 16"},
		{text: "Aug 20"},
	},
	{
		{colsize: 6, ls: lineSpacing, text: "Site No."},
		{text: "2"},
		{text: "3"},
		{text: "5"},
	},
	{
		{colsize: 6, ls: lineSpacing, text: "Site Coordinates"},
		{text: "0°42.8'N"},
		{text: "0°21.2'N"},
		{text: "1°40.7'N"},
	},
	{
		{colsize: 2, ls: lineSpacing, text: ""},
		{text: "23°42.5'E"},
		{text: "1°18.0'W"},
		{text: "41°54.0'W"},
	},
}

var eventTable = [][]tabrow{
	{
		{colsize: 15, ls: lineSpacing, text: "Event", color: textcolor},
		{text: "Launch", color: tlicolor},
		{text: "Earth Orbit", color: tlicolor},
		{text: "Translunar Injection", color: tlicolor},
		{text: "Translunar Coast", color: tlicolor},
		{text: "Lunar Orbit Insertion", color: tlicolor},
		{text: "Lunar Landing", color: tlicolor},
		{text: "Lunar Lift-off", color: teicolor},
		{text: "Transearth Injection", color: teicolor},
		{text: "Transearth Coast", color: teicolor},
		{text: "Entry Interface", color: teicolor},
		{text: "Splashdown", color: teicolor},
	},
	{
		{ls: lineSpacing, text: "Time (EDT)", color: textcolor},
		{text: "July 16 9:32", color: tlicolor},
		{text: "July 16 9:43", color: tlicolor},
		{text: "July 16 12:16", color: tlicolor},
		{text: "July 16-19", color: tlicolor},
		{text: "July 19 13:21", color: tlicolor},
		{text: "July 20 16:17", color: tlicolor},
		{text: "July 21 13:54", color: teicolor},
		{text: "July 22 00:55", color: teicolor},
		{text: "July 22-24", color: teicolor},
		{text: "July 24 12:39", color: teicolor},
		{text: "July 24 12:50", color: teicolor},
	},
}

var flightTable = [][]tabrow{
	{
		{colsize: 22, ls: lineSpacing, text: "Event"},
		{text: "Launch to Translunar Injection"},
		{text: "Translunar Injection to Lunar Orbit Insertion"},
		{text: "Lunar Orbit Insertion to Transearth Injection"},
		{text: "Transearth Injection to Entry Interface"},
		{text: "Entry Interface to Pacific Ocean Touchdown"},
		{text: ""},
		{text: "Total Flight Time"},
	},
	{
		{colsize: 2, ls: lineSpacing, text: "hrs"},
		{text: "02"},
		{text: "73"},
		{text: "59"},
		{text: "59"},
		{text: "00"},
		{text: ""},
		{text: "195"},
	},
	{
		{colsize: 2, ls: lineSpacing, text: "min"},
		{text: "44"},
		{text: "16"},
		{text: "24"},
		{text: "37"},
		{text: "14"},
		{text: ""},
		{text: "15"},
	},
}

var moons = []moon{
	{size: moonSize, radius: outerRadius6, angle: 350, note: "July 13", color: textcolor},
	{size: moonSize, radius: outerRadius6, angle: 346, note: "Aug 9", color: textcolor},
	{size: moonSize, radius: outerRadius5, angle: 335, note: "Aug 8", color: textcolor},
	{size: moonSize, radius: outerRadius4, angle: 325, note: "Aug 7", color: textcolor},
	{size: moonSize, radius: outerRadius3, angle: 313, note: "Aug 6", color: textcolor},
	{size: moonSize, radius: outerRadius2, angle: 301, note: "Aug 5", color: textcolor},
	{size: moonSize, radius: outerRadius1, angle: 289, note: "Aug 4", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 276, note: "Aug 3", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 263, note: "Aug 2", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 248, note: "Aug 1", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 234, note: "July 31", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 220, note: "July 30", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 205, note: "July 29", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 190, note: "July 28", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 173, note: "July 27", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 157, note: "July 26", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 144, note: "July 25", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 128, note: "July 24", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 114, note: "July 23", color: textcolor},
	{size: moonSize, radius: outerRadius, angle: 97, note: "", color: textcolor, orbitcolor: teicolor},
	{size: moonSize, radius: outerRadius, angle: 85, note: "", color: textcolor, orbitcolor: tlicolor},
	{size: moonSize, radius: outerRadius, angle: 73, note: "", color: textcolor, orbitcolor: tlicolor},
	{size: moonSize, radius: outerRadius1, angle: 60, note: "July 19", color: textcolor, orbitcolor: tlicolor},
	{size: moonSize, radius: outerRadius2, angle: 47, note: "July 18", color: textcolor},
	{size: moonSize, radius: outerRadius3, angle: 36, note: "July 17", color: textcolor},
	{size: moonSize, radius: outerRadius4, angle: 25, note: "Jul 16 (Launch)", color: textcolor},
	{size: moonSize, radius: outerRadius5, angle: 13, note: "July 15", color: textcolor},
	{size: moonSize, radius: outerRadius6, angle: 2, note: "Jul 14 (New Moon)", color: textcolor},
}

// textblockfile reads a file and makes a textblock from its contents
func textblockfile(canvas *gpdf.Canvas, filename string, x, y, w, size, ls float64, color string) {
	s, err := os.ReadFile(filename)
	if err != nil {
		canvas.TextWrapStrict(x, y, w, size, ls, "--- unable to open "+filename+" ----", "red")
		return
	}
	canvas.TextWrapStrict(x, y, w, size, ls, string(s), color)
}

// dist show distance in nautical miles as concentric circles
func dist(canvas *gpdf.Canvas, begin, end, incr float64) {
	canvas.ImageScaleName(cx, cy, earthSize, 0, 100, "globe0.png")
	canvas.Circle(cx, cy, earthSize*0.6, tlicolor, 10)
	nm := 50.0
	for d := begin; d <= end; d += incr {
		rad := d / 2
		canvas.Circle(cx, cy, rad, distcolor, 7)
		lx, ly := canvas.PolarDegrees(cx, cy, rad, 270)
		label := fmt.Sprintf("%v", nm) + ",000"
		canvas.CText(lx, ly, ats, label, textcolor)
		if d == begin {
			canvas.CText(lx, ly-1.2, ats, "nautical miles", textcolor)
		}
		nm += 50.0
	}
}

// showtables lays out tables
func showtable(canvas *gpdf.Canvas, x, y, size float64, t [][]tabrow) {
	xp := x
	for i := range t {
		yp := y
		ls := t[i][0].ls
		cs := t[i][0].colsize
		for j := range t[i] {
			canvas.Text(xp, yp, size, t[i][j].text, t[i][j].color)
			yp -= (size * ls)
		}
		xp += cs
	}
}

// trajectory plots a trajectory curve
func trajectory(canvas *gpdf.Canvas) {
	tlix, tliy := canvas.PolarDegrees(cx, cy, 4, 220)   // translunar injection
	teix, teiy := canvas.PolarDegrees(cx, cy, 41.5, 96) // transearth injection
	loix, loiy := canvas.PolarDegrees(cx, cy, 39.5, 60) // lunar orbit insertion
	tei2x, tei2y := canvas.PolarDegrees(cx, cy, 20, 97) // transearth injection (2)
	splx, sply := canvas.PolarDegrees(cx, cy, 3, 200)   // splashdown

	// control points
	c1x, c1y := canvas.PolarDegrees(cx, cy, 14, 300)
	c2x, c2y := canvas.PolarDegrees(cx, cy, 41, 90)
	c3x, c3y := canvas.PolarDegrees(cx, cy, 6, 200)

	cw := 0.1
	cop := 100.0
	canvas.Curve(tlix, tliy, c1x, c1y, loix, loiy, cw, tlicolor, cop)
	canvas.Curve(teix, teiy, c2x, c2y, tei2x, tei2y, cw, teicolor, cop)
	canvas.Curve(tei2x, tei2y, c3x, c3y, splx, sply, cw, teicolor, cop)
}

// showlocations annotates locations using polar coordinates
func showlocations(canvas *gpdf.Canvas, locations []locnote) {
	for _, l := range locations {
		x, y := canvas.PolarDegrees(cx, cy, l.radius, l.angle)
		ts := 1.0
		if l.textsize > 0 {
			ts = l.textsize
		}
		switch {
		case l.rotation > 0:
			canvas.RText(x, y, ts, l.rotation, l.note, l.color)
		case l.align == "e":
			canvas.EText(x, y, ts, l.note, l.color)
		case l.align == "c":
			canvas.CText(x, y, ts, l.note, l.color)
		default:
			canvas.Text(x, y, ts, l.note, l.color)
		}
		if l.showpoint {
			canvas.Circle(x, y, ts/3, l.color, 50)
		}
	}
}

// showmoons plots the position of the moon
func showmoons(canvas *gpdf.Canvas, moonpos []moon) {
	for _, m := range moonpos {
		mx, my := canvas.PolarDegrees(cx, cy, m.radius, m.angle)
		if len(m.orbitcolor) > 0 {
			canvas.Circle(mx, my, m.size+(m.size/3), m.orbitcolor, 20)
		}
		canvas.Circle(mx, my, m.size, m.color, 20)
		canvas.Arc(mx, my, m.size, m.size, 100, 280, m.size, m.color, 20)
		lx, ly := canvas.PolarDegrees(mx, my, m.size+(m.size*.7), 90)
		canvas.CText(lx, ly, 1, m.note, m.color)
	}
}

// att combines the various components of a Apollo translunar/transearth chart
func att(canvas *gpdf.Canvas) {
	c1 := 3.0
	c2 := c1 + 37
	c3 := c2 + 30
	top := 25.0
	als := lineSpacing
	aw := 30.0
	canvas.CText(50, ttop, tts, title, textcolor)
	dist(canvas, 20, 80, 20)
	showmoons(canvas, moons)
	showlocations(canvas, locations)
	trajectory(canvas)
	textblockfile(canvas, "desc1.txt", c1, top, aw, ats, als, textcolor)
	showtable(canvas, c1, top-8, ats, missionTable)
	textblockfile(canvas, "desc2.txt", c1, top-14, aw, ats, als, textcolor)
	showtable(canvas, c2, top, ats*1.2, eventTable)
	canvas.TextWrapStrict(c3, top, aw, ats, als, "Approximate flight time for a July 16, 1969 launch date is as follows:", textcolor)
	showtable(canvas, c3, top-3, ats, flightTable)
}

func main() {
	// set up the canvas
	canvas, err := gpdf.SetupCanvas(cw, ch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	// load the font
	font, err := canvas.LoadFontFile("sans.ttf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	canvas.SetFont(font)
	// build the visualization, write the output
	att(canvas)
	err = canvas.Creator.WriteToFile("att.pdf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
}
