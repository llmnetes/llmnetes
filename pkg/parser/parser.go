package parser

import (
	"fmt"
	"regexp"

	"github.com/ghodss/yaml"
)

var (
	re = regexp.MustCompile(`yaml file: (.*)
	file name: (.*)
	command to run: (.*)
	explanation: (.*)`)
)

func NewParser() *Parser {
	return &Parser{}
}

type Parser struct{}

/*
	yaml file: (if applicable)
	file name: (if applicable)
	command to run:
	explanation:
*/

type GPT3Response struct {
	YamlFile     string `json:"yaml_file"`
	FileName     string `json:"file_name"`
	CommandToRun string `json:"command_to_run"`
	Explanation  string `json:"explanation"`
}

// ParseGPT3Response parses the response from GPT-3 and returns the relevant
// information.
// Responses are in the format of
//
//	yaml file: (if applicable)
//	file name: (if applicable)
//	command to run:
//	explanation:
func ParseGPT3Response(input string) (*GPT3Response, error) {
	fmt.Println("===> input", input)
	// remove all trailing spaces but keep the newlines
	// input = strings.ReplaceAll(input, " ", "")

	var resp GPT3Response
	err := yaml.Unmarshal([]byte(input), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
