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
	Best_Of       int     `json:"best_of"`
	Suffix        string  `json:"suffix"`
}

// the text in the choices the response gives
type APIText struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
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

// the response the api gives
type APIResponse struct {
	Err     *ApiErr   `json:"error"`
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Choices []APIText `json:"choices"`
	Usage   APIUsage  `json:"usage"`
}

// request the api
func RequestApi(gitDiff string, model Model, maxtokens int, temperature float64, top_p float64, frequence float64, presence float64, bestof int) {
	// get all the improvements
	improvements, err := RequestImprovements(globals.OpenaiKey, gitDiff, model, maxtokens, temperature, top_p, frequence, presence, bestof)
	if err != nil {
		globals.Sugar.Fatalln("Error while requesting the improvments",
            "Error", err,
            "Open API Key", globals.OpenaiKey,
            "Model", model,
            "Max Tokens", maxtokens,
            "Temperature", temperature,
            "Top_P", top_p,
            "Frequence Penalty", frequence,
            "Presence Penalty", presence,
            "Best Of", bestof,
            "Improvements received", improvements,
        )
	}
	// print each improvement
	for _, improvement := range improvements {
		fmt.Println(improvement)
	}
}

// checking the format
func CheckFormat(body Body) error {
	// model
	if body.Model != Davinci && body.Model != Ada && body.Model != Curie && body.Model != Babbage {
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
	// best of
	if body.Best_Of < BestOfMin || body.Best_Of > BestOfMax {
		return ErrWrongBestOfRange
	}
	// if its all good
	return nil
}

// request the improvements
func RequestImprovements(key string, gitDiff string, model Model, maxtokens int, temperature float64, top_p float64, frequence float64, presence float64, bestof int) ([]string, error) {
	answers := []string{}
	// request url
	url := "https://api.openai.com/v1/completions"
	// the instruction
	promptInstruction := "Explain the git diff below, and from a code reviewers perspective, tell me what I can improve on in the code (the '+' in the git diff is an added line, the '-' is a removed line). DO NOT SUGGEST CHANGES ALREADY MADE IN THE GIT DIFF. DO NOT EXPLAIN THE GIT DIFF. ONLY  SAY WHAT COULD BE IMPROVED. Also go into more detail, but not too much"
	// get the prompt using sprintf
	prompt := fmt.Sprintf("%#v\n%#v\n", promptInstruction, gitDiff)
	// get the body struct
	params := Body{
		Model:         model,
		Prompt:        prompt,
		Temperature:   temperature,
		Max_Tokens:    maxtokens,
		Top_P:         top_p,
		Frequence_Pen: frequence,
		Presence_Pen:  presence,
		Best_Of:       bestof,
	}
	// if the params are in the wrong format return an error
	if err := CheckFormat(params); err != nil {
		return answers, err
	}
	// marshal the params
	jsonParams, err := json.Marshal(params)
	if err != nil {
        return answers, err
	}
	// get the request body in bytes
	reqBody := bytes.NewBuffer(jsonParams)
	// form a new request
	req, err := http.NewRequest("POST", url, reqBody)
	// set the api key
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", key))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	// execute the request
	resp, err := client.Do(req)
	if err != nil {
        return answers, err
	}
	defer resp.Body.Close()
	// get the body
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
		if len(c.Text) != 0 {
			answers = append(answers, c.Text)
		}
	}
	return answers, nil
}
