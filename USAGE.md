# USAGE

# Example on terminal

![Usage Example](/assets/examples/usage/base_usage.png)

**More in the [examples/usage](https://github.com/vibovenkat123/review-gpt/tree/main/assets/examples/usage) directory**

<br />

## Example with all flags:

```console
foo@bar:~$ rgpt --input "$(git diff file.txt)" --model "davinci" --max 500 --temp 0.2 --topp 1 --freq 1.2 --pres 0.3 --bestof 1 --json --verbose
```

## Example with necessary flags:

```console
foo@bar:~$ rgpt --input "$(git diff file.txt)"
```

## You can see all the flags and more info on [The manual page](https://github.com/vibovenkat123/review-gpt/blob/main/man/rgpt.1.ronn)
