package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Body struct {
	Model         string  `json:"model"`
	Prompt        string  `json:"prompt"`
	Temperature   int     `json:"temperature"`
	Max_Tokens    int     `json:"max_tokens"`
	Top_P         int     `json:"top_p"`
	Frequence_Pen float32 `json:"frequency_penalty"`
	Presence_Pen  float32 `json:"presence_penalty"`
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

func Request() {
	pwd := os.Getenv("OPENAI_KEY")
	url := "https://api.openai.com/v1/completions"
	prompt := "package expenseFunctions\n\nfunc Validate(id *string, name *string, spent *int) (idIsGood bool, nameIsGood bool, spentIsGood bool) {\n\tidIsGood = false\n\tnameIsGood = false\n\tspentIsGood = false\n\tif id != nil && len(*id) == 36 {\n\t\tidIsGood = true\n\t}\n\tif name != nil && len(*name) > 0 {\n\t\tnameIsGood = true\n\t}\n\tif spent != nil && *spent >= 0 {\n\t\tspentIsGood = true\n\t}\n\treturn idIsGood, nameIsGood, spentIsGood\n}\nFrom a code reviewer's perspective, what is your comments on the code above, and what should be done differently? Do not include any code in your answer.\n\n"
	params := Body{
		Model:         "code-davinci-002",
		Prompt:        prompt,
		Temperature:   0,
		Max_Tokens:    256,
		Top_P:         1,
		Frequence_Pen: 1.27,
		Presence_Pen:  0.0,
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
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	apiReq := APIRequest{}
	json.Unmarshal([]byte(string(body)), &apiReq)
	choices := apiReq.Choices
	for _, c := range choices {
		fmt.Println(c.Text)
	}
}
