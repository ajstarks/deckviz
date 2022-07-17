#!/bin/sh
if test $# -ne 3
then
	echo "girl-text mario-text laura-text" 1>&2 
	exit 1
fi

cat <<!  | decksh |  pdfdeck -stdout -pagesize 1200x800 -sans impact - > db.pdf
deck
	slide "black" "white"
		mariox=63
		marioy=45
		laurax=85
		lauray=35
		girlx=29
		girly=20
		ts=6
		image "db.png" 50 50 100 0
		ctext "$1" girlx girly ts
		ctext "$2" mariox marioy ts
		ctext "$3" laurax lauray ts
	eslide
edeck
!