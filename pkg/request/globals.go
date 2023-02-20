package request
import (
    "errors"
    "fmt"
)
type Model = string
type Modeltype = string
const (
    Davinci string = "text-davinci-003"
    Curie string = "text-curie-001"
    Babbage  string = "text-babbage-001"
    Ada string = "text-ada-001"
    Codex Modeltype = "codex"
    Gpt Modeltype = "gpt"
)
var (
    TempRangeMin float32 = 0
    TempRangeMax  float32 = 1
    TopPMin float32 = 0
    TopPMax   float32 = 1
    PresenceMin float32 = 0
    PresenceMax   float32 = 2
    FrequenceMin float32 = 0
    FrequenceMax float32 = 2
    BestOfMin int = 1
    BestOfMax int = 20
)
var (
    ErrWrongModel error = errors.New(fmt.Sprintf("The model you entered was not correct"))
    ErrWrongTempRange error = errors.New(fmt.Sprintf("The temperature is not in the correct range (%v <= temp <= %v)", TempRangeMin, TempRangeMax))
    ErrWrongTopRange error = errors.New(fmt.Sprintf("The top_p is not in the correct range (%v <= top_p <= %v)", TopPMin, TopPMax))
    ErrWrongPresenceRange error = errors.New(fmt.Sprintf("The presence penalty is not in the correct range (%v <= presence <= %v)", PresenceMin, PresenceMax))
    ErrWrongFrequenceRange error = errors.New(fmt.Sprintf("The presence penalty is not in the correct range (%v <= frequence <= %v)", FrequenceMin, FrequenceMax))
    ErrWrongBestOfRange error = errors.New(fmt.Sprintf("The best of variable is not in the correct range (%v <= best of  <= %v)", BestOfMin, BestOfMax))
)
