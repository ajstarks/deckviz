# deckviz -- Visualizations created with [decksh](https://github.com/ajstarks/deck/blob/master/cmd/decksh/README.md)

### Contents

* 2020election -- 2020 Election
* aapl -- Apple 2018 Results
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
* complex-table -- charts and tables
* covid-usa -- Covid data
* cyano -- #RecreationThursday -- Will Creech grid
* diptych -- American Diptych
* domestica -- Output from the Domestica course "Data Visualization and Information Design"
* elections -- elections through the years
* fam -- family trees
* famfountain -- family fountains
* festival -- Stage from the 1969 Harlem Cultural Festival
* fire -- German wildfires
* flag -- Flags
* functable -- Mathematical Functions Table
* go -- go built from arcs and lines
* gomod -- go build flows
* hlito -- Hlito for #RecreationThursday
* hsv -- HSV color model
* human-freedom -- Remix of Human Freedom Index visualization
* iam -- "I AM A MAN" poster
* juneteenth -- Juneteenth flag
* lamp -- Bauhaus lamp
* lift -- Lift every voice and sing
* litany -- Du Bois's "A Litany of Atlanta"
* lynch -- "A Man Was Lynched" poster
* messiah -- A passage from Handel's Messiah
* middle -- the Middle Passage
* mlb -- Major League Baseball Brackets
* modelt -- Manufacturing and Marketing the Model T - 1908-1916
* multigen -- Multi-generational households
* nba -- NBA standings
* netflix -- Netflix viewing frequency
* odita -- #recreationthursday from Odili Donald Odita
* oscars -- 2022 Academy Awards
* pic -- diagrams from pic
* prison -- Comparing the rate of incarceration vs. US presidential terms
* qs --  Quadratspirale schwarz, 1952 (#RecreationThursday)
* reqheader -- Request headers diagram
* slave -- African Slave Routes
* speed -- Comparison of Indy 500 speed data
* stockgrid -- stocks using seasonal palette
* stolen -- The Stolen Series
* swd -- Examples from "Storytelling with Data"
* t3 -- Dieter Rams T3 Radio (1958)
* tailscale -- TailScale style diagrams
* thomas -- Iris, Tulips, Jonquils, and Crocuses, 1969, Alma Thomas
* toni -- Toni Morrison quote from Charlie Rose Interview, 1993
* top-netflix -- Netflix viewing data
* tuskegee -- The Tuskegee Airmen
* ur -- quote from Jenkins' "Underground Railroad"
* voyager -- the Voyager spacecraft
* warlife -- How much of your life has the U.S. been at war
* wealth-gap -- US Wealth Gap
* weather -- Weather trends

Each directory has the '.dsh' scripts, associated data files, and rendered output in PDF and PNG formats.

## Running the examples

### get the deck tools

	go install github.com/ajstarks/deck/cmd/decksh@latest	# get decksh
	go install github.com/ajstarks/deck/cmd/pdfdeck@latest  # get pdfdeck to render PDFs
	go install github.com/ajstarks/deck/cmd/pngdeck@latest	# get pngdeck to render PNGs
	go install github.com/ajstarks/deck/cmd/dchart@latest   # get the dchart command
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
	

	


