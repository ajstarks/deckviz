# Election Visualization with deck markup

## Input data

Input data file format (tab-separated), with columns in this order:
(1) state Abbreviation (for example NJ), 
(2) grid row, 
(3) grid column, 
(4) party which carried the state,
(5) population

A record that begis with the '#' is treated as a comment, and is used to build the header

	# Year Democratic-Candidate Republican-Candidate Third-Party-Candidate

The states are arranged in a grid with rows numbered 0-n west to east, and and columns numbered 0-n north
to south.

The party is designated as 'd' for the Democratic Party, and 'r' for Republican, 'i' for independent or third-party.
Other party affiliations may be added to the name of the candidate:

* name:f -- Federalist
* name:w -- Whig
* name:dr -- Democratic Republican

The population is from the US Census State Intercensal tables, using the data closest 
to the election year. For example, for election year 2020, the data estimated from 2019 is used. 
For the election of 2000, the official 2000 census data is used.

This repository has data files for two layout schemes: files named for year like ```2008.d``` use the layout shown on the Five-Thirty Eight
website. 

![538-layout](5.png)

Files named with the nyt prefix, like ```nyt-2008.d``` use the layout from the New York Times.

![nyt-layout](n.png)

## Running

here are the command line options:
```
   -bgcolor string
        background color (default "black")
  -colsize float
        colsize (default 5.4)
  -height int
        canvas height (default 900)
  -left float
        left (default 25)
  -rowsize float
        rowsize (default 7.2)
  -shape string
        shape for states: "c": circle, "h": hexagon, "s": square, "l": line, "g": geographic, "p": plain text (default "c")
  -textcolor string
        textcolor (default "white")
  -top float
        top (default 75)
  -width int
        canvas width (default 1200)
```

This command will make a PDF showing the elections from 1792-2024 using NYT layout style:

	go run elections.go -width 1200 -height 900 -left 15 nyt-????.d | pdfdeck -pagesize 1200,900 -sans PublicSans-Regular -stdout - > elections.pdf


## Data from US Census: 

* https://www.census.gov/search-results.html?q=State+Intercensal+Tables&page=1&stateGeo=none&searchtype=web&cssp=SERP&_charset_=UTF-8
* https://www.census.gov/data/tables/time-series/demo/popest/intercensal-2000-2010-state.html
* https://www.census.gov/data/tables/time-series/demo/popest/pre-1980-state.html
* https://www.census.gov/data/tables/time-series/demo/popest/pre-1980-state.html
