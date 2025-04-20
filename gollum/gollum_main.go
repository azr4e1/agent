package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	// "github.com/anthropics/anthropic-sdk-go"
	"github.com/azr4e1/gollum"
	"github.com/azr4e1/gollum/message"
)

func main() {
	// client := anthropic.NewClient()
	key := os.Getenv("OPENAI_API_KEY")
	client, err := gollum.NewClient(gollum.WithProvider(gollum.OPENAI), gollum.WithAPIKey(key))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	scanner := bufio.NewScanner(os.Stdin)
	getUserMessage := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}
		return scanner.Text(), true
	}

	agent := NewAgent(&client, getUserMessage)
	err = agent.Run(context.TODO())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func NewAgent(client *gollum.LLMClient, getUserMessage func() (string, bool)) *Agent {
	return &Agent{
		client:         client,
		getUserMessage: getUserMessage,
	}
}

type Agent struct {
	client         *gollum.LLMClient
	getUserMessage func() (string, bool)
}

func (a *Agent) Run(ctx context.Context) error {
	// conversation := []anthropic.MessageParam{}
	conversation := message.NewChat()

	fmt.Println("Chat with OpenAI (use 'ctrl-c' to quit)")

	for {
		fmt.Print("\u001b[94mYou\u001b[0m: ")
		userInput, ok := a.getUserMessage()
		if !ok {
			break
		}

		// userMessage := anthropic.NewUserMessage(anthropic.NewTextBlock(userInput))
		userMessage := message.UserMessage(userInput)
		// conversation = append(conversation, userMessage)
		conversation.Add(userMessage)

		response, err := a.runInference(ctx, conversation)
		if err != nil {
			return err
		}
		// conversation = append(conversation, message.ToParam())
		conversation.Add(message.AssistantMessage(response.Content()))

		// for _, content := range message.Content {
		// 	switch content.Type {
		// 	case "text":
		// 	fmt.Printf("\u001b[93mClaude\u001b[0m: %s\n", content.Text)
		// }
		// }
		fmt.Printf("\u001b[93mOpenAI\u001b[0m: %s\n", response.Content())
	}

	return nil
}

func (a *Agent) runInference(ctx context.Context, conversation message.Chat) (gollum.CompletionResponse, error) {
	// message, err := a.client.Messages.New(ctx, anthropic.MessageNewParams{
	// 	Model:     anthropic.ModelClaude3_7SonnetLatest,
	// 	MaxTokens: int64(1024),
	// 	Messages:  conversation,
	// })
	_, response, err := a.client.Complete(gollum.WithChat(conversation), gollum.WithMaxCompletionTokens(1024), gollum.WithModel("gpt-4o"))
	return response, err
}
