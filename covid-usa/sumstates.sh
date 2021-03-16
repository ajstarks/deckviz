#!/bin/sh
# make a tsv file summing the columns for each date
awk -F, '
# first record is the column headers
NR==1 {
	for(i=2;i<NF;i++) {
		head[i]=$i
	}
}
# sum each column
NR > 1 {
	for (i=1; i <= NF; i++) {
		sum[i]+=$i
	}
}
# print the sums with the columns
END {
	for (i=1; i <= NF; i++) {
		print head[i-1] "\t" sum[i]
	}
}' "$@"
