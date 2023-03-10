package chatgo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/xavierxcn/chatgo/utils"
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

// Init initializes the robot
func (r *Robot) Init() {
	if r.token == "" {
		panic("token is empty")
	}

	r.messages = []*message{
		{
			Role:    RoleSystem,
			Content: fmt.Sprintf("hello, you are %s, you are a helpful assistant. ", r.Name()),
		},
	}

	_, err := r.tell()
	if err != nil {
		panic(err)
	}
}

// SetName sets the robot name
func (r *Robot) SetName(name string) *Robot {
	r.name = name
	return r
}

// Name gets the robot name
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

// SetMessagesFromFile sets the robot messages
func (r *Robot) SetMessagesFromFile(path string) *Robot {
	m, err := utils.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var messages []*message
	err = json.Unmarshal([]byte(m), &messages)
	if err != nil {
		panic(err)
	}

	r.messages = messages
	return r
}

// Replay 回放之前的聊天记录
func (r *Robot) Replay() {
	for _, m := range r.messages[1:] {
		role := m.Role
		if m.Role == RoleAssistant {
			role = r.name
		}
		fmt.Printf("%s: %s\n", role, m.Content)
	}
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

// TellStream tells the robot something and returns a stream
func (r *Robot) TellStream(sentence string) (<-chan string, error) {
	var err error

	r.messages = append(r.messages, &message{
		Role:    RoleUser,
		Content: strings.TrimSpace(sentence),
	})

	stream, err := r.tellStream()
	if err != nil {
		return nil, err
	}

	result := make(chan string)

	message := &message{
		Role:    RoleAssistant,
		Content: "",
	}

	go func() {
		defer func() {
			close(result)
		}()
		for {

			rsp := <-stream
			if rsp == nil {
				break
			}

			for _, choice := range rsp.Choices {
				result <- choice.Delta.Content
				message.Content += choice.Delta.Content
			}
		}
	}()

	r.messages = append(r.messages, message)
	return result, nil
}

// Save saves the messages to file
func (r *Robot) Save(path string) error {
	// 将messages保存到HistoryPath
	f, err := utils.CreateAndOpenFile(path)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	message, err := json.Marshal(r.messages)
	if err != nil {
		return err
	}

	_, err = f.Write(message)
	if err != nil {
		return err
	}

	fmt.Println("history saved to", path)

	return nil
}
