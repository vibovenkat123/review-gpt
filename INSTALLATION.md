
# Setup CLI (required)

## Clone the repo

```console
foo@bar:~$ git clone git@github.com:vibovenkat123/review-gpt.git ~/.rgpt
```
NOTE: For a standard install, it is recommended you clone to ~/.rgpt for tidiness
## Set the environment variable
Setup the environment variable in the file .env.example

Then move it to the file .env 

```console
foo@bar:~/.rgpt$ mv .env.example .env
```

## Set up your path 

```console
foo@bar:~/.rgpt$ ./setupcli
Password: 
success
```

It needs sudo privillegas to copy it to your PATH

# Actual CLI (required)

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

4. Add it to your path (In git repo)

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

# Usage

For now, rgpt only supports commits

To use commits:

```console
foo@bar:~/random/git/repo$ rgpt --action commits
```
