#!/bin/bash
if test $# -ne 2
then
	echo orig-name month-name
	exit 1
fi
oname=$1
mname=$2
sed -n 1,22p $oname-2019.csv > $mname-2019-01.csv
sed -n 23,41p $oname-2019.csv > $mname-2019-02.csv
sed -n 24,62p $oname-2019.csv > $mname-2019-03.csv
sed -n 63,83p $oname-2019.csv > $mname-2019-04.csv
sed -n 84,105p $oname-2019.csv > $mname-2019-05.csv
sed -n 106,125p $oname-2019.csv > $mname-2019-06.csv
sed -n 107,147p $oname-2019.csv > $mname-2019-07.csv
sed -n 148,169p $oname-2019.csv > $mname-2019-08.csv
sed -n 170,189p $oname-2019.csv > $mname-2019-09.csv
sed -n 190,212p $oname-2019.csv > $mname-2019-10.csv
sed -n 213,232p $oname-2019.csv > $mname-2019-11.csv
sed -n 233,252p $oname-2019.csv > $mname-2019-12.csv

ed < puthead.ed $mname-2019-02.csv
ed < puthead.ed $mname-2019-03.csv
ed < puthead.ed $mname-2019-04.csv
ed < puthead.ed $mname-2019-05.csv
ed < puthead.ed $mname-2019-06.csv
ed < puthead.ed $mname-2019-07.csv
ed < puthead.ed $mname-2019-08.csv
ed < puthead.ed $mname-2019-09.csv
ed < puthead.ed $mname-2019-10.csv
ed < puthead.ed $mname-2019-11.csv
ed < puthead.ed $mname-2019-12.csv

awk -F, 'NR > 1 {print $5}' $oname-2019.csv | sort -nr  > $mname.limit
echo "-min=" $(head -1 $mname.limit) "-max=" $(tail -1 $mname.limit)
