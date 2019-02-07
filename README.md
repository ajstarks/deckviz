# deckviz -- Visualizations created with [decksh](https://github.com/ajstarks/deck/blob/master/cmd/decksh/README.md)

### Contents

* aapl -- Apple financial results
* babyname -- Baby names through the decades
* ceo-stock -- The relationship of CEO tenure to stock price
* gomod -- Go module flows
* fire -- German wildfires
* multigen -- Multi-generational households
* nba -- NBA standings
* prison -- Comparing the rate of incarceration vs. US presidential terms
* speed -- Comparison of Indy 500 speed data
* swd -- Examples from "Storytelling with Data"
* weather -- Weather trends

Each directory has the '.dsh' scripts, associated data files, and rendered output in PDF and PNG formats.

## Running the examples

### get the deck tools

	$ go get github.com/ajstarks/deck/cmd/decksh		# get decksh
	$ go get github.com/ajstarks/deck/cmd/pdfdeck		# get pdfdeck to render PDFs
	$ go get github.com/ajstarks/deck/cmd/pngdeck		# get pngdeck to render PNGs
	
### clone the repo
	
	$ git clone https://github.com/ajstarks/deckviz	# clone this repo
	Cloning into 'deckviz'...
	...
	$ cd deckviz
	
### move the a directory, execute and render, for example


	$ cd fire
	$ decksh fire.dsh > fire.xml
	$ pdfdeck fire.xml  # produces fire.pdf
	$ pngdeck fire.xml  # produces fire-00001.png
	

	


