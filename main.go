package main

import (
	"github.com/916RasnitaRadu/todo-cli/cli"
	"github.com/916RasnitaRadu/todo-cli/repository"
	"github.com/916RasnitaRadu/todo-cli/service"
)

func main() {
	repo := repository.NewFileRepository("tasks.csv")
	srv := service.NewService(repo)
	cli := cli.NewCLI(srv)
	cli.Run()
}

/*
1. A task must have: ID, Task (name), Created At, Done
2. Must support commands
3. Must be a runnable application (?)
4. Must support CRUD (create, list, update (write done/not done), delete)
5. Must support file storage (CSV), and SQLITE (choose the storage through env vars)

------ packages
encoding/csv
text/tabwriter for writing out tab aligned output
github.com/spf13/cobra for the command line interface
github.com/mergestat/timediff for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)
*/
