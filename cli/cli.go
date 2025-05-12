package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/916RasnitaRadu/todo-cli/service"
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

func (cli *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printMenu()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			cli.listTasks()
		case "2":
			// add()
		case "3":
			// delete()
		case "4":
			// edit()
		case "x":
			fmt.Println("Goodbye! ;)")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}
