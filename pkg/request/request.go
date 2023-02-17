package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
    "bufio"
    "github.com/joho/godotenv"
    "log"
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
    err := godotenv.Load(fmt.Sprintf("%v/.rgpt.env", os.Getenv("HOME")))
    if err != nil {
        log.Fatalln(err)
//        log.Fatalln("Env file not found, Did you forget to save it? Look in INSTALLATION.md for more details")
    }
	pwd := os.Getenv("OPENAI_KEY")
	url := "https://api.openai.com/v1/completions"
    promptInstruction := "From a code reviewer's perspective, what is your improvements on the code above, and what should be done differently? Do not include any code in your answer.\n/*"
	prompt := fmt.Sprintf("%v\n%v", promptText, promptInstruction)
	params := Body{
		Model:         "code-davinci-002",
		Prompt:        prompt,
		Temperature:   0,
		Max_Tokens:    300,
		Top_P:         1,
		Frequence_Pen: 1.27,
		Presence_Pen:  0.58,
        Best_Of: 3,
	}
	jsonParams, err := json.Marshal(params)
	if err != nil {
		log.Fatalln(err)
	}
	reqBody := bytes.NewBuffer(jsonParams)
	req, err := http.NewRequest("POST", url, reqBody)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", pwd))
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
        fmt.Println("A:", c.Text)
	}
}
func Request() {
    dir := ".rgpt"
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatalln("Run `rgptsetup commit` before running")
    }
    for _, file := range files {
        fileContent := ""
        f, err := os.Open(fmt.Sprintf("%v/%v", dir, file.Name()))
        if err != nil {
            log.Fatalln(err)
        }
        defer f.Close()
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            fileContent += scanner.Text()
            fileContent += "\n"
        }
        if err := scanner.Err(); err != nil {
           log.Fatalln(err)
        }
        Call(fileContent)
    }
}
