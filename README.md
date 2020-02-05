# deckviz -- Visualizations created with [decksh](https://github.com/ajstarks/deck/blob/master/cmd/decksh/README.md)

### Contents

* aid -- International aid as defined in the #SWDchallenge
* aapl -- Apple financial results
* astrobang -- Houston Astros 2017 sign-stealing
* babynames -- Baby names through the decades
* bp -- Blood Pressure and Heart Rate
* ceo-stock -- The relationship of CEO tenure to stock price
* fire -- German wildfires
* gomod -- Go module flows
* modelt -- Manufacturing and Marketing the Model T - 1908-1916
* multigen -- Multi-generational households
* nba -- NBA standings
* prison -- Comparing the rate of incarceration vs. US presidential terms
* speed -- Comparison of Indy 500 speed data
* swd -- Examples from "Storytelling with Data"
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
	

	


