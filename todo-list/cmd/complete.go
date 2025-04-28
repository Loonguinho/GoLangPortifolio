package cmd

import (
	"fmt"
	"todo-list/config"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the ID of the task to mark as complete.")
			return
		}
		id := args[0]
		tasks, err := config.GetTasks()
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}
		found := false
		for i, task := range tasks {
			if task.ID == id {
				found = true
				tasks[i].IsComplete = true
				break
			}
		}
		if !found {
			fmt.Println("Task not found.")
			return
		}
		err = config.WriteTasks(tasks)
		if err != nil {
			fmt.Println("Error writing tasks:", err)
			return
		}
		fmt.Println("Task marked as complete.")
	},
}