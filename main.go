package main

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/yolo-operator/yolo-operator/pkg/parser"
)

const (
	responseStyle = `. Answer using the following format:
yaml file content: (if applicable)
command to run for the file: (if applicable)
explanation: (if applicable)
`
)

func main() {
	client := openai.NewClient("TTTTT")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Deploy a CronJob that will delete pod randomly in the cluster every 2hours" + responseStyle,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println("full resp====>\n", resp.Choices[0].Message.Content)

	parsed, err := parser.ParseGPT3Response(resp.Choices[0].Message.Content)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp yaml file", parsed.YamlFile)
	fmt.Println("resp file name", parsed.FileName)
	fmt.Println("resp command to run", parsed.CommandToRun)
	fmt.Println("resp explanation", parsed.Explanation)
}
