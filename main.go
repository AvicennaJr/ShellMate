package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	URL = "https://api.openai.com/v1/chat/completions"
)

type (
	Prompt struct {
		Model    string     `json:"model"`
		Messages []Messages `json:"messages"`
	}
	Messages struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}

	Response struct {
		ID      string    `json:"id"`
		Object  string    `json:"object"`
		Created int       `json:"created"`
		Model   string    `json:"model"`
		Usage   Usage     `json:"usage"`
		Choices []Choices `json:"choices"`
	}
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	}
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
	Choices struct {
		Message      Message `json:"message"`
		FinishReason string  `json:"finish_reason"`
		Index        int     `json:"index"`
	}
)

func main() {
	fmt.Printf("\x1b[%dm%s\x1b[0m\n\n", 33, "Hi, I'm Shell Mate. How can I help you?")

	token := os.Getenv("OPENAPI_TOKEN")
	if token == "" {
		fmt.Println("I could not find the 'OPENAPI_TOKEN' value on your environment.")
		fmt.Printf("Please provide your token: ")
		fmt.Scanln(&token)
	}

	Start(os.Stdin, os.Stdout, token)
}

func Start(in io.Reader, out io.Writer, token string) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Print("🤖 >> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			fmt.Println("\nSee you soon! 👋")
			os.Exit(0)
		}

		var prompt Prompt

		prompt.Model = "gpt-3.5-turbo"
		prompt.Messages = make([]Messages, 1)
		prompt.Messages[0].Role = "user"
		prompt.Messages[0].Content = line

		jsonData, err := json.Marshal(prompt)
		if err != nil {
			fmt.Printf("\x1b[%dm%s\x1b[0m\n\n%s\n", 31, "Sorry, something went wrong 😔", err.Error())
			os.Exit(1)
		}

		req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("\x1b[%dm%s\x1b[0m\n\n%s\n", 31, "Sorry, something went wrong 😔", err.Error())
			os.Exit(1)
		}

		req.Header.Set("Content-Type", "application/json")
		bearer := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", bearer)

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("\x1b[%dm%s\x1b[0m\n\n%s\n", 31, "Sorry, something went wrong 😔", err.Error())
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("\x1b[%dm%s\x1b[0m\n\n%s\n", 31, "Sorry, something went wrong 😔", err.Error())
			os.Exit(1)
		}

		if resp.StatusCode == 200 {
			var response Response
			if err := json.Unmarshal(body, &response); err != nil {
				fmt.Printf("\x1b[%dm%s\x1b[0m\n\n%s\n", 31, "Sorry, something went wrong 😔", err.Error())
				os.Exit(1)
			}
			responseMessage := strings.Replace(strings.TrimSpace(response.Choices[0].Message.Content), "As an AI language model,", "You do know I'm just an AI right?", 1)
			io.WriteString(out, fmt.Sprintf("\n\x1b[%dm%s\x1b[0m\n", 32, responseMessage))
			io.WriteString(out, "\n")
		} else {
			fmt.Printf("\x1b[%dm%s\x1b[0m\n", 31, "Sorry, something went wrong 😔.")
			fmt.Println(string(body))
			os.Exit(1)
		}
	}
}
