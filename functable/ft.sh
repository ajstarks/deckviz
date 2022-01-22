#!/bin/sh
./mfunc -f sine   > sine.d
./mfunc -f cosine > cos.d
./mfunc -f sincos > sincos.d
./mfunc -f exp    > exp.d
./mfunc -f sqrt   > sqrt.d
./mfunc -f gamma -x 0.1,4,0.05 > gamma.d

decksh <<!

deck
slide

ctext "Math Functions" 50 90 5

left=7
top=70
cw=25
ch=10
rowheight=40
gutter=5

opts="-bar=f -vol -linewidth=0.1 -line -volop=10 -xlabel=10 -val=f -textsize=1.5"

t=top
b=t-ch

l=left
r=l+cw
dchart opts -left l -bottom b -right r -top t -color red sine.d

l=r+gutter
r=l+cw
dchart opts -left l -bottom b -right r -top t -color green cos.d

l=r+gutter
r=l+cw
dchart opts -left l -bottom b -right r -top t -color blue sincos.d

t-=rowheight
b=t-ch

l=left
r=l+cw
dchart opts -left l -bottom b -right r -top t -color orange exp.d

l=r+gutter
r=l+cw
dchart opts -left l -bottom b -right r -top t -color seagreen sqrt.d

l=r+gutter
r=l+cw
dchart opts -left l -bottom b -right r -top t -color steelblue gamma.d

eslide
edeck

!



