package commands

import (
	"MultiGit/log"
	"MultiGit/types"
	"bytes"
	"errors"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

func RunGitCommand(repoPath string, cmdArgs ...string) (string, error) {
	args := append([]string{"-C", repoPath}, cmdArgs...)
	cmd := exec.Command("git", args...)
	var outBuff, errBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	err := cmd.Run()
	if err != nil || errBuff.Len() > 0 {
		var resultError error
		if err != nil {
			resultError = err
		} else {
			resultError = errors.New(errBuff.String())
		}
		return "", resultError
	}

	return outBuff.String(), nil
}

func RunCustomCommand(repoPath string, command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = repoPath

	var outBuff, errBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	if err := cmd.Start(); err != nil {
		return "", err
	}

	// Wait asynchronously for the command to complete
	if err := cmd.Wait(); err != nil {
		stderrLines := strings.Split(errBuff.String(), "\n")
		var filteredErrors []string

		for _, line := range stderrLines {
			if !strings.Contains(line, "npm warn") && line != "" {
				filteredErrors = append(filteredErrors, line)
			}
		}

		if len(filteredErrors) > 0 {
			return "", errors.New(strings.Join(filteredErrors, "\n"))
		}

		return "", err
	}

	return outBuff.String(), nil
}

func ProcessReposWithProgress(repos *[]types.Repo, action func(repo types.Repo) (err error), limited bool) {
	var wg sync.WaitGroup
	var semaphore chan struct{}

	if limited {
		maxConcurrent := runtime.NumCPU() // Set a limit to the number of concurrent executions
		semaphore = make(chan struct{}, maxConcurrent)
	}

	errorChan := make(chan string, len(*repos))

	// Initialize the progress bar
	progressBar := progressbar.NewOptions(len(*repos),
		progressbar.OptionSetDescription("Processing repositories..."),
		progressbar.OptionShowCount(),
		progressbar.OptionSetElapsedTime(true),
	)

	for _, repo := range *repos {
		wg.Add(1)
		if limited {
			semaphore <- struct{}{} // Block if max concurrency is reached
		}

		go func(repo types.Repo) {
			defer wg.Done()
			if limited {
				defer func() { <-semaphore }() // Release the slot
			}

			err := action(repo)
			if err != nil {
				errorChan <- err.Error()
				_ = progressBar.Add(1)
				return
			}

			_ = progressBar.Add(1)
		}(repo)
	}

	wg.Wait()
	_ = progressBar.Finish()
	close(errorChan)

	fmt.Println("")

	// Print all errors after completion
	for err := range errorChan {
		log.Error(err)
	}

	fmt.Println("Finished")
}
