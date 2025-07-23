package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"

	"golings/internal/exercises"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var (
	success  = color.New(color.FgGreen).SprintFunc()
	errColor = color.New(color.FgRed).SprintFunc()
	warning  = color.New(color.FgYellow).SprintFunc()
	info     = color.New(color.FgCyan).SprintFunc()
)

func verifyExercise(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("please specify an exercise to verify")
	}

	exerciseName := c.Args().First()
	ex, err := findExercise(exerciseName)
	if err != nil {
		return err
	}

	// Run tests in the exercise directory
	cmd := exec.Command("go", "test")
	cmd.Dir = filepath.Dir(ex.Path)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		fmt.Println(errColor("Tests failed:"))
		fmt.Println(errColor(out.String()))
		return fmt.Errorf(errColor("exercise %q failed", exerciseName))
	}

	// If tests passed, mark exercise as complete
	if err := exercises.MarkComplete(exerciseName); err != nil {
		return fmt.Errorf(errColor("failed to mark exercise as complete: %w", err))
	}

	fmt.Println(success("\n✅ Exercise passed! Marked as complete."))
	return nil
}

func findExercise(name string) (*exercises.Exercise, error) {
	exList, err := exercises.List()
	if err != nil {
		return nil, fmt.Errorf("failed to list exercises: %w", err)
	}

	for _, ex := range exList {
		if ex.Name == name {
			return &ex, nil
		}
	}
	return nil, fmt.Errorf("exercise %q not found", name)
}
