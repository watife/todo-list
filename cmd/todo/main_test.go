package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	fileName = ".todo.json"
	binary   = "todo"
)

func TestMain(m *testing.M) {
	fmt.Println("Building binary")

	if runtime.GOOS == "windows" {
		binary = binary + ".exe"
	}

	build := exec.Command("go", "build", "-o", binary)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to build binary for %s: %v\n", binary, err)
		os.Exit(1)
	}

	fmt.Println("Running tests")
	result := m.Run()

	fmt.Println("Cleaning up")
	os.Remove(binary)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	cmdPath := filepath.Join(dir, binary)

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, strings.Split(task, "")...)

		if err := cmd.Run(); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		expected := task + "\n"
		if expected != string(out) {
			t.Errorf("Expected %q, got %q", expected, string(out))
		}

	})
}
