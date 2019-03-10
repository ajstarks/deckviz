#!/bin/sh
IFS="
"
for i in $(cat donor-countries)
do
	file=$(echo "$i" | sed -e 's/ /-/g')
	# echo "# $i" > sorted-$file-aid.d
	csvread 4 8 < aid.csv | 
	awk -v country="$i" -F\t '$1 == country { print $2 }' | dhist | sed 10q > sorted-$file-aid.d
done