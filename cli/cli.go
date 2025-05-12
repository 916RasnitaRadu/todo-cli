package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/916RasnitaRadu/todo-cli/service"
	"github.com/916RasnitaRadu/todo-cli/types"
)

type CLI struct {
	srv service.Service
}

func NewCLI(service service.Service) CLI {
	return CLI{
		srv: service,
	}
}

func printMenu() {
	fmt.Println(`
	1. View task list
	2. Add
	3. Delete
	4. Check
	x. Exit
	`)
}

func (cli *CLI) getNextId() int {
	var maxId int
	tasks, err := cli.srv.GetTasks()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return -1
	}

	for _, t := range tasks {
		if t.ID > maxId {
			maxId = t.ID
		}
	}

	return maxId + 1
}

func (cli *CLI) listTasks() {
	tasks, err := cli.srv.GetTasks()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	for _, t := range tasks {
		fmt.Printf("%d. %s	%v	%v\n", t.ID, t.Description, t.CreatedAt, t.Done)
	}
}

func (cli *CLI) createTask() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the task: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	newTask := types.Task{
		ID:          cli.getNextId(),
		Description: taskName,
		CreatedAt:   time.Now(),
	}

	if err := cli.srv.Create(newTask); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Println("New task created successfully!")
}

func (cli *CLI) deleteTask() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the task id: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	if err := cli.srv.Delete(id); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Println("Task deleted successfully!")
}

func (cli *CLI) updateTask() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the task id: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	if err := cli.srv.ChangeStatus(id); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Println("Task updated successfully!")
}

func (cli *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printMenu()

		fmt.Printf("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			cli.listTasks()
		case "2":
			cli.createTask()
		case "3":
			cli.deleteTask()
		case "4":
			cli.updateTask()
		case "x":
			fmt.Println("Goodbye! ;)")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}
