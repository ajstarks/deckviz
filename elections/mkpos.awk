BEGIN {
	startx=10
	starty=75
	colsize=7
	rowsize=9
	cs=2
	ts=1.5
	printf "deck\nslide\n"
	partycolor["r"]="red"
	partycolor["d"]="blue"
	Pi=3.141592653589793238
}
{
	name=$1
	col=$2
	row=$3
	party=$4
	population=$5
	cs=vmap(area(population), 800, 7100, 1, colsize-1)
	x = startx + (row * colsize)
	y = starty - (col * rowsize)
	printf "ctext \"%s\" %g %g %g\n", name, x, y+(colsize*0.65) , ts
	printf "circle %g %g %g \"%s\"\n", x, y, cs, partycolor[party]
}

function vmap(value, low1, high1, low2, high2) {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

function area(v) {
	return sqrt((v / Pi)) * 2
}

END {
	printf "eslide\nedeck\n"
}