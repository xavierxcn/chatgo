// Package chatgo is a chatbot based on openai api
package chatgo

// TokenPath is the path of token
const TokenPath = "~/.chatgo/token"

// HistoryPath is the path of history
const HistoryPath = "~/.chatgo/%s.%s.history"

// defaultModel is the default openai model
const defaultModel = "gpt-3.5-turbo-0301"

// charURL is the openai api url
const charURL = "https://api.openai.com/v1/chat/completions"
