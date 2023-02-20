package main

import (
	"flag"
	"github.com/vibovenkat123/review-gpt/pkg/request"
	"github.com/vibovenkat123/review-gpt/pkg/globals"
	"log"
)

var input string
var model string
var maxtokens int
var temperature float64
var top_p float64
var frequence float64
var presence float64
var bestof int
func main() {
    globals.Setup()
    flag.StringVar(&input, "input", "", "The input (git diff file.txt)")
    flag.StringVar(&input, "i", "", "The input (git diff file.txt)")
    flag.StringVar(&model, "model", "text-davinci-003", "The model for GPT (see USAGE.md for more details)")
    flag.StringVar(&model, "m", "text-davinci-003", "The model for GPT (see USAGE.md for more details)")
    flag.IntVar(&maxtokens, "max", 500, "The length of the max tokens (see USAGE.md for more details)")
    flag.Float64Var(&temperature, "temp", 0.2, "What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.")
    flag.Float64Var(&temperature, "t", 0.2, "What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.")
    flag.Float64Var(&temperature, "temperature", 0.2, "What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.")
    flag.Float64Var(&top_p, "topp", 1, "An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.")
    flag.Float64Var(&frequence, "freq", 1.2, "Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.")
    flag.Float64Var(&frequence, "fr", 1.2, "Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.")
    flag.Float64Var(&frequence, "f", 1.2, "Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.")
    flag.Float64Var(&frequence, "frequence", 1.2, "Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.")
    flag.Float64Var(&presence, "presence", 0.3, "Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.")
    flag.Float64Var(&presence, "pr", 0.3, "Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.")
    flag.Float64Var(&presence, "p", 0.3, "Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.")
    flag.Float64Var(&presence, "pres", 0.3, "Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.")
    flag.IntVar(&bestof, "bestof", 1, "Generates best_of completions server-side and returns the 'best' (the one with the highest log probability per token). Results cannot be streamed.")
    flag.IntVar(&bestof, "bo", 1, "Generates best_of completions server-side and returns the 'best' (the one with the highest log probability per token). Results cannot be streamed.")
    flag.IntVar(&bestof, "best", 1, "Generates best_of completions server-side and returns the 'best' (the one with the highest log probability per token). Results cannot be streamed.")
    flag.Parse()
    if len(input) == 0 {
        log.Fatalln("Enter an argument for --file or --input")
    }
	request.RequestApi(input, model,  maxtokens, temperature, top_p, frequence, presence, bestof)
}
