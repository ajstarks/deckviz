#!/bin/bash
IFS="
"
for i in $(cat donor-countries)
do
	csvread -headskip 4 5 6 < aid.csv | awk -v donor="$i" -F\t '$1 == donor {sum+=$3} END {print donor "\t" sum}'
done > total-donors.d

for i in $(cat recipient-countries)
do
	csvread -headskip 4 5 6 < aid.csv | awk -v recip="$i" -F\t '$2 == recip {sum+=$3} END {print recip "\t" sum}'
done > total-recipients.d

