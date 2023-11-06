package openai

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

const (
	// responseStyle is the style of the response we expect from the model.
	// TODO: make this configurable.
	responseStyle = `. Answer using the following YAML format:
	yaml file: (if applicable)
	file name: (if applicable)
	command to run:
	explanation:`

	responseStyle2 = `. Answer using the exact following format (reponse with the exact same format as below):
	YAML file content: COMPLETE HERE
	File name: COMPLETE HERE
	Command to deploy the file: COMPLETE HERE
	Explanation: COMPLETE HERE
`
)

// clientWraper is a wrapper around the openai.Client.
type clientWraper struct {
	client *openai.Client
}

// NewClientFromEnv creates a new client from the OPENAI_TOKEN environment variable.
func NewClientFromEnv() (*clientWraper, error) {
	token := os.Getenv("OPENAI_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("OPENAI_TOKEN is not set")
	}
	client := &clientWraper{
		openai.NewClient(token),
	}
	return client, nil
}

// RunQuery runs a query against the GPT-3 model.
func (c *clientWraper) RunQuery(query string) (string, error) {
	fmt.Println("Running query")
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query + responseStyle2,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	// We pick the first choice, which is the most likely one.
	return resp.Choices[0].Message.Content, nil
}
