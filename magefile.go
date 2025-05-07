//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Generate regenerates the protobuf code
func Generate() error {
	fmt.Println("Generating protobuf code...")

	// Ensure protoc is installed
	if err := exec.Command("protoc", "--version").Run(); err != nil {
		return fmt.Errorf("protoc not found: %v", err)
	}

	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	// Define the proto file path
	protoFile := filepath.Join(dir, "pkg", "pluginapi", "plugin.proto")

	// Run protoc command
	cmd := exec.Command("protoc",
		"--go_out=.", "--go_opt=paths=source_relative",
		"--go-grpc_out=.", "--go-grpc_opt=paths=source_relative",
		protoFile,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate protobuf code: %v", err)
	}

	fmt.Println("Protobuf code generated successfully!")
	return nil
}

// Clean removes generated files
func Clean() error {
	fmt.Println("Cleaning generated files...")

	// Add patterns of files to clean
	patterns := []string{
		"pkg/pluginapi/*.pb.go",
	}

	for _, pattern := range patterns {
		files, err := filepath.Glob(pattern)
		if err != nil {
			return fmt.Errorf("failed to glob pattern %s: %v", pattern, err)
		}

		for _, file := range files {
			if err := os.Remove(file); err != nil {
				return fmt.Errorf("failed to remove %s: %v", file, err)
			}
			fmt.Printf("Removed %s\n", file)
		}
	}

	return nil
}

// All runs all tasks
func All() error {
	if err := Clean(); err != nil {
		return err
	}
	return Generate()
}
