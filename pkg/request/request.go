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
	Model         globals.Model   `json:"model"`
	Prompt        string  `json:"prompt"`
	Temperature   float64 `json:"temperature"`
	Max_Tokens    int     `json:"max_tokens"`
	Top_P         float64 `json:"top_p"`
	Frequence_Pen float64 `json:"frequency_penalty"`
	Presence_Pen  float64 `json:"presence_penalty"`
    Best_Of       int     `json:"best_of"`
}
// the struct to use for gpt3.5
type TurboBody struct {
	Model         globals.Model     `json:"model"`
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
func RequestApi(gitDiff string, model globals.Model, maxtokens int, temperature float64, top_p float64, frequence float64, presence float64, bestof int) {
	LogVerbose("Requesting for improvements")
	// get all the improvements
	improvements, err := RequestImprovements(globals.OpenaiKey, gitDiff, model, maxtokens, temperature, top_p, frequence, presence, bestof)
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
	if body.Model != globals.Davinci && body.Model != globals.Ada && body.Model != globals.Curie && body.Model != globals.Babbage && body.Model != globals.Turbo {
		return globals.ErrWrongModel
	}
	// temperature
	if body.Temperature < globals.TempRangeMin || body.Temperature > globals.TempRangeMax {
		return globals.ErrWrongTempRange
	}
	// top_p
	if body.Top_P < globals.TopPMin || body.Top_P > globals.TopPMax {
		return globals.ErrWrongToppRange
	}
	// presense penalty
	if body.Presence_Pen < globals.PresenceMin || body.Presence_Pen > globals.PresenceMax {
		return globals.ErrWrongPresenceRange
	}
	// frequence penalty
	if body.Frequence_Pen < globals.FrequenceMin || body.Frequence_Pen > globals.FrequenceMax {
		return globals.ErrWrongFrequenceRange
	}
    // the best_of
    if  body.Best_Of < globals.BestOfMin || body.Best_Of > globals.BestOfMax {
        return globals.ErrWrongBestOfRange
    }
	// if its all good
	return nil
}

// request the improvements
func RequestImprovements(key string, gitDiff string, model globals.Model, maxtokens int, temperature float64, top_p float64, frequence float64, presence float64, bestof int) ([]string, error) {
	answers := []string{}
	// get the normal GPT3 body struct 
	params := Body{
		Model:         model,
		Temperature:   temperature,
		Max_Tokens:    maxtokens,
		Top_P:         top_p,
		Frequence_Pen: frequence,
		Presence_Pen:  presence,
        Best_Of: bestof,
	}
	// Get the strict for GPT3.5
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
	if model == globals.Turbo {
		endUrl = "chat/completions"
	}
	// request url
	url := fmt.Sprintf("https://api.openai.com/v1/%v", endUrl)
	// the instruction
	promptInstruction := "explain the git diff below, and from a code reviewers perspective, tell me what i can improve on in the code (the '+' in the git diff is an added line, the '-' is a removed line). do not suggest changes already made in the git diff. do not explain the git diff. only  say what could be improved. also go into more detail, but not too much"
	// The background information for turbo/gpt3.5
	turboPromptInstruction := "You are a very intelligent code reviewer. You will take in a git diff, and tell the user what they could have improved (like a code review) based on analyzing the git diff in order to see whats changed.\nYou will not provide any examples/code snippets in your answer"
	// get the prompt using sprintf
	prompt := fmt.Sprintf("%#v\n%#v\n", promptInstruction, gitDiff)
	if model == globals.Turbo {
		// The background message
		sysMessage := Message{
			Role:    "system",
			Content: turboPromptInstruction,
		}
		// The input (what they respond to)
		usrMessage := Message{
			Role:    "user",
			Content: gitDiff,
		}
		// the message for turbo
		turboParams.Messages = []Message{sysMessage, usrMessage}
	} else {
		// set the gpt3 prompt to the prompt defined before
		params.Prompt = prompt
	}
	// marshal the params
	var jsonParams []byte
	var err error
	// marshal the correct param struct
	if model == globals.Turbo {
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
		// if its GPT3.5, its structured differently
		if model == globals.Turbo {
			answers = append(answers, c.Message.Content)
			continue
		}
		// if its not empty
		if len(c.Text) != 0 {
			answers = append(answers, c.Text)
		}
	}
	return answers, nil
}
