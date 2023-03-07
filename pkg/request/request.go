package request

import (
	// for marshalling
	"bytes"
	// for json
	"encoding/json"
	"fmt"

	// for getting the api key
	"github.com/vibovenkat123/review-gpt/pkg/globals"
	// for reading the response
	"io/ioutil"
	// for errors
	"errors"
	// for the requests
	"net/http"
)

// the struct to use for the body of the request
type Body struct {
	Model         Model   `json:"model"`
	Prompt        string  `json:"prompt"`
	Temperature   float64 `json:"temperature"`
	Max_Tokens    int     `json:"max_tokens"`
	Top_P         float64 `json:"top_p"`
	Frequence_Pen float64 `json:"frequency_penalty"`
	Presence_Pen  float64 `json:"presence_penalty"`
}
type TurboBody struct {
	Model         Model     `json:"model"`
	Messages      []Message `json:"messages"`
	Temperature   float64   `json:"temperature"`
	Max_Tokens    int       `json:"max_tokens"`
	Top_P         float64   `json:"top_p"`
	Frequence_Pen float64   `json:"frequency_penalty"`
	Presence_Pen  float64   `json:"presence_penalty"`
}

// the text in the choices the response gives
type APIText struct {
	Text    string  `json:"text"`
	Message Message `json:"message"`
	Index   int     `json:"index"`
}

// the usage the response gives
type APIUsage struct {
	Prompt_Tokens     int `json:"prompt_tokens"`
	Completion_Tokens int `json:"completion_tokens"`
	Total_Tokens      int `json:"total_tokens"`
}
type ApiErr struct {
	Message string  `json:"message"`
	Type    string  `json:"type"`
	Param   *string `json:"param"`
	Code    string  `json:"code"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// the response the api gives
type APIResponse struct {
	Err     *ApiErr   `json:"error"`
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Choices []APIText `json:"choices"`
	Usage   APIUsage  `json:"usage"`
}

func LogVerbose(msg string) {
	if globals.Verbose {
		globals.Log.Info().
			Msg(msg)
	}
}

// request the api
func RequestApi(gitDiff string, model Model, maxtokens int, temperature float64, top_p float64, frequence float64, presence float64) {
	LogVerbose("Requesting for improvements")
	// get all the improvements
	improvements, err := RequestImprovements(globals.OpenaiKey, gitDiff, model, maxtokens, temperature, top_p, frequence, presence)
	LogVerbose("Got improvements")
	if err != nil {
		globals.Log.Error().
			Str("Model", model).
			Int("Max Tokens", maxtokens).
			Float64("Temperature", temperature).
			Float64("Top_P", top_p).
			Float64("Frequence Penalty", frequence).
			Float64("Presence Penalty", presence).
			Err(err).
			Msg("Error while getting improvements")
	}
	// print each improvement
	for _, improvement := range improvements {
		fmt.Println(improvement)
	}
}

// checking the format
func CheckFormat(body Body) error {
	// model
	if body.Model != Davinci && body.Model != Ada && body.Model != Curie && body.Model != Babbage && body.Model != Turbo {
		return ErrWrongModel
	}
	// temperature
	if body.Temperature < TempRangeMin || body.Temperature > TempRangeMax {
		return ErrWrongTempRange
	}
	// top_p
	if body.Top_P < TopPMin || body.Top_P > TopPMax {
		return ErrWrongToppRange
	}
	// presense penalty
	if body.Presence_Pen < PresenceMin || body.Presence_Pen > PresenceMax {
		return ErrWrongPresenceRange
	}
	// frequence penalty
	if body.Frequence_Pen < FrequenceMin || body.Frequence_Pen > FrequenceMax {
		return ErrWrongFrequenceRange
	}
	// if its all good
	return nil
}

// request the improvements
func RequestImprovements(key string, gitDiff string, model Model, maxtokens int, temperature float64, top_p float64, frequence float64, presence float64) ([]string, error) {
	answers := []string{}
	// get the body struct
	params := Body{
		Model:         model,
		Temperature:   temperature,
		Max_Tokens:    maxtokens,
		Top_P:         top_p,
		Frequence_Pen: frequence,
		Presence_Pen:  presence,
	}
	turboParams := TurboBody{
		Model:         model,
		Temperature:   temperature,
		Max_Tokens:    maxtokens,
		Top_P:         top_p,
		Frequence_Pen: frequence,
		Presence_Pen:  presence,
	}
	// if the params are in the wrong format return an error
	if err := CheckFormat(params); err != nil {
		return answers, err
	}
	// the end of the url
	endUrl := "completions"
	if model == Turbo {
		endUrl = "chat/completions"
	}
	// request url
	url := fmt.Sprintf("https://api.openai.com/v1/%v", endUrl)
	// the instruction
	promptInstruction := "explain the git diff below, and from a code reviewers perspective, tell me what i can improve on in the code (the '+' in the git diff is an added line, the '-' is a removed line). do not suggest changes already made in the git diff. do not explain the git diff. only  say what could be improved. also go into more detail, but not too much"
	turboPromptInstruction := "You are a very intelligent code reviewer. You take in a git diff from a user(the '+' in the git diff is an added line, the '-' is a removed line), and then list all the improvements the user could have made. Go in to more detail, but not to the point where its too much. You will never write any code, only tell the improvements"
	// get the prompt using sprintf
	prompt := fmt.Sprintf("%#v\n%#v\n", promptInstruction, gitDiff)
	if model == Turbo {
		// get the prompt using sprintf
		sysMessage := Message{
			Role:    "system",
			Content: turboPromptInstruction,
		}
		usrMessage := Message{
			Role:    "user",
			Content: gitDiff,
		}
		turboParams.Messages = []Message{sysMessage, usrMessage}
	} else {
		params.Prompt = prompt
	}
	// marshal the params
	var jsonParams []byte
	var err error
	if model == Turbo {
		jsonParams, err = json.Marshal(turboParams)
	} else {
		jsonParams, err = json.Marshal(params)
	}
	if err != nil {
		return answers, err
	}
	// get the request body in bytes
	reqBody := bytes.NewBuffer(jsonParams)
	// form a new request
	LogVerbose("Creating new request")
	req, err := http.NewRequest("POST", url, reqBody)
	// set the api key
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", key))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	// execute the request
	LogVerbose("Requesting GPT")
	resp, err := client.Do(req)
	if err != nil {
		return answers, err
	}
	defer resp.Body.Close()
	// get the body
	LogVerbose("Got back the request information")
	body, _ := ioutil.ReadAll(resp.Body)
	apiReq := APIResponse{}
	// unmarshal (put the json in a struct) the body
	json.Unmarshal([]byte(string(body)), &apiReq)
	if apiReq.Err != nil {
		err := apiReq.Err
		return answers, errors.New(err.Message)
	}
	// get all the choices
	choices := apiReq.Choices
	// append it to the answers array
	for _, c := range choices {
		// if its not empty
		if model == Turbo {
			answers = append(answers, c.Message.Content)
			continue
		}
		if len(c.Text) != 0 {
			answers = append(answers, c.Text)
		}
	}
	return answers, nil
}
