package exercises

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Exercise struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Completed   bool   `json:"completed"`
}

var progressPath = filepath.Join("data", "progress.json")

func List() ([]Exercise, error) {
	var exercises []Exercise

	// Walk through the exercises directory
	err := filepath.Walk("exercises", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process directories that match the exercise pattern
		if info.IsDir() && strings.HasPrefix(filepath.Base(path), "0") {
			// Read the README.md for the description
			readmePath := filepath.Join(path, "README.md")
			desc := ""
			if data, err := os.ReadFile(readmePath); err == nil {
				// Get the first non-empty line of description
				for _, line := range strings.Split(string(data), "\n") {
					if strings.TrimSpace(line) != "" {
						desc = strings.TrimSpace(line)
						break
					}
				}
			}

			exercisePath := filepath.Join(path, "exercise.go")
			exercises = append(exercises, Exercise{
				Name:        filepath.Base(path),
				Description: desc,
				Path:        exercisePath,
				Completed:   false,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Sort exercises by name
	sort.Slice(exercises, func(i, j int) bool {
		return exercises[i].Name < exercises[j].Name
	})

	// Load progress and update completion status
	progress, err := LoadProgress()
	if err != nil {
		return nil, err
	}

	for i := range exercises {
		if completed, exists := progress[exercises[i].Name]; exists {
			exercises[i].Completed = completed
		}
	}

	return exercises, nil
}

func LoadProgress() (map[string]bool, error) {
	progress := make(map[string]bool)

	if _, err := os.Stat(progressPath); os.IsNotExist(err) {
		return progress, nil
	}

	data, err := os.ReadFile(progressPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &progress)
	return progress, err
}

func SaveProgress(progress map[string]bool) error {
	if err := os.MkdirAll(filepath.Dir(progressPath), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(progress, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(progressPath, data, 0644)
}

func MarkComplete(name string) error {
	progress, err := LoadProgress()
	if err != nil {
		return err
	}

	progress[name] = true
	return SaveProgress(progress)
}
