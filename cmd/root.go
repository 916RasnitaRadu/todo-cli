package cmd

import (
	"os"

	"github.com/916RasnitaRadu/todo-cli/repository"
	"github.com/916RasnitaRadu/todo-cli/service"
	"github.com/spf13/cobra"
)

func getService() service.Service {
	repo := repository.NewFileRepository("tasks.csv")
	return service.NewService(repo)
}

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A todo list that is very fancy",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toogle", "t", false, "Help message for toggle")
}
