BEGIN {
	printf "deck\nslide\n"

	n = 0
	left = 0
	right = 200
	width = 4
	height = 5
	top = 50
	bottom = top - height
	for (x=left; x < right; x+=width) {
		if (n % 2 == 0) {
			y = top
		} else {
			y = bottom
		}
		xp[n] = x
		yp[n] = y
		n++
	}

	printf "polygon \""
	for (i=0; i < n; i++) {
		printf "%g ", xp[i]
	}
	printf "%g\" \"", left
	for (i=0; i < n; i++) {
		printf "%g ", yp[i]
	}
	printf "%g\"\n", bottom

	printf "eslide\nedeck\n"
}
