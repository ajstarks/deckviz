#!/bin/sh
if test $# -ne 1
then
	echo "specify an image" 1>&2
	exit 1
fi
input="$1" ; type=`echo $input|awk -F. '{print $2}'`
(echo r $input; echo blur 10)                  | giftsh > blur.$type
(echo r $input; echo brightness 20)            | giftsh > brightness.$type
(echo r $input; echo colorbalance 200 0 0)     | giftsh > colorbalance.$type
(echo r $input; echo colorize 200 100 100 )    | giftsh > colorize.$type
(echo r $input; echo colorspace l)             | giftsh > colorspace-l.$type
(echo r $input; echo colorspace s)             | giftsh > colorspace-s.$type
(echo r $input; echo contrast 20)              | giftsh > contrast.$type
(echo r $input; echo crop 0 0 200 200)         | giftsh > crop.$type
(echo r $input; echo cropsize 100 100)         | giftsh > cropsize.$type
(echo r $input; echo edge)                     | giftsh > edge.$type
(echo r $input; echo emboss)                   | giftsh > emboss.$type
(echo r $input; echo fliph)                    | giftsh > fliph.$type
(echo r $input; echo flipv)                    | giftsh > flipv.$type
(echo r $input; echo gamma 2)                  | giftsh > gamma.$type
(echo r $input; echo gray)                     | giftsh > gray.$type
(echo r $input; echo hue 75)                   | giftsh > hue.$type
(echo r $input; echo invert)                   | giftsh > invert.$type
(echo r $input; echo max 3)                    | giftsh > max.$type
(echo r $input; echo mean 5)                   | giftsh > mean.$type
(echo r $input; echo median 5)                 | giftsh > median.$type
(echo r $input; echo min 5)                    | giftsh > min.$type
(echo r $input; echo opacity 60)               | giftsh > opacity.$type
(echo r $input; echo pixelate 50)              | giftsh > pixelate.$type
(echo r $input; echo resizefill 512 512)       | giftsh > resizefill.$type
(echo r $input; echo resizefit 512 512)        | giftsh > resizefit.$type
(echo r $input; echo resize 200 200)           | giftsh > resize.$type
(echo r $input; echo rotate 45)                | giftsh > rotate.$type
(echo r $input; echo saturation 200)           | giftsh > saturation.$type
(echo r $input; echo sepia 100)                | giftsh > sepia.$type
(echo r $input; echo sigmoid 0.5 0)            | giftsh > sigmoid.$type
(echo r $input; echo sobel)                    | giftsh > sobel.$type
(echo r $input; echo threshold 60)             | giftsh > threshold.$type
(echo r $input; echo transpose)                | giftsh > transpose.$type
(echo r $input; echo transverse)               | giftsh > transverse.$type
(echo r $input; echo unsharp 1 1 0.05)         | giftsh > unsharp.$type
