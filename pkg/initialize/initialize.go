package initialize

import (
	// flag for system flags
	"errors"
	"flag"
	// globals (for initializing and using globals)
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

// if its json
var json bool

// the raw pretty flag
var rawJSON bool

// Amount of times to call
var bestof int

// if its verbose
var verbose bool

// the raw verbose flag
var rawVerbose bool

// if a flag is passed
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
// if any of the flags are passed
func areFlagsPassed(names []string) bool {
	for _, name := range names {
		if isFlagPassed(name) {
			return true
		}
	}
	return false
}
// set a flag
func setFlag(names []string, flagVar interface{}, defaultValue interface{}, help string) {
    // apply the correct flag function accordingly
	switch flagVar.(type) {
	case *string:
		for _, name := range names {
			flag.StringVar(flagVar.(*string), name, defaultValue.(string), help)
		}
	case *bool:
		for _, name := range names {
			flag.BoolVar(flagVar.(*bool), name, defaultValue.(bool), help)
		}
	case *int:
		for _, name := range names {
			flag.IntVar(flagVar.(*int), name, defaultValue.(int), help)
		}
	case *float64:
		for _, name := range names {
			flag.Float64Var(flagVar.(*float64), name, defaultValue.(float64), help)
		}
	}
}
func getFlags() {
    // set all the flags
	setFlag(globals.InputFlag.Names, &input, "", globals.InputFlag.Help)
	setFlag(globals.VerboseFlag.Names, &rawVerbose, false, globals.VerboseFlag.Help)
	setFlag(globals.JsonFlag.Names, &rawJSON, false, globals.JsonFlag.Help)
	setFlag(globals.ModelFlag.Names, &model, "text-davinci-003", globals.ModelFlag.Help)
	setFlag(globals.MaxTokenFlag.Names, &maxtokens, 500, globals.MaxTokenFlag.Help)
	setFlag(globals.TemperatureFlag.Names, &temperature, 0.2, globals.TemperatureFlag.Help)
	setFlag(globals.ToppFlag.Names, &top_p, 1.0, globals.ToppFlag.Help)
	setFlag(globals.FrequenceFlag.Names, &frequence, 1.2, globals.FrequenceFlag.Help)
	setFlag(globals.PresenceFlag.Names, &presence, 0.3, globals.PresenceFlag.Help)
	setFlag(globals.BestOfFlag.Names, &bestof, 1, globals.BestOfFlag.Help)
	flag.Parse()
    // the json and verbose flags dont need values
	json = areFlagsPassed(globals.JsonFlag.Names)
	verbose = areFlagsPassed(globals.VerboseFlag.Names)
}
// to check for flag validation
func validate(input string) error {
	if len(input) == 0 {
		err := errors.New("Input flag is empty (did you enter it, or is there any git diff?).")
		return err
	}
	return nil
}
// initialize review-gpt
func Init() {
	// Get the flags
	getFlags()
	// parse the flags
	// setup the globals
	globals.Setup(json, verbose)
	// if the input is empty
	if err := validate(input); err != nil {
		msg := err.Error()
		globals.Log.Fatal().
			Msg(msg)
	}
	// request the api
	request.RequestApi(input, model, maxtokens, temperature, top_p, frequence, presence, bestof)
}
