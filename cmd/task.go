package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/916RasnitaRadu/todo-cli/cli"
	"github.com/916RasnitaRadu/todo-cli/types"
	"github.com/spf13/cobra"
)

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks",
	Run:   listTasks,
}

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Run:   addTask,
}

var deleteTaskCmd = &cobra.Command{
	Use:   "del",
	Short: "delete a task",
	Run:   deleteTask,
}

var updateTaskCmd = &cobra.Command{
	Use:   "upd",
	Short: "Change the status of a task",
	Run:   updateTask,
}

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run:   startServer,
}

func updateTask(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	srv := getService()
	if err := srv.ChangeStatus(id); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}

func deleteTask(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	srv := getService()
	if err := srv.Delete(id); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}

func addTask(cmd *cobra.Command, args []string) {
	taskName := strings.Join(args, " ")

	srv := getService()
	cli := cli.NewCLI(srv)

	newTask := types.Task{
		ID:          cli.GetNextId(),
		Description: taskName,
		CreatedAt:   time.Now(),
	}

	if err := srv.Create(newTask); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}

func listTasks(cmd *cobra.Command, args []string) {
	fmt.Println("Listing the tasks:")

	srv := getService()
	cli := cli.NewCLI(srv)

	cli.ListTasks()
}

func startServer(cmd *cobra.Command, args []string) {
	srv := getService()
	cli := cli.NewCLI(srv)
	cli.Run()
}

func init() {
	rootCmd.AddCommand(listTasksCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(deleteTaskCmd)
	rootCmd.AddCommand(updateTaskCmd)
}
