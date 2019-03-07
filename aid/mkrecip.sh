#!/bin/sh
# for every donor country, collect the recipients, total the aid for each recipient
# donor-countries is a list of donor countries
# <country-name>-recip is a list of donors for each country
# total-<country-name>-recip.d is the total amount of aid for each country
IFS="
"
for d in $(cat donor-countries)
do
	cf=$(echo "$d" | sed -e 's/ /-/g')
	csvread 4 5 6 < aid.csv | awk -F\t -v donor="$d" '$1 == donor { print $2 }' | sort -u > "$cf-recip"
	for r in $(cat $cf-recip)
	do
		csvread 4 5 6 < aid.csv | 
		awk -F\t -v recip="$r" -v donor="$d" '$1 == donor && $2 == recip {sum += $3} END {print recip "\t" sum}'
	done > total-$cf-recip.d
done
