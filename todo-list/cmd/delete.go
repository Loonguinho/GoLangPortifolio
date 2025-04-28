package cmd

import (
	"fmt"
	"todo-list/config"
	"todo-list/models"

	"github.com/spf13/cobra"
)
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the ID of the task to delete.")
			return
		}
		var updatedTasks []models.Task
		found := false
		id := args[0]
		tasks, err := config.GetTasks()
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}
		for _, task := range tasks {
			if task.ID == id {
				found = true
				continue
			}
			updatedTasks = append(updatedTasks, task)
		}
		if !found {
			fmt.Println("Task not found.")
			return
		}

		err = config.WriteTasks(updatedTasks)
		if err != nil {
			fmt.Println("Error writing tasks:", err)
			return
		}
		fmt.Println("Task deleted.")
	},
	}