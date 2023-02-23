# CLI

# One-Liner

```console
foo@bar:~$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/vibovenkat123/review-gpt/HEAD/install.sh)"
```

# If that doesnt work, build it from source
## Prereqs:

[Go](https://go.dev)

## Clone the repo
```console
foo@bar:~$ git clone git@github.com:vibovenkat123/review-gpt.git ~/.rgpt
```
### NOTE: For tidiness, it is highly recommended you install the repo to `~/.rgpt`, unless if you are going to work on this

## Set the environment variable (In git repo)

Create a new file called .rgpt.env in your home directory

```console
foo@bar:~$ touch ~/.rgpt.env
```

Add the environment variable to the file, e.g OPENAI_KEY="<key_here>". It must be formatted like that.

```console
foo@bar:~$ echo OPENAI_KEY="sk-1234" > ~/.rgpt.env
```

## Build the binaries

Run `make build` to build the binaries (in git repo)

```console
foo@bar:~/.rgpt$ make build
```

## Add the binaries to your path (in git repo)

```console
foo@bar:~/.rgpt$ sudo mv ./bin/rgpt /usr/local/bin/rgpt
foo@bar:~/.rgpt$ sudo chmod +x /usr/local/bin/rgpt
```

## Done!!
