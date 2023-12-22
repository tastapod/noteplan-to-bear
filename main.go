package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	startDir := filepath.Join(homeDir,
		"Library/Containers/co.noteplan.NotePlan3",
		"Data/Library/Application Support/co.noteplan.NotePlan3/Notes")

	err := os.Chdir(startDir)
	if err != nil {
		exitError(err)
	}

	err = filepath.WalkDir(".", migrateFunc)
	if err != nil {
		exitError(err)
	}
}

func exitError(err error) {
	fmt.Printf("Error: %v", err)
	os.Exit(1)
}

func migrateFunc(path string, d fs.DirEntry, fileErr error) error {
	// escalate errors
	if fileErr != nil {
		return fileErr
	}

	// skip NotePlan system dirs (start with @)
	if d.IsDir() {
		if strings.HasPrefix(d.Name(), "@") {
			return filepath.SkipDir
		}
		return nil
	}

	// skip dot-files
	if strings.HasPrefix(d.Name(), ".") {
		return nil
	}

	// migrate the note
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// note content
	content := string(contentBytes)

	sourceTag := "#_noteplan/" + strings.ToLower(filepath.Dir(path))
	content = strings.Replace(content, "\n", "\n"+sourceTag+"\n", 1)

	err = clipboard.Init()
	if err != nil {
		exitError(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(content))

	url := "bear://x-callback-url/create?clipboard=yes"

	cmdErr := exec.Command("open", url).Run()
	if cmdErr != nil {
		return cmdErr
	}

	return nil
}
