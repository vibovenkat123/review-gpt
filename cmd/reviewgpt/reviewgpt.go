package main

import (
	// flag for system flags
	"flag"
	// globals (for initializing globals
	"github.com/vibovenkat123/review-gpt/pkg/globals"
	// to call the API
	"github.com/vibovenkat123/review-gpt/pkg/request"
	// for errors
	"log"
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
// Amount of times to call
var bestof int

func main() {
    // all the help messages
    inputHelp := "The input (git diff file.txt)"
    modelHelp := "The model for GPT (see USAGE.md for more details)"
    maxTokensHelp := "The length of the max tokens (see USAGE.md for more details)"
    temperatureHelp := "What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic."
    topPHelp := "An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered."
    frequenceHelp := "Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim."
    presenceHelp := "Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics." 
    bestOfHelp := "Generates best_of completions server-side and returns the 'best' (the one with the highest log probability per token). Results cannot be streamed."
    // the flag arrays
    inputFlags := []string{"input", "i"}
    modelFlags := []string{"model", "m"}
    maxTokensFlags:= []string{"max"}
    temperatureFlags := []string{"temp", "t"}
    toppFlags:= []string{"topp"}
    frequenceFlags:= []string{"frequence", "freq", "fr", "f"}
    presenceFlags  := []string{"pr", "presence", "p", "pres"}
    bestOfFlags := []string{"bo", "bestof", "best"}
    // setup the globals
	globals.Setup()
    // setup the flag by looping through the flags array and setting them
    for _, inputFlag := range inputFlags {
        flag.StringVar(&input, inputFlag, "", inputHelp)
    }
    for _, modelFlag := range modelFlags {
        flag.StringVar(&model, modelFlag, "text-davinci-003", modelHelp)
    }
    for _, maxTokensFlag := range maxTokensFlags {
        flag.IntVar(&maxtokens, maxTokensFlag, 500, maxTokensHelp)
    }
    for _, temperatureFlag := range temperatureFlags {
        flag.Float64Var(&temperature, temperatureFlag, 0.2, temperatureHelp)
    }
    for _, toppFlag := range toppFlags {
        flag.Float64Var(&top_p, toppFlag, 1, topPHelp)
    }
    for _, frequenceFlag := range frequenceFlags {
        flag.Float64Var(&frequence, frequenceFlag, 1.2, frequenceHelp)
    }
    for _, presenceFlag := range presenceFlags {
        flag.Float64Var(&presence, presenceFlag, 0.3, presenceHelp)
    }
    for _, bestOfFlag := range bestOfFlags {
        flag.IntVar(&bestof, bestOfFlag, 1, bestOfHelp)
    }
    // parse the flags
	flag.Parse()
    // if the input is empty
	if len(input) == 0 {
		log.Fatalln("Input flag is empty (did you enter it, or is there any git diff?)")
	}
    // request the api
	request.RequestApi(input, model, maxtokens, temperature, top_p, frequence, presence, bestof)
}
