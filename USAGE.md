# USAGE

# Example on terminal

![Usage Example](/assets/examples/usage/base_usage.png)

**More in the [examples/usage](https://github.com/vibovenkat123/review-gpt/tree/main/assets/examples/usage) directory**

<br />

## Example with all flags:

```console
foo@bar:~$ rgpt --input "$(git diff file.txt)" --model "text-davinci-003" --max 500 --temp 0.2 --topp 1 --freq 1.2 --pres 0.3 --bestof 1
```

## Example with necessary flags:

```console
foo@bar:~$ rgpt --input "$(git diff file.txt)"
```

## Flags:

### Required:

`input`, `i`: The input file (the git diff). You can get the git diff of a file by executing `git diff file.txt` where `file.txt` is the file

### Optional
`pretty`, `pret`: If the output should be pretty (recommended for CLI yes, extension no)

Default: true

`model`, `m`: The model for GPT to use. Can be ("text-davinci-003", "text-curie-001", "text-ada-001", "text-babbage-001"), Davinci is recommended

Default: text-davinci-003

`max`: The maximum tokens to use (go to [OpenAI's tokenizer](https://platform.openai.com/tokenizer) to convert your text to tokens). More tokens = More expensive.

Default: 500

`temp`, `t`, `temperature`: What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.

Default: 0.2

`topp`:  An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.

Default: 1

`frequence`, `freq`, `fr`, `f`: Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.

Default: 1.2

`presence`, `pr`, `p`, `pres`: Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.

Default: 0.3

`bestof`, `bo`, `best`: Generates best_of completions server-side and returns the 'best' (the one with the highest log probability per token). Results cannot be streamed.

Default: 1


