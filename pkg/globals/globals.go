package globals

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)
// the openapi key
var OpenaiKey string
// the path of the environment variable
var EnvFile string

func Setup() {
    // set the environment file
	EnvFile = ".rgpt.env"
	home := os.Getenv("HOME")
    // load the environment file
	err := godotenv.Load(fmt.Sprintf("%v/%v", home, EnvFile))
	if err != nil {
        // if the error says the environement file doesn't exist
		if strings.Contains(err.Error(), "no such file") {
			log.Fatalln(errors.New(".rgpt.env not found. Did you follow the instructions in the INSTALLATION.md?"))
		}
		log.Fatalln(err)
	}
    // set the openapi key to the environment variable
	OpenaiKey = os.Getenv("OPENAI_KEY")
}
