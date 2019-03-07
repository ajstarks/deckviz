#!/bin/sh
IFS="
"
if test $# -ne 1
then
	echo "specify a total-recipient data file" 1>&2
	exit 1
fi
begin=$(grep $1 recip-loc.d | awk '{print $2, $3}')
for i in $(awk '{print $1}' total-$1-recip.d)
do 
	grep $i recip-loc.d | awk -F\t -v begin=$begin '{printf "line %s  %s %s ls lcolor lop\t// %s\n", begin, $2, $3, $1}'
done
