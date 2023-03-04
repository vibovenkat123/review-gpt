RGPT(1) -- Automate code reviews from your terminal
===================================================

## SYNOPSIS
`rgpt` [<options>] `-i`|`--input` <diff>

## DESCRIPTION
The **rgpt** command automatically code reviews your code - straight from your terminal

In the default settings, `rgpt` takes a <diff> file (e.g, a git(1) diff) and asks GPT3 on how to improve it.
The `--json` and `--verbose` options change the output of the CLI. These options can be used together. 

The `max`,  `temp`, `frequence`, `presence`, and `bestof` options dictate how GPT3 will parse the <diff>.

The `rgpt` needs an input <diff> in order to work. Every other flag is optional.

## THE DIFF

The `rgpt` command requires a input flag. This is input flag is recommended to be a <diff>.
Since GPT3 only allows a certain amount of <tokens>, It is recommended you only input a git diff for one file at a time.
