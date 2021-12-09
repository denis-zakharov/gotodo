package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/denis-zakharov/gotodo/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to pase the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.ReadTasks()
		if err != nil {
			fmt.Println("Failed to read tasks:", err.Error())
			os.Exit(1)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Println("Failed to delete task:", err.Error())
			} else {
				fmt.Printf("Completed %v\n", task)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
