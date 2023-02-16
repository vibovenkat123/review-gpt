package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
    "bufio"
)

type Body struct {
	Model         string  `json:"model"`
	Prompt        string  `json:"prompt"`
    Temperature   int     `json:"temperature"`
	Max_Tokens    int     `json:"max_tokens"`
	Top_P         int     `json:"top_p"`
	Frequence_Pen float32 `json:"frequency_penalty"`
	Presence_Pen  float32 `json:"presence_penalty"`
    Best_Of int `json:"best_of"`
    Suffix string `json:"suffix"`
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
func Call(promptText string) {
	pwd := os.Getenv("OPENAI_KEY")
	url := "https://api.openai.com/v1/completions"
	prompt := fmt.Sprintf("%v\nI WILL NOT write more code below, and if the code above is good, I will not write anything. As a code reviewer, I think that you could improve this code by doing the following improvements:\n", promptText)
    suffix := "\n\nThats all the improvements!!"
	params := Body{
		Model:         "code-davinci-002",
		Prompt:        prompt,
		Temperature:   0,
		Max_Tokens:    300,
		Top_P:         1,
		Frequence_Pen: 1.27,
		Presence_Pen:  0.58,
        Best_Of: 3,
        Suffix: suffix,
	}
	jsonParams, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	reqBody := bytes.NewBuffer(jsonParams)
	req, err := http.NewRequest("POST", url, reqBody)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", pwd))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	apiReq := APIRequest{}
	json.Unmarshal([]byte(string(body)), &apiReq)
	choices := apiReq.Choices
	for _, c := range choices {
        fmt.Println("A:", c.Text)
	}
}
func Request() {
    dir := ".rgpt"
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        panic("Run `rgptsetup commit` before running")
    }
    for _, file := range files {
        fileContent := ""
        f, err := os.Open(fmt.Sprintf("%v/%v", dir, file.Name()))
        if err != nil {
            panic(err)
        }
        defer f.Close()
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            fileContent += scanner.Text()
            fileContent += "\n"
        }
        if err := scanner.Err(); err != nil {
           panic(err)
        }
        Call(fileContent)
    }
}
