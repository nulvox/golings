package main

import (
	"fmt"
	"os"

	"golings/internal/exercises"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "golings",
		Usage: "Interactive Go exercises",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all available exercises",
				Action:  listExercises,
			},
			{
				Name:    "verify",
				Aliases: []string{"v"},
				Usage:   "Verify a specific exercise",
				Action:  verifyExercise,
			},
			{
				Name:    "progress",
				Aliases: []string{"p"},
				Usage:   "Show progress through exercises",
				Action:  showProgress,
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run a specific exercise",
				Action:  runExercise,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listExercises(c *cli.Context) error {
	fmt.Println("Available exercises:")

	exList, err := exercises.List()
	if err != nil {
		return fmt.Errorf("failed to list exercises: %w", err)
	}

	for _, ex := range exList {
		status := " "
		if ex.Completed {
			status = "✔"
		}
		fmt.Printf("[%s] %s - %s\n\t%s\n", status, ex.Name, ex.Description, ex.Path)
	}
	return nil
}

func showProgress(c *cli.Context) error {
	exList, err := exercises.List()
	if err != nil {
		return fmt.Errorf("failed to list exercises: %w", err)
	}

	completed := 0
	for _, ex := range exList {
		if ex.Completed {
			completed++
		}
	}

	percent := float64(completed) / float64(len(exList)) * 100
	fmt.Printf("Progress: %d/%d (%0.1f%%)\n", completed, len(exList), percent)

	fmt.Println("\nRecently completed:")
	for _, ex := range exList {
		if ex.Completed {
			fmt.Printf("✔ %s - %s\n", ex.Name, ex.Description)
		}
	}

	return nil
}

func runExercise(c *cli.Context) error {
	fmt.Println("Running exercise...")
	return nil
}
