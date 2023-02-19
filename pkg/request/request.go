package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Body struct {
	Model         string  `json:"model"`
	Prompt        string  `json:"prompt"`
	Temperature   float32 `json:"temperature"`
	Max_Tokens    int     `json:"max_tokens"`
	Top_P         int     `json:"top_p"`
	Frequence_Pen float32 `json:"frequency_penalty"`
	Presence_Pen  float32 `json:"presence_penalty"`
	Best_Of       int     `json:"best_of"`
	Suffix        string  `json:"suffix"`
}
type APIText struct {
	Text         string  `json:"text"`
	Index        int     `json:"index"`
	LogProbs     *string `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
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

func Request(promptText string, fileContent string) {
    err := godotenv.Load(fmt.Sprintf("%v/.rgpt.env", os.Getenv("HOME")))
    if err != nil {
        log.Fatalln(err)
    }
    answers := RequestChanges(promptText, os.Getenv("OPENAI_KEY"))
    answer := answers[0]
    fmt.Println(answer)
    code :=RequestCode(answer, fileContent, os.Getenv("OPENAI_KEY"))
    fmt.Println(code[0])
}
func RequestCode(answer string, fileContent string, key string) []string {
    answers := []string{}
	url := "https://api.openai.com/v1/completions"
    prompt := fmt.Sprintf("%v%v%v", "Original File:\n\n", fileContent, "\nNew code:")
    suffix := fmt.Sprintf("%v%v", "\n\nChanges:", answer)
	params := Body{
		Model:         "code-davinci-002",
		Prompt:        prompt,
		Temperature:   0,
		Max_Tokens:    1000,
		Top_P:         1,
		Frequence_Pen: 0,
		Presence_Pen:  0,
		Best_Of:       1,
        Suffix: suffix,
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
    return answers
}
func RequestChanges(promptText string, key string) []string{
    answers := []string{}
	url := "https://api.openai.com/v1/completions"
	promptInstruction := "State the differences (where a 'difference' is e.g. an addition or removal of some small number of lines; deciding what constitutes a difference is a hard task on its own, that could be approached through e.g. matching identifiers appearing in consecutive lines with other consecutive lines). Then tell me the intentions of these differences."
	prompt := fmt.Sprintf("%#v\n%#v\n\n", promptText, promptInstruction)
	params := Body{
		Model:         "text-davinci-003",
		Prompt:        prompt,
		Temperature:   0.1,
		Max_Tokens:    100,
		Top_P:         1,
		Frequence_Pen: 1.2,
		Presence_Pen:  0.11,
		Best_Of:       1,
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
    return answers
}
