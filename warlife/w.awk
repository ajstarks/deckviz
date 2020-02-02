# How much of your life has the US been at war?
# https://twitter.com/ianbremmer/status/1215389030013132800
# Generates decksh markup
BEGIN {
	if (pie == "") {
		pie="t"
	}
	linew=4

	x=10
	y=90
	d=4
	r=d/2

	tx=30
	nx=90
	ny=10

	xincr=d*2
	yincr=d*2
	op=60
	ts=2.1
	bgcolor="black"
	tcolor="white"
	ccolor="white"
	dcolor="maroon"
	title="How much of your life the U.S. has been at war"
	ref="https://twitter.com/ianbremmer/status/1215389030013132800"

	printf "deck\nslide \"%s\" \"%s\"\n", bgcolor, tcolor
	printf "text \"%s\" %d %d %.2f\n", title, tx, ny, ts
	printf "textblock \"Source: Washington Post: %s\" %d %d 50 1 \"sans\" \"%s\" 100 \"%s\"\n", ref, tx, ny-3, tcolor, ref
	if (pie == "t") {
		printf "ctext \"birthdate\" %d 13 1 \"sans\" \"%s\"\n", nx, tcolor
		printf "circle %d %d %.2f \"%s\"\n", nx, ny, d*1.1, ccolor
		printf "arc %d %d %.2f %.2f 0 270 %.2f \"%s\" %g\n", nx, ny, r, r, r, dcolor, op
		printf "ctext \"% war\" %d %d  1 \"serif\" \"%s\"\n", nx, ny, bgcolor
	} else {
		printf "text \"birthdate\" %d 13 1 \"sans\" \"%s\"\n", nx, tcolor
		printf "text \"war %\" %d %d  1 \"serif\" \"%s\"\n", nx, ny+2, tcolor
		printf "hline %.2f %.2f %.2f 1 \"gray\"\n", nx, ny, linew
		printf "hline %.2f %.2f %.2f 1 \"%s\"\n", nx, ny, linew*(0.75), dcolor
	}
}  
$0 ~! /^#/ && NF == 2  {
	if (x > 95) {
		x=10
		y-=yincr
	}
	if (pie == "t") {
		printf "ctext \"%s\" %.2f %.2f 1\n", $1, x, y+3
		printf "ctext \"%.1f%%\" %.2f %.2f 1.2 \"serif\" \"%s\"\n", $2, x, y-0.5, bgcolor
		printf "circle %.2f %.2f %.2f \"%s\" \n", x, y, d*1.1, ccolor
		if ($2 == 100.0) {
			printf "circle %.2f %.2f %.2f \"%s\" %g\n", x, y, d, dcolor, op
		} else {
			printf "arc %.2f %.2f %.2f %.2f 0 %.2f %.2f \"%s\" %g\n", x, y, r, r, (($2/100)*360), r, dcolor, op
		}
	} else {
		printf "text \"%s\" %.2f %.2f 1\n", $1, x, y+3
		printf "text \"%.1f%%\" %.2f %.2f 1.2 \"serif\" \"%s\"\n", $2, x, y+1.75, tcolor
		printf "hline %.2f %.2f %.2f 1 \"gray\"\n", x, y, linew
		printf "hline %.2f %.2f %.2f 1 \"%s\"\n", x, y, linew*($2/100), dcolor
	}

	x+=xincr
}  
END {
	printf "eslide\nedeck\n"
}
