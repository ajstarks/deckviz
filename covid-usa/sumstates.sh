#!/bin/sh
# make a tsv file summing the columns for each date
awk -F, -f sumstates.awk "$@"