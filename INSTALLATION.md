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

4. Add it to your path 

### In git repo

```console
foo:~/.rgpt$ ./basecli
```
### Outside

```console
foo:~$ sudo cp rgpt /usr/local/bin
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

To use rgpt:

```console
foo@bar:~/random/git/repo$ rgpt --i "$(git diff file.txt)" --f "$(git show file.txt)"
```
