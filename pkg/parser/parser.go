package parser

import (
	"fmt"
	"regexp"
	"strings"
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

const (
	YamlFileToken     = "YAML file content:"
	FileNameToken     = "File name:"
	CommandToRunToken = "Command to deploy the file:"
	ExplanationToken  = "Explanation:"
)

// ParseGPT3Response parses the response from GPT-3 and returns the relevant
// information.
// Responses are in the format of
//
//	yaml file: (if applicable)
//	file name: (if applicable)
//	command to run:
//	explanation:
func ParseGPT3Response(input string) (*GPT3Response, error) {
	// We're trying to parse a string that looks like this:
	// yaml file: (if applicable)
	// file name: (if applicable)
	// command to run:
	// explanation:

	// While this sounds like a job for a regex, it's actually not
	// that simple. The problem is that the input is not always in the
	// same format. Sometimes new lines are used, sometimes not. Sometimes
	// there are spaces after the colon, sometimes not. Sometimes there are
	// spaces before the colon, sometimes not. Sometimes there are spaces
	// after the token, sometimes not.

	// We avoid using regex for the sake of solving this problem in a
	// faster way. We use a simple state machine instead.

	// We can get the indexes of each token in the input string. We can
	// then use those indexes to extract the relevant information.

	// First let's remove any trailing new lines and spaces.
	input = trimTrailingNewLinesAndSpaces(input)

	YamlFileTokenIndex := strings.Index(input, YamlFileToken)
	FileNameTokenIndex := strings.Index(input, FileNameToken)
	CommandToRunTokenIndex := strings.Index(input, CommandToRunToken)
	ExplanationTokenIndex := strings.Index(input, ExplanationToken)

	// We need to make sure that all tokens are present.
	if YamlFileTokenIndex == -1 || FileNameTokenIndex == -1 || CommandToRunTokenIndex == -1 || ExplanationTokenIndex == -1 {
		return nil, fmt.Errorf("unparsable text: missing token index ?") // todo: improve error message
	}

	// We need to make sure that the tokens are in the right order.
	if YamlFileTokenIndex > FileNameTokenIndex ||
		FileNameTokenIndex > CommandToRunTokenIndex ||
		CommandToRunTokenIndex > ExplanationTokenIndex {
		return nil, fmt.Errorf("unparsable text: tokens are not in the right order") // todo: improve error message
	}

	// We need to query the text between each token.

	yamlFileContent := input[YamlFileTokenIndex+len(YamlFileToken) : FileNameTokenIndex]
	fileNameContent := input[FileNameTokenIndex+len(FileNameToken) : CommandToRunTokenIndex]
	commandToRunContent := input[CommandToRunTokenIndex+len(CommandToRunToken) : ExplanationTokenIndex]
	explanationContent := input[ExplanationTokenIndex+len(ExplanationToken):]

	return &GPT3Response{
		YamlFile:     yamlFileContent,
		FileName:     fileNameContent,
		CommandToRun: commandToRunContent,
		Explanation:  explanationContent,
	}, nil
}

func trimTrailingNewLinesAndSpaces(input string) string {
	// We remove any trailing new lines and spaces.
	for input[len(input)-1] == '\n' || input[len(input)-1] == ' ' {
		input = input[:len(input)-1]
	}
	return input
}

func (resp *GPT3Response) Sanitize() {
	// Remove any trailing new lines and spaces.
	resp.YamlFile = strings.TrimSpace(resp.YamlFile)
	resp.FileName = strings.TrimSpace(resp.FileName)
	resp.CommandToRun = strings.TrimSpace(resp.CommandToRun)
	resp.Explanation = strings.TrimSpace(resp.Explanation)

	resp.YamlFile = strings.ReplaceAll(resp.YamlFile, "```yaml\n", "")
	resp.YamlFile = strings.ReplaceAll(resp.YamlFile, "```", "")

	resp.CommandToRun = strings.ReplaceAll(resp.CommandToRun, "```bash\n", "")
	resp.CommandToRun = strings.ReplaceAll(resp.CommandToRun, "```", "")
	resp.CommandToRun = strings.ReplaceAll(resp.CommandToRun, "\n", "")
}
