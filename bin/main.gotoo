package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/yolo-operator/yolo-operator/pkg/parser"
)

const (
	responseStyle = `. Answer using the exact following format (reponse with the exact same format as below):
YAML file content: COMPLETE HERE
File name: COMPLETE HERE
Command to deploy the file: COMPLETE HERE
Explanation: COMPLETE HERE
`
)

func main() {
	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Deploy a CronJob that will delete pod randomly in the cluster every 2hours. We have access to an image called kubikubectl that has the kubectl binary in it" + responseStyle,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println("full resp====>\n", resp.Choices[0].Message.Content)

	fmt.Println("-------")
	gpt3Response, err := parser.ParseGPT3Response(resp.Choices[0].Message.Content)
	if err != nil {
		fmt.Println("error parsing GPT3 response:", err)
		return
	}
	fmt.Println("gpt3Response====>\n")
	b, _ := json.MarshalIndent(gpt3Response, "", "    ")
	fmt.Println(string(b))

	fmt.Println("gpt3Response SANI====>\n")
	gpt3Response.Sanitize()
	b, _ = json.MarshalIndent(gpt3Response, "", "    ")
	fmt.Println(string(b))

}
