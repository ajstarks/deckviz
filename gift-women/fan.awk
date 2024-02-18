#!/bin/sh
# make decksh commands for a upper fan chart, given data in the form of 
# Label\tnumber\tcolor
# usage:
# fanchart.awk cx cy size file
if test $# -lt 4
then
	echo "fanchart.awk cx cy size startangle file" 1>&2
	exit 1
fi
cx=$1
cy=$2
size=$3
startangle=$4
shift 4
awk \
-v cx=$cx -v cy=$cy -v size=$size -v startangle=$startangle '
BEGIN {
	FS="\t"
	span=startangle-(180-startangle)
}
{
	sum += $2
	label[NR] = $1
	data[NR] = $2
	colors[NR] = $3
	outlables="no"
}
END {
	start=startangle
	for (i=1; i <= NR; i++) {
		pct = data[i]/sum
		m = pct * span
		a1 = start - m
		a2 = start
		labelangle =  a2-((a2-a1)/2)
		
		printf "edge=polar %g %g %g %g\n", cx, cy, size, labelangle

		if (labelangle > 90) {
			ttype="ctext"
		} else {
			ttype="text"
		}
		
		if (outables == "yes") {
			r=size+2
			tsize=1.5
			printf "lab=polar %g %g %g %g\n", cx, cy, r, labelangle
			printf "line edge_x edge_y lab_x lab_y 0.1 \"%s\"\n", colors[i]
			printf "%s \"%s (%.2f%%)\" lab_x lab_y tsize \"sans\" \"%s\"\n", ttype, label[i], pct*100, tsize, "black"
		} else {

			if (pct*100 > 30) {
				tsize=2.25
				r=size*0.7
				printf "lab=polar %g %g %g %g\n", cx, cy, r, labelangle
				
			} else {
				r = size+2
				tsize=1.5
				printf "lab=polar %g %g %g %g\n", cx, cy, r, labelangle
				printf "line edge_x edge_y lab_x lab_y 0.1 \"%s\"\n", colors[i]
				
			}
			
			printf "%s \"%.2f%%\" lab_x lab_y %.2f \"sans\" \"%s\"\n", ttype, pct*100, tsize, "black"
		}
		printf "arc %g %g %g %g %g %g %g \"%s\" wop\t// %s %.1f%%\n", cx, cy, size, size, a1, a2,  size, colors[i], label[i], pct*100
		start = a1
	}
}
' $*