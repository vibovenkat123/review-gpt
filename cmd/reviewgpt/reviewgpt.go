package main

import (
	"flag"
	"github.com/vibovenkat123/review-gpt/pkg/request"
	"log"
	"os"
	"os/exec"
)

var action request.Action

func main() {
	flag.StringVar(&action, "action", "merge", "The action you want to do (commit | merge)")
	flag.Parse()
	setupPath, err := exec.LookPath("rgptsetup")
	if err != nil {
		log.Fatalln("Set up rgptsetup in order to use rgpt, see instructions in INSTALLATION.md for more info")
	}
	cmd := &exec.Cmd{
		Path:   setupPath,
		Args:   []string{setupPath, "invalid"},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if action == request.Commit {
		cmd.Args[1] = "commits"
	} else if action == request.Merge {
		cmd.Args[1] = "merge"
	} else {
		log.Fatalln("enter right arguments")
	}
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	request.Request()
}
