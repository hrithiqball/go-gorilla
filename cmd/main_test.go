package main

import (
	"bytes"
	"log"
	"os/exec"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestMainFunction(t *testing.T) {
	cmd := exec.Command("go", "run", "cmd/main.go")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to execute main.go: %v", err)
	}

	expected := "Expected output from your app\n"
	if out.String() != expected {
		t.Errorf("Unexpected output: got %v, want %v", out.String(), expected)
	}
}
