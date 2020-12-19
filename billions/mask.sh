#!/bin/sh
for i in $*
do
	b=$(basename $i .jpg)
	echo $b
	convert $i -gravity Center \( -size 300x300 xc:Black -fill White -draw 'circle 150 150 150 1'  -alpha Copy \) -compose CopyOpacity -composite -trim $b-mask.png
done