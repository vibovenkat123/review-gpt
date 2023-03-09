# Pricing

### Preface

Since this project is open source, you don't need to pay ***me*** anything. However, GPT 3 costs money.


### The `model` flag

The `--model` flag represents the models for GPT. You can find the list at [The GPT3.5 model page](https://platform.openai.com/docs/models/gpt-3-5) and [The GPT3 model page](https://platform.openai.com/docs/models/gpt-3)

Review-GPT only supports gpt-3.5-turbo, text-davinci-003, text-curie-001, text-babbage-001, and text-ada-001

Out of those 5 models, I would only recommend **gpt-3-5-turbo**

Turbo is a GPT3.5 model. It is more powerful and cheaper than any of the rest of the GPT3.5 models as well as the GPT3 models.

### The `max` flag

The `--max` flag represents the max tokens that GPT can use while generating code. If you decrease the flag, it will generate less tokens, thus making it cheaper.
