package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// C19 is the json data returned by https://coronavirus.projectpage.app/.json?period=0
type C19 struct {
	Dates     []string `json:"dates"`
	Deaths    []int    `json:"deaths"`
	Confirmed []int    `json:"confirmed"`
	AllTimeDeaths int `json:"alltimeDeaths"`
	allTimeConfirmed int `json:"allTimeConfirmed"`
}

func main() {
	var data C19
	if err := json.NewDecoder(os.Stdin).Decode(&data); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Println("date,deaths,confirmed")
	for i := 0; i < len(data.Dates); i++ {
		fmt.Printf("\"%s\",%d,%d\n", data.Dates[i], data.Deaths[i], data.Confirmed[i])
	}
}
