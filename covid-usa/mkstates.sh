#!/bin/sh
# create a file for each state or region
# the format of the input: each record is the number of cases per county, columns are dates
if test $# -ne 1
then
	echo "specify an input file" 2>&1
	exit 1
fi
input="$1"
csvread -plain=f  6 $(seq 50 428) < $input > state-cases.csv		# select the columns state-name, date....
head -1 state-cases.csv > head.csv									# get the column header
csvread 0 < state-cases.csv | sort -u > states						# get the list of states/regions
IFS="
"
# for each state/region, make a the state CSV along with a summation TSV file ready for charting
for i in $(cat states)
do
	name=$(echo "$i" | sed -e 's/ /-/g')
	echo $name
	cat head.csv > "$name".csv
	grep "$i" state-cases.csv >> "$name".csv
	echo "#" "$i" > $name.d
	./sumstates.sh "$name".csv >> "$name".d
done
