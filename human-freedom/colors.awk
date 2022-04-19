#
# reads the human freedom data (ISO-CODE, score, region, direction)
# makes a decksh grid file
#
BEGIN {
	# map regions to colors
	color["Caucasus & Central Asia"]="#0b8ca6"
	color["East Asia"]="#0b631e"
	color["Eastern Europe"]="#eac91a"
	color["Latin America & the Caribbean"]="#a60018"
	color["Middle East & North Africa"]="#f0e9ce"
	color["North America"]="#e89313"
	color["Oceania"]="#a63e3e"
	color["South Asia"]="#8d8475"
	color["Sub-Saharan Africa"]="#143f83"
	color["Western Europe"]="#222222"

	# map direction code to decksh function
	dir["u"]="up"
	dir["d"]="down"
	dir["s"]="same"
	dir[""]="same"
}
{
	if (color[$3] == "#f0e9ce" || color[$3] == "#eac91a") {
		textcolor="black"
	} else {
		textcolor="white"
	}
	printf "%s \"%s\" x y %g %.2f \"%s\" \"%s\"\n", dir[$4], $1, size, $2, color[$3], textcolor
}
