package globals;
import (
    "os"
	"github.com/joho/godotenv"
    "log"
    "fmt"
)
var OpenaiKey string
func Setup() {
    err := godotenv.Load(fmt.Sprintf("%v/.rgpt.env", os.Getenv("HOME")))
    if err != nil {
        log.Fatalln(err)
    }
    OpenaiKey = os.Getenv("OPENAI_KEY")
}
