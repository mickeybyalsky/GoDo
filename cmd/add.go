/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"GoDo/task"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Created     time.Time
	Completed   bool
	CompletedAt *time.Time
}

var taskName string
var taskDetails string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long: `The add command allows you to add a new task to your Todo app.
You can specify details like the name and description of the task.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		if taskName == "" {
			fmt.Println("Error: Task name is required")
			cmd.Help()
			os.Exit(1)
		}

		now := time.Now()
		newTask := task.Task{
			ID:          0,
			Title:       taskName,
			Description: taskDetails,
			Created:     now,
			Completed:   false,
			CompletedAt: nil,
		}

		taskJSON, err := json.MarshalIndent(newTask, "", " ")
		if err != nil {
			fmt.Println("Encountered an error:", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully!")
		// fmt.Printf("Name: %s\n", taskName)
		fmt.Println(string(taskJSON))
		// fmt.Printf("Description: %s\n", taskDetails)
		// fmt.Printf("Created Time: %s\n", time.Now())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&taskName, "name", "n", "", "Task Name (required)")
	addCmd.Flags().StringVarP(&taskDetails, "details", "d", "", "Task Description")
	addCmd.MarkFlagRequired("name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
