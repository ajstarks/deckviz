BEGIN {
	color["l1"] = "hsv(0,100,100)" # "darkred"
	color["l2"] = "hsv(0,75,100)" # "firebrick"
	color["l3"] = "hsv(0,60,100)" # "indianred"
	color["l4"] = "hsv(0,45,100)" # "hotpink"
	color["l5"] = "hsv(0,30,100)" # "palevioletred"
	color["l6"] = "hsv(0,15,100)"  # "pink"
	# color["l7"] = "silver"
}

$2 <= 400 && $2 > 200	{printf "georegion \"%s.json\" \"%s\"\n", $1, color["l1"]}
$2 <= 200 && $2 > 100	{printf "georegion \"%s.json\" \"%s\"\n", $1, color["l2"]}
$2 <= 100 && $2 > 50  	{printf "georegion \"%s.json\" \"%s\"\n", $1, color["l3"]}
$2 <= 50 && $2 > 30  	{printf "georegion \"%s.json\" \"%s\"\n", $1, color["l4"]}
$2 <= 30 && $2 > 20  	{printf "georegion \"%s.json\" \"%s\"\n", $1, color["l5"]}
$2 <= 20 && $2 > 0  	{printf "georegion \"%s.json\" \"%s\"\n", $1, color["l6"]}
# $2 <= 10 && $2 > 0   	{printf "georegion \"%s.json\" \"%s\"\n", $1, color["l7"]}
END {
	ly=25
	printf "square keyc1 %d keysize \"%s\"\n", ly+1, color["l1"]
	printf "text \"200-360\" keyc2 %d keysize \"sans\" \"%s\"\n", ly, color["l1"]
	ly-=5

	printf "square keyc1 %d keysize \"%s\"\n", ly+1, color["l2"]
	printf "text \"100-200\" keyc2 %d keysize \"sans\" \"%s\"\n", ly, color["l2"]
	ly-=5

	printf "square keyc1 %d keysize \"%s\"\n", ly+1, color["l3"]
	printf "text \"50-100\" keyc2 %d keysize \"sans\" \"%s\"\n", ly, color["l3"]
	ly-=5

	ly = 25
	printf "square keyc3 %d keysize \"%s\"\n", ly+1, color["l4"]
	printf "text \"30-50\" keyc4 %d keysize \"sans\" \"%s\"\n", ly, color["l4"]
	ly-=5

	printf "square keyc3 %d keysize \"%s\"\n", ly+1, color["l5"]
	printf "text \"20-30\" keyc4 %d keysize \"sans\" \"%s\"\n", ly, color["l5"]
	ly-=5

	printf "square keyc3 %d keysize \"%s\"\n", ly+1, color["l6"]
	printf "text \"20 and under\" keyc4 %d keysize \"sans\" \"%s\"\n", ly, color["l6"]
	ly-=5

	# printf "square keyc1 %d keysize \"%s\"\n", ly+1, color["l7"]
	# printf "text \"10 and under\" keyc1 %d keysize \"sans\" \"%s\"\n", lx+5, ly, color["l7"]
}