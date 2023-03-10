package chatgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

const (
	// RoleAssistant is the assistant role
	RoleAssistant = "assistant"
	// RoleUser is the user role
	RoleUser = "user"
	// RoleSystem is the system role
	RoleSystem = "system"
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

// ChatStreamResponse is the response body for chatgpt stream
type ChatStreamResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Delta struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"delta"`
		Index        int `json:"index"`
		FinishReason any `json:"finish_reason"`
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

	req, err := http.NewRequest("POST", charURL, bytes.NewBuffer(data))
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

	defer func() {
		_ = resp.Body.Close()
	}()

	var result ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *Robot) tellStream() (<-chan *ChatStreamResponse, error) {
	payload := map[string]interface{}{
		"model":    r.GetModel(),
		"stream":   true,
		"messages": r.messages,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", charURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http2.Transport{},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	stream := make(chan *ChatStreamResponse, 10)

	go func() {
		defer func() {
			stream <- nil
		}()
		for {
			// 流式读取body
			buffer := make([]byte, 65535)
			n, err := resp.Body.Read(buffer)
			if err != nil {
				break
			}
			buffer = buffer[:n]
			messages := bytes.Split(buffer, []byte("data:"))
			for _, m := range messages {
				if len(m) != 0 {
					rdata := &ChatStreamResponse{}
					m = bytes.TrimSpace(m)
					if err := json.Unmarshal(m, rdata); err != nil {
						continue
					}
					stream <- rdata
				}
			}
		}
	}()

	return stream, nil
}
