package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/denis-zakharov/gotodo/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Cannot create a task:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your task list with DB ID %d\n", task, id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
