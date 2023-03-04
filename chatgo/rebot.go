package chatgo

import (
	"strings"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Robot is a chatgpt robot
type Robot struct {
	name     string
	token    string
	model    string
	messages []*message
}

// NewRobot creates a new robot
func NewRobot() *Robot {
	return &Robot{}
}

// SetName sets the robot name
func (r *Robot) SetName(name string) *Robot {
	r.name = name
	return r
}

func (r *Robot) Name() string {
	if r.name == "" {
		return "default_robot"
	}
	return r.name
}

// SetToken sets the robot openai token
func (r *Robot) SetToken(token string) *Robot {
	r.token = token
	return r
}

// SetModel sets the robot openai model
func (r *Robot) SetModel(model string) *Robot {
	r.model = model
	return r
}

// GetModel gets the robot openai model
func (r *Robot) GetModel() string {
	if r.model == "" {
		return defaultModel
	}
	return r.model
}

// Tell tells the robot something
func (r *Robot) Tell(sentence string) string {
	r.messages = append(r.messages, &message{
		Role:    "user",
		Content: sentence,
	})

	result, err := r.tell()
	if err != nil {
		panic(err)
	}

	if len(result.Choices) == 0 {
		return "sorry, i don't know what to say."
	}

	latestMessage := result.Choices[len(result.Choices)-1].Message

	r.messages = append(r.messages, &message{
		Role:    latestMessage.Role,
		Content: latestMessage.Content,
	})

	return strings.TrimSpace(latestMessage.Content)
}
