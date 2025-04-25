package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"tasks/config"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a description for the task.")
			return
		}
		
		desc := args[0]
		if len(desc) < 10 {
			fmt.Println("Task description must be at least 10 characters long.")
			return
		}
		fmt.Println("Adding task:", desc)
		
		file, err := os.OpenFile(config.TaskFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		id := uuid.New().String()[:5]
		createdAt := time.Now()

		task := []string{
			id,
			desc,                              
			createdAt.Format(time.RFC3339),    
			"false",                           
		}

		if err := writer.Write(task); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("Task added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}