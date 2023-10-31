package k8s

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

const (
	// defaultRootDir is the default directory where files will be written.
	defaultRootDir = "/tmp/yolo-operator"
)

// NewShellAccess creates a new shellAccess object.
func NewShellAcess(rootDir string) *ShellAccess {
	if rootDir == "" {
		rootDir = defaultRootDir
	}
	return &ShellAccess{
		rootDir: rootDir,
	}
}

type ShellAccess struct {
	rootDir string
	// TODO: support swapping out kubeconfigs
	kubeConfig string
}

// Commands usually need to apply some files to the cluster before they can be run.
// This method should be used to prepare those files.
// It will create a temporary directory, write the files to it, and return the path
// to the directory.
// It will return a cleanup function that should be called when the command is done
// running.
func (s *ShellAccess) PrepareFiles(files map[string]string) (string, func(), error) {
	randomDirName := randString(10)

	dir := s.rootDir + "/" + randomDirName
	// Create temp directory
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return "", func() {}, err
	}
	// Write files to it
	for fileName, fileContent := range files {
		f, err := os.Create(dir + "/" + fileName)
		if err != nil {
			return "", func() {}, err
		}
		_, err = f.WriteString(fileContent)
		if err != nil {
			return "", func() {}, err
		}
		err = f.Close()
		if err != nil {
			return "", func() {}, err
		}
	}

	// Create the cleanup function
	cleanup := func() {
		os.RemoveAll(randomDirName)
	}

	// Return path to directory and cleanup function
	return dir, cleanup, nil
}

// RunCommand runs a command on the shell and returns the output
// the function will use workDir as the working directory.
func (s *ShellAccess) RunCommand(workDir, command string) (string, error) {
	// Run command on shell
	fmt.Println("Running command", command, "in directory", workDir)
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = workDir
	stdout, err := cmd.CombinedOutput()
	fmt.Println("stdout=====>", string(stdout))
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}

// randString generates a random string of length n
func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
