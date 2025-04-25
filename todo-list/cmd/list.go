package cmd

import (
	"fmt"
	"os"
	"tasks/config"
	"tasks/models"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var showAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing tasks...")
		tasks, err := config.GetTasks()
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}
		printTasks(tasks, showAll)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks, including completed ones")
	rootCmd.AddCommand(listCmd)
}

func printTasks(tasks []models.Task, showAll bool) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	if showAll{
		fmt.Fprintln(writer, "ID\tDescription\tCreated At\tIs Complete")
	} else {
		fmt.Fprintln(writer, "ID\tDescription\tCreated At")
	}

	for _, task := range tasks {
		if !showAll &&	task.IsComplete {
			continue
		}

		created := timediff.TimeDiff(task.CreatedAt)//, timediff.WithLocale("pt-BR"))
		
		if showAll {
			fmt.Fprintf(writer, "%s\t%s\t%s\t%t\n", task.ID, task.Description, created, task.IsComplete)
		} else {
			fmt.Fprintf(writer, "%s\t%s\t%s\n", task.ID, task.Description, created)
		}
}

	writer.Flush()
}