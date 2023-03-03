package main

import (
	// flag for system flags
	"flag"
	// globals (for initializing globals
	"github.com/vibovenkat123/review-gpt/pkg/globals"
	// to call the API
	"github.com/vibovenkat123/review-gpt/pkg/request"
)

// flag variables
// the git diff
var input string

// the model
var model string

// the maxiumum amount of tokens received
var maxtokens int

// the temperature
var temperature float64

// the top_p
var top_p float64

// the frequence penalty
var frequence float64

// the presence penalty
var presence float64

// if its pretty
var pretty bool

// Amount of times to call
var bestof int

func init() {
	// setup the flag by looping through the flags array and setting them
	for _, inputFlag := range globals.InputFlag.Names {
		flag.StringVar(&input, inputFlag, "", globals.InputFlag.Help)
	}
	for _, prettyFlag := range globals.PrettyFlag.Names {
		flag.BoolVar(&pretty, prettyFlag, true, globals.PrettyFlag.Help)
	}
	for _, modelFlag := range globals.ModelFlag.Names {
		flag.StringVar(&model, modelFlag, "text-davinci-003", globals.ModelFlag.Help)
	}
	for _, maxTokensFlag := range globals.MaxTokenFlag.Names {
		flag.IntVar(&maxtokens, maxTokensFlag, 500, globals.MaxTokenFlag.Help)
	}
	for _, temperatureFlag := range globals.TemperatureFlag.Names {
		flag.Float64Var(&temperature, temperatureFlag, 0.2, globals.TemperatureFlag.Help)
	}
	for _, toppFlag := range globals.ToppFlag.Names {
		flag.Float64Var(&top_p, toppFlag, 1, globals.ToppFlag.Help)
	}
	for _, frequenceFlag := range globals.FrequenceFlag.Names {
		flag.Float64Var(&frequence, frequenceFlag, 1.2, globals.FrequenceFlag.Help)
	}
	for _, presenceFlag := range globals.PresenceFlag.Names {
		flag.Float64Var(&presence, presenceFlag, 0.3, globals.PresenceFlag.Help)
	}
	for _, bestOfFlag := range globals.BestOfFlag.Names {
		flag.IntVar(&bestof, bestOfFlag, 1, globals.BestOfFlag.Help)
	}
}
func main() {
	// setup the globals
	// parse the flags
	flag.Parse()
	globals.Setup(pretty)
	// if the input is empty
	if len(input) == 0 {
		globals.Log.Fatal().
			Msg("Input flag is empty (did you enter it, or is there any git diff?).")
	}
	// request the api
	request.RequestApi(input, model, maxtokens, temperature, top_p, frequence, presence, bestof)
}
