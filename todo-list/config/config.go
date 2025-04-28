package config

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"todo-list/models"
)

const TaskFile = "data/tasks.csv"

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	file, err := os.Open(TaskFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		id := record[0]
		desc := record[1]
		createdAt, _ := time.Parse(time.RFC3339, record[2])
		isComplete, _ := strconv.ParseBool(record[3])

		tasks = append(tasks, models.Task{
			ID:          id,
			Description: desc,
			CreatedAt:   createdAt,
			IsComplete:  isComplete,
		})
	}

	return tasks, nil

}

func WriteTasks(tasks []models.Task) error {
	file, err := os.Create(TaskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		err := writer.Write([]string{
			task.ID,
			task.Description,
			task.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(task.IsComplete),
		})
		if err != nil {
			return err
		}
	}
	return nil
}