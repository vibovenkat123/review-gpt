# Info

# Pricing

### Preface

Since this project is open source, you don't need to pay ***me*** anything. However, GPT 3 costs money.

### The `type` flag

In order to cut costs, you can first change the `--type` flag to `codex`, since codex is free at the moment. However, for analyzing the git diff, it still uses GPT 3. Go to to [Codex vs GPT3](#codex-vs-gpt3)

### The `model` flag

The `--model` flag represents the models for GPT. You can find a list at [their official model page]([https://platform.openai.com/docs/models/gpt-3](https://platform.openai.com/docs/models/gpt-3)). You can also find the prices of the individual models at [their pricing page]([https://openai.com/api/pricing](https://openai.com/api/pricing/))

Out of those 4 models, I would only recommend two, **Davinci** and **Curie.**

**Davinci is best model**, as it can allow more text, has more knowlege, and is better at analyzing the git diffs and producing code. **NOTE, IF YOU CHOOSE `gpt` FOR THE TYPE PARAMETER, IT IS RECOMMENDED YOU CHOOSE DAVINCI**

**Curie is way cheaper**, to be exact, 10x cheaper (meaning if something in davinci was $10, it would be $1 in curie). But that comes at a cost. It is not as accurate while talking about the git diff, and can produce faulty information.

### The `max` flag

The `--max` flag represents the max tokens that GPT can use while generating code. If you decrease the flag, it will generate less tokens, thus making it cheaper.

# Codex vs GPT3

Codex is built for completing code. **Codex will be better at generating “correct” (correct syntax, etc.) code**, while GPT3 may occasionally produce wrong code. However, since codex is meant for completing code, it isn’t as good for following instructions, and **GPT3 understands the instructions better than Codex.**

**Codex is free at this time**, and **GPT still costs money**.
