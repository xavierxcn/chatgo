package chatgo

import (
	"fmt"
	"os"
	"strings"
	"time"
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
	CreateAt time.Time `json:"create_at"`
}

// NewRobot creates a new robot
func NewRobot() *Robot {
	return &Robot{
		CreateAt: time.Now(),
	}
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
func (r *Robot) Tell(sentence string) []string {
	r.messages = append(r.messages, &message{
		Role:    "user",
		Content: sentence,
	})

	result, err := r.tell()
	if err != nil {
		panic(err)
	}

	if len(result.Choices) == 0 {
		return []string{"sorry, i don't know what to say."}
	}

	latestMessage := result.Choices[len(result.Choices)-1].Message

	r.messages = append(r.messages, &message{
		Role:    latestMessage.Role,
		Content: latestMessage.Content,
	})

	return strings.Split(latestMessage.Content, "\n")
}

// Save saves the messages to file
func (r *Robot) Save(filename string) error {
	// 将messages保存到HistoryPath
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	for _, m := range r.messages {
		_, err := f.WriteString(fmt.Sprintf("%s: %s", m.Role, m.Content))
		if err != nil {
			return err
		}
	}

	fmt.Println("history saved to", filename)

	return nil
}
