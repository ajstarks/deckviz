# How much of your life has the US been at war?
# https://twitter.com/ianbremmer/status/1215389030013132800
# Generates decksh markup (using grid)
BEGIN {
	bgcolor="black"
	tcolor="white"
	ccolor="white"
	dcolor="maroon"
	d=4
	r=d/2
	op=40
}  
$0 ~! /^#/ && NF == 2  {
	if (layer == "year") {
		printf "ctext \"%s\" x y 1\n", $1
	}
	if (layer == "pct") {
		printf "ctext \"%.1f%%\" x y 1.2 \"serif\" \"%s\"\n", $2, bgcolor
	}
	if (layer == "circle") {
		printf "circle x y  %.2f \"%s\" \n", d*1.1, ccolor
	}
	if (layer == "arc") {
		if ($2 == 100.0) {
			printf "circle x y  %.2f \"%s\" %g\n", d, dcolor, op
		} else {
			printf "arc x y  %.2f %.2f 0 %.2f %.2f \"%s\" %g\n", r, r, (($2/100)*360), r, dcolor, op
		}
	}
}