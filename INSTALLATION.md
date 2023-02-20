# CLI

# One-Liner

```console
foo@bar:~$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/vibovenkat123/review-gpt/HEAD/install.sh)"
```

# If that doesnt work, follow the steps below

## Clone the repo
```console
foo@bar:~$ git clone git@github.com:vibovenkat123/review-gpt.git ~/.rgpt
```
### NOTE: For tidiness, it is highly recommended you install the repo to `~/.rgpt`, unless if you are going to work on this

## Set the environment variable (In git repo)

Setup the environment variable in the file .env.example

Then move it to the file .env 

```console
foo@bar:~/.rgpt$ mv .env.example .env
```

Then copy the file to ~/.rgpt.env

```console
foo@bar:~/.rgpt$ cp .env ~/.rgpt.env
```

## Prebuilt Binaries

1. Go to the [releases page](https://github.com/vibovenkat123/review-gpt/releases) and download the right binary

2. Rename the binary file to `rgpt`

```console
foo@bar:~$ mv rgpt-os-arch rgpt
```

3. Add the binary file as a variable

```console
foo@bar:~$ export RGPTPATH_BIN="/path/to/binary/file"
```

4. Add the binaries to your path 

### In git repo

```console
foo:~/.rgpt$ ./basecli
```

## From source

1. Have [Go](https://go.dev) installed

3. Build the binaries  (In git repo)

```console
foo@bar:~/.rgpt$ make build
```

4. Add them to your path
```console
foo@bar:~/.rgpt$ make update
```

# Updates

To update the git repo, run `make update` in the directory
