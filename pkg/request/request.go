package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
    "github.com/vibovenkat123/review-gpt/pkg/globals"
	"net/http"
)

type Body struct {
	Model         Model  `json:"model"`
	Prompt        string  `json:"prompt"`
	Temperature   float64 `json:"temperature"`
	Max_Tokens    int     `json:"max_tokens"`
	Top_P         float64 `json:"top_p"`
	Frequence_Pen float64 `json:"frequency_penalty"`
	Presence_Pen  float64 `json:"presence_penalty"`
	Best_Of       int     `json:"best_of"`
	Suffix        string  `json:"suffix"`
}
type EditBody struct {
    Model string `json:"model"`
    Input string `json:"input"`
    Instruction string `json:"instruction"`
    N int `json:"n"`
    Temperature float64 `json:"temperature"`
    Top_P float64 `json:"top_p"`
}
type APIText struct {
	Text         string  `json:"text"`
	Index        int     `json:"index"`
}
type APIUsage struct {
	Prompt_Tokens     int `json:"prompt_tokens"`
	Completion_Tokens int `json:"completion_tokens"`
	Total_Tokens      int `json:"total_tokens"`
}
type APIRequest struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Choices []APIText `json:"choices"`
	Usage   APIUsage  `json:"usage"`
}

func RequestApi(gitDiff string, model Model, maxtokens int, temperature float64, top_p float64, frequence float64, presence float64, bestof int) {
    improvements, err := RequestImprovements(globals.OpenaiKey, gitDiff, model, maxtokens, temperature, top_p, frequence, presence, bestof)
    if err != nil {
        log.Fatalln(err)
    }
    for _, improvement:= range improvements {
        fmt.Println(improvement)
    }
}
func CheckFormat(body Body) error {
    if body.Model != Davinci && body.Model != Ada && body.Model != Curie && body.Model != Babbage {
        return ErrWrongModel
    }
    if body.Temperature < TempRangeMin || body.Temperature > TempRangeMax {
        return ErrWrongTempRange
    }
    if body.Top_P < TopPMin || body.Top_P > TopPMax {
        return ErrWrongTopRange
    }
    if body.Presence_Pen < PresenceMin || body.Presence_Pen >  PresenceMax {
        return ErrWrongPresenceRange
    }
    if body.Frequence_Pen < FrequenceMin || body.Frequence_Pen > FrequenceMax {
        return ErrWrongFrequenceRange
    }
    if body.Best_Of < BestOfMin || body.Best_Of > BestOfMax {
        return ErrWrongBestOfRange
    }
    return nil
}
func RequestImprovements(key string, gitDiff string, model Model, maxtokens int, temperature float64, top_p float64, frequence float64, presence float64, bestof int) ([]string, error){
    answers := []string{}
	url := "https://api.openai.com/v1/completions"
	promptInstruction := "Explain the git diff below, and from a code reviewers perspective, tell me what I can improve on in the code (the '+' in the git diff is an added line, the '-' is a removed line). DO NOT SUGGEST CHANGES ALREADY MADE IN THE GIT DIFF. DO NOT EXPLAIN THE GIT DIFF. ONLY  SAY WHAT COULD BE IMPROVED. Also go into more detail, but not too much"
	prompt := fmt.Sprintf("%#v\n%#v\n", promptInstruction, gitDiff)
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
    if err := CheckFormat(params); err != nil {
       return answers, err
    }
	jsonParams, err := json.Marshal(params)
	if err != nil {
		log.Fatalln(err)
	}
	reqBody := bytes.NewBuffer(jsonParams)
	req, err := http.NewRequest("POST", url, reqBody)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", key))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	apiReq := APIRequest{}
	json.Unmarshal([]byte(string(body)), &apiReq)
	choices := apiReq.Choices
	for _, c := range choices {
        if len(c.Text) != 0 {
            answers = append(answers, c.Text)
        }
	}
    return answers, nil
}
