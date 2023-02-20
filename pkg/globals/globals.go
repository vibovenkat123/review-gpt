package globals

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

var OpenaiKey string
var EnvFile string

func Setup() {
	EnvFile = ".rgpt.env"
	home := os.Getenv("HOME")
	err := godotenv.Load(fmt.Sprintf("%v/%v", home, EnvFile))
	if err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Fatalln(errors.New(".rgpt.env not found. Did you follow the instructions in the INSTALLATION.md?"))
		}
		log.Fatalln(err)
	}
	OpenaiKey = os.Getenv("OPENAI_KEY")
}
