package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func readInput(input chan string) {
	for {
		var line string
		_, err := fmt.Scanf("%s\n", &line)
		if err != nil {
			panic(err)
		}
		input <- line
	}
}

func main() {
	hearCmd := exec.Command("hear")
	stdout, _ := hearCmd.StdoutPipe()

	// listen to text on stdin to start and stop recording
	// and print the text to stdout

	input := make(chan string)
	go readInput(input)
	for i := range input {
		switch i {
		case "start":
			err := hearCmd.Start()
			if err != nil {
				fmt.Println("error starting the command: ", err)
				close(input)
			}
		case "stop":
			err := hearCmd.Process.Kill()
			if err != nil {
				fmt.Println("error stopping the command: ", err)
			}
			close(input)
		}
	}

	allOutput := []string{}
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		allOutput = append(allOutput, m)
	}
	fmt.Println(getFullSentence(allOutput))
	myAudioInputFile := "my-audio-input.yaml"
	_ = writeOutputToKubernetesCommandFile(myAudioInputFile, getFullSentence(allOutput))
}

func getFullSentence(allOutput []string) string {
	// traverse the array in reverse order and find the first
	// element that matches the last element.
	lastElement := allOutput[len(allOutput)-1]
	for i := len(allOutput) - 2; i >= 0; i-- {
		if allOutput[i] == lastElement {
			return strings.Join(allOutput[i+1:], " ") + "."
		}
	}
	return "ERROR"
}

// writeOutputToKubernetesCommandFile writes the output to a kubernetes command file
// that can be used to run the command in kubernetes.
/*
apiVersion: batch.yolo.ahilaly.dev/v1alpha1
kind: Command
metadata:
	name: my-command
spec:
	input: INPUT
*/
func writeOutputToKubernetesCommandFile(filename, input string) error {
	content := fmt.Sprintf(`apiVersion: batch.yolo.ahilaly.dev/v1alpha1
kind: Command
metadata:
	name: my-command
spec:
	input: %s`, input)
	return os.WriteFile(filename, []byte(content), 0644)
}
