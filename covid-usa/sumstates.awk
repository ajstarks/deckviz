# first record is the column headers
NR==1 {
	for (i=2; i <= NF; i++) {
		head[i-1]=$i
	}
}
# sum each column
NR > 1 {
	for (i=2; i <= NF; i++) {
		sum[i-1]+=$i
	}
}
# print the sums with the columns
END {
	for (i=1; i <= NF-1; i++) {
		print head[i] "\t" sum[i]
	}
}
