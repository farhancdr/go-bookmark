package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestMainExecution(t *testing.T) {
	// Create a temporary directory to simulate a clean environment
	tempDir, err := ioutil.TempDir("", "go-bookmark-test")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up the temporary directory

	// Change to the temporary directory to avoid affecting actual files
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("failed to change to temp directory: %v", err)
	}
	defer os.Chdir(originalDir) // Change back after the test

	// Capture stdout and stderr
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	os.Stdout = wOut
	os.Stderr = wErr

	// Restore stdout and stderr after the test
	defer func() {
		wOut.Close()
		wErr.Close()
		os.Stdout = oldStdout
		os.Stderr = oldStderr
	}()

	// Simulate running the main function (without arguments for a basic test)
	// You might want to simulate arguments for specific command tests,
	// but for main.go, a no-op run is often sufficient for a basic check.
	os.Args = []string{"go-bookmark"} // Simulate calling the executable without subcommands

	// Run the main function in a goroutine to allow capturing output
	done := make(chan struct{})
	go func() {
		main()
		close(done)
	}()

	// Wait for main to complete
	<-done

	// Read captured output
	stdoutBytes, _ := ioutil.ReadAll(rOut)
	stderrBytes, _ := ioutil.ReadAll(rErr)

	// In a typical application with no default action or a simple help message,
	// you might expect some output to stdout, and no errors to stderr.
	// This specific main.go will likely print the help message if no command is given.
	if len(stderrBytes) > 0 {
		t.Errorf("expected no errors to stderr, but got: %s", string(stderrBytes))
	}

	// You can add more specific checks based on your default root command behavior.
	// For example, if it prints a help message by default:
	if !bytes.Contains(stdoutBytes, []byte("A CLI for managing your bookmarks")) {
		t.Errorf("expected help message in stdout, got:\n%s", string(stdoutBytes))
	}
}
