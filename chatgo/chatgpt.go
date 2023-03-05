package chatgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ChatResponse is the response body for chatgpt
type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

// ChatRequest is the request body for chatgpt
type ChatRequest struct {
	Token   string `json:"token"`
	Model   string `json:"model"`
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (r *Robot) tell() (*ChatResponse, error) {
	payload := map[string]interface{}{
		"model":    r.GetModel(),
		"messages": r.messages,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", charUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *Robot) tellStream() (*ChatResponse, error) {

	return nil, nil
}
