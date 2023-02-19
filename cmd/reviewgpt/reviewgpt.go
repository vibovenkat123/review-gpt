package main

import (
	"flag"
//	"github.com/vibovenkat123/review-gpt/pkg/request"
	"log"
)

var input string
var file string
var model string
var typeof string
var strict bool
var maxtokens int
func main() {
    flag.StringVar(&input, "input", "", "The input (git diff file.txt)")
    flag.StringVar(&input, "i", "", "The input (git diff file.txt)")
    flag.StringVar(&file, "file", "", "The original file (git show HEAD:file.txt)")
    flag.StringVar(&file, "f", "", "The original file (git show HEAD:file.txt)")
    flag.StringVar(&model, "model", "text-davinci-003", "The model for GPT (see USAGE.md for more details)")
    flag.StringVar(&model, "m", "text-davinci-003", "The model for GPT (see USAGE.md for more details)")
    flag.StringVar(&typeof, "type", "codex", "The type of the model (gpt | codex)")
    flag.BoolVar(&strict, "strict", false, "If it is on strict mode or not (see USAGE.md for more details)")
    flag.IntVar(&maxtokens, "max", len(input) + 200, "The length of the max tokens (see USAGE.md for more details)")
    flag.Parse()
    log.Println(input, file, model, typeof, strict, maxtokens)
    if len(input) == 0 || len(file) == 0 {
        log.Fatalln("Enter an argument for --file or --input")
    }
//	request.Request(input, file)
}
