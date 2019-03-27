#!/bin/sh
awk -F$'\t' '
	BEGIN {
		print "deck\n\tslide\n\t\timage \"ww.png\" 50 50 1920 1080"
	}
	$0 !~ /^#/ {
		name=$1
		x=$2
		y=$3
		printf "\t\tcircle %.2f %.2f 0.25 \"red\" 50\n", x, y
		printf "\t\tctext \"%s\" %.2f %.2f 0.25\n", name, x, y
	}
	END {
		print "\teslide\nedeck"
	}' $*
