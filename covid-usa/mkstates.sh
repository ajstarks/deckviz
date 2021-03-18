#!/bin/bash
# create a file for each state or region
# the format of the input: each record is the number of cases per county, columns are dates
# input data is from 
# https://github.com/CSSEGISandData/COVID-19/blob/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_US.csv
# Direct download: https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_US.csv
if test $# -ne 1
then
	echo "specify an input file" 2>&1
	exit 1
fi
input="$1"
nf=$(awk -F, 'NR==1{print NF}' $input)						# compute the number of columns == dates
nf=$(expr $nf - 1)											# substract the label column
cols=$(seq 50 $nf)											# enumerate the columns starting at 3/1/20 (column 50)
csvread -plain=f 6 $cols < $input > state-cases.csv			# select the columns state-name, date....
head -1 state-cases.csv > head.csv							# get the column header
csvread 0 < state-cases.csv | sort -u > states				# get the list of states/regions
# set IFS because state names have embedded spaces
IFS="
"
# for each state/region, make a the state CSV along with a summation TSV file ready for charting
for i in $(cat states)
do
	name=$(echo "$i" | sed -e 's/ /-/g')					# convert spaces to dashes for file names
	echo $name												# report progress
	cat head.csv > "$name".csv								# begin with the the header
	grep "$i" state-cases.csv >> "$name".csv				# select the state
	echo "#" "$i" > $name.d									# add the comment header for the TSV file
	awk -F, -f sumstates.awk "$name".csv >> "$name".d		# sum each state
done
