
# Setup CLI (required)
## Clone repo
Clone the repo to the preffered spot

NOTE: IT IS HIGHLY RECOMMENDED THAT YOU CLONE TO `~/.review-gpt` FOR A STANDARD INSTALL

```console
foo@bar:~$ git clone git@github.com:vibovenkat123/review-gpt.git ~/.review-gpt
```
NOTE: FOR A STANDARD INSTALL IT IS HIGHLY RECOMMENDED THAT YOU DON'T CHANGE ANY OF THE SOURCE CODE

## Add to the path

Execute the addtopath command with bash:
```console
foo@bar:~$ bash addtopath.sh
```
Follow the instructions it tells you to do (add the correct destination to the correct source file)
## Run the command

For now, the rgptsetup command only supports commits, it will support merges in the future

To setup rgpt with commits on a repo, execute this

```console
foo@bar:~$ rgptsetup commits
```

## Uninstallation

Script coming soon, for now just delete the directory and the path commands


