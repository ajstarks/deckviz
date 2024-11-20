#!/bin/sh
(echo r ajs.jpg; echo blur 10)                  | giftsh > blur.jpg
(echo r ajs.jpg; echo brightness 20)            | giftsh > brightness.jpg
(echo r ajs.jpg; echo colorbalance 200 0 0)     | giftsh > colorbalance.jpg
(echo r ajs.jpg; echo colorize 200 100 100 )    | giftsh > colorize.jpg
(echo r ajs.jpg; echo colorspace l)             | giftsh > colorspace-l.jpg
(echo r ajs.jpg; echo colorspace s)             | giftsh > colorspace-s.jpg
(echo r ajs.jpg; echo contrast 20)              | giftsh > contrast.jpg
(echo r ajs.jpg; echo crop 0 0 200 200)         | giftsh > crop.jpg
(echo r ajs.jpg; echo cropsize 100 100)         | giftsh > cropsize.jpg
(echo r ajs.jpg; echo edge)                     | giftsh > edge.jpg
(echo r ajs.jpg; echo emboss)                   | giftsh > emboss.jpg
(echo r ajs.jpg; echo fliph)                    | giftsh > fliph.jpg
(echo r ajs.jpg; echo flipv)                    | giftsh > flipv.jpg
(echo r ajs.jpg; echo gamma 2)                  | giftsh > gamma.jpg
(echo r ajs.jpg; echo gray)                     | giftsh > gray.jpg
(echo r ajs.jpg; echo hue 75)                   | giftsh > hue.jpg
(echo r ajs.jpg; echo invert)                   | giftsh > invert.jpg
(echo r ajs.jpg; echo max 3)                    | giftsh > max.jpg
(echo r ajs.jpg; echo mean 5)                   | giftsh > mean.jpg
(echo r ajs.jpg; echo median 5)                 | giftsh > median.jpg
(echo r ajs.jpg; echo min 5)                    | giftsh > min.jpg
(echo r ajs.jpg; echo opacity 60)               | giftsh > opacity.jpg
(echo r ajs.jpg; echo pixelate 50)              | giftsh > pixelate.jpg
(echo r ajs.jpg; echo resizefill 512 512)       | giftsh > resizefill.jpg
(echo r ajs.jpg; echo resizefit 512 512)        | giftsh > resizefit.jpg
(echo r ajs.jpg; echo resize 200 200)           | giftsh > resize.jpg
(echo r ajs.jpg; echo rotate 45)                | giftsh > rotate.jpg
(echo r ajs.jpg; echo saturation 200)           | giftsh > saturation.jpg
(echo r ajs.jpg; echo sepia 100)                | giftsh > sepia.jpg
(echo r ajs.jpg; echo sigmoid 0.5 0)            | giftsh > sigmoid.jpg
(echo r ajs.jpg; echo sobel)                    | giftsh > sobel.jpg
(echo r ajs.jpg; echo threshold 60)             | giftsh > threshold.jpg
(echo r ajs.jpg; echo transpose)                | giftsh > transpose.jpg
(echo r ajs.jpg; echo transverse)               | giftsh > transverse.jpg
(echo r ajs.jpg; echo unsharp 1 1 0.05)         | giftsh > unsharp.jpg
