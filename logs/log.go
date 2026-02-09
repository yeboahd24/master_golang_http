package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

var envLoadOnce sync.Once

func ensureEnvLoaded() {
	envLoadOnce.Do(func() {
		// Get current directory
		cwd, err := os.Getwd()
		if err != nil {
			cwd = "."
		}

		// Try current dir and parent dir
		paths := []string{
			filepath.Join(cwd, ".env"),
			filepath.Join(cwd, "..", ".env"),
		}

		for _, path := range paths {
			if err := godotenv.Load(path); err == nil {
				fmt.Println("Loaded .env from:", path)
				return
			}
		}

		fmt.Println("No .env file found, using system env vars")
	})
}
