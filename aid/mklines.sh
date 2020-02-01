#!/bin/sh
# makes a slide showing donations radiating from the map location
IFS="
"
if test $# -ne 1
then
	echo "specify a total-recipient data file" 1>&2
	exit 1
fi
filebase=$(echo "$1" | sed -e 's/ /-/g')
begin=$(grep "$1" recip-loc.d | awk -F\t '{print $2, $3}')
donation=$(grep "$1" td.d | awk -F\t '{printf "%.2f", $2}')

awk -F\t -v begin=$begin -v country="$1" -v donation=$donation '
BEGIN { 
	printf "\n\tslide\n\t\ttext \"%s\" dleft dtop dsize \"sans\" donorcolor\n", country
	printf "\t\ttext \"Donation Frequency\" 60 laby 1 \"sans\" \"gray\"\n"
	printf "\t\ttext \"Amount Received\"    78 laby 1 \"sans\" \"gray\"\n"
	printf "\t\timage mapimage 36 50 3200 1800 40\n"
	printf "\t\tcircle %s dcrad donorcolor 50\n", begin
	printf "\t\tctext \"$ %s Billion\" %s dtsize\n", donation, begin
}'

awk -F\t -v country="$1" '
	$1 == country {
		x=7
		for (f=2; f<=NF; f++) {
			printf "\t\timage \"%s\.png\"\t%d icony iconw iconh\n", $f, x
			x += 5
		}
	
	}' donor-icons.d

for i in $(awk -F\t '{print $1}' total-$filebase-recip.d)
do
	grep $i recip-loc.d | 
	awk -F\t -v begin=$begin -v country=$1 '{
		printf "\t\tline %s %s %s\tls lcolor lop\t// %s\n", begin, $2, $3, $1
		printf "\t\tcircle %s %s\t\tdotsize dotcolor dotop\n", $2, $3
	}'

done
awk -v fb=$filebase 'BEGIN { 
	printf "\t\tdchart opts  sorted-total-%s-recip.d\n", fb
	printf "\t\tdchart dopts sorted-%s-aid.d\n\teslide\n", fb
}'
