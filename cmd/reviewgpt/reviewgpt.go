package main

import (
	"flag"
	"github.com/vibovenkat123/review-gpt/pkg/request"
	"github.com/vibovenkat123/review-gpt/pkg/globals"
	"log"
)

var input string
var model string
var modelType string
var strict bool
var maxtokens int
func main() {
    globals.Setup()
    flag.StringVar(&input, "input", "", "The input (git diff file.txt)")
    flag.StringVar(&input, "i", "", "The input (git diff file.txt)")
    flag.StringVar(&model, "model", "text-davinci-003", "The model for GPT (see USAGE.md for more details)")
    flag.StringVar(&model, "m", "text-davinci-003", "The model for GPT (see USAGE.md for more details)")
    flag.BoolVar(&strict, "strict", false, "If it is on strict mode or not (see USAGE.md for more details)")
    flag.IntVar(&maxtokens, "max", 500, "The length of the max tokens (see USAGE.md for more details)")
    flag.Parse()
    if len(input) == 0 {
        log.Fatalln("Enter an argument for --file or --input")
    }
	request.RequestApi(input, model,  maxtokens)
}
