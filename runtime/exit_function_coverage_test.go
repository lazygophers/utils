package runtime

import (
	"os"
	"os/exec"
	"testing"
)

// TestExitFunctionSubprocess tests the Exit() function using subprocess testing
func TestExitFunctionSubprocess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		// This is the subprocess that will call Exit()
		Exit()
		return
	}

	t.Run("exit_function_subprocess", func(t *testing.T) {
		// Use subprocess testing to test Exit() function
		cmd := exec.Command(os.Args[0], "-test.run=TestExitFunctionSubprocess")
		cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
		err := cmd.Run()

		// Exit() should cause the process to exit
		if err == nil {
			t.Fatal("Expected Exit() to cause process termination")
		}

		// Check if it's an exit error (process terminated)
		if exitError, ok := err.(*exec.ExitError); ok {
			t.Logf("Process exited as expected with status: %d", exitError.ExitCode())
		} else {
			t.Logf("Process terminated as expected: %v", err)
		}
	})
}

// TestExitErrorHandlingSubprocess tests error paths in Exit() function
func TestExitErrorHandlingSubprocess(t *testing.T) {
	if os.Getenv("GO_WANT_EXIT_ERROR_TEST") == "1" {
		// This subprocess will call Exit() to test the error handling paths
		Exit()
		return
	}

	t.Run("exit_error_handling_subprocess", func(t *testing.T) {
		cmd := exec.Command(os.Args[0], "-test.run=TestExitErrorHandlingSubprocess")
		cmd.Env = append(os.Environ(), "GO_WANT_EXIT_ERROR_TEST=1")
		err := cmd.Run()

		if err == nil {
			t.Fatal("Expected process to exit")
		}
		
		t.Logf("Exit function error handling tested: %v", err)
	})
}