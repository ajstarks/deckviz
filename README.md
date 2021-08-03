# deckviz -- Visualizations created with [decksh](https://github.com/ajstarks/deck/blob/master/cmd/decksh/README.md)

### Contents

* 2020election -- 2020 Election
* absense -- Theaster Gates from "Black Art: Absence of Light"
* aid -- International aid as defined in the #SWDchallenge
* aapl -- Apple financial results
* astrobang -- Houston Astros 2017 sign-stealing
* b17 -- Bomber casuality rates
* babynames -- Baby names through the decades
* billions -- top 10 US billionaires
* bp -- Blood Pressure and Heart Rate
* c19 -- COVID-19
* ceo-stock -- The relationship of CEO tenure to stock price
* covid-usa -- Covid data
* complex-table -- charts and tables
* diptych -- American Diptych
* elections -- elections through the years
* fam -- family trees
* festival -- Stage from the 1969 Harlem Cultural Festival
* fire -- German wildfires
* flag -- Flags
* go -- go built from arcs and lines
* gomod -- go build flows
* hlito -- Hlito for #RecreationThursday
* iam -- "I AM A MAN" poster
* juneteenth -- Juneteenth flag
* gomod -- Go module flows
* lynch -- "A Man eas Lynched" poster
* middle -- the Middle Passage
* modelt -- Manufacturing and Marketing the Model T - 1908-1916
* multigen -- Multi-generational households
* nba -- NBA standings
* netflix -- Netflix viewing frequency
* prison -- Comparing the rate of incarceration vs. US presidential terms
* qs --  Quadratspirale schwarz, 1952 (#RecreationThursday)
* reqheader -- Request headers diagram
* speed -- Comparison of Indy 500 speed data
* stockgrid -- stocks using seasonal palette
* swd -- Examples from "Storytelling with Data"
* tailscale -- TailScale style diagrams
* toni -- Toni Morrison quote from Charlie Rose Interview, 1993
* voyager -- the Voyager spacecraft
* warlife -- How much of your life has the U.S. been at war
* weather -- Weather trends

Each directory has the '.dsh' scripts, associated data files, and rendered output in PDF and PNG formats.

## Running the examples

### get the deck tools

	go get github.com/ajstarks/deck/cmd/decksh		# get decksh
	go get github.com/ajstarks/deck/cmd/pdfdeck		# get pdfdeck to render PDFs
	go get github.com/ajstarks/deck/cmd/pngdeck		# get pngdeck to render PNGs
	go get github.com/ajstarks/deck/cmd/dchart      # get the dchart command
	# assume the DECKFONTS environment variable is set to $HOME/deckfonts
	cd $HOME
	git clone https://github.com/ajstarks/deckfonts
	
### clone the repo
	
	git clone https://github.com/ajstarks/deckviz	# clone this repo
	Cloning into 'deckviz'...
	...
	cd deckviz
	
### move the a directory, execute and render

	cd fire
	decksh fire.dsh > fire.xml
	pdfdeck fire.xml  # produces fire.pdf
	pngdeck fire.xml  # produces fire-00001.png

Note that some directories require more scripting, and will have a script called ```mkdeck```

	cd astrobang
	./mkdeck
	

	


