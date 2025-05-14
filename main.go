package main

import (
	"github.com/916RasnitaRadu/todo-cli/cmd"
)

func main() {
	cmd.Execute()
}

/*
	TODO:
		- SHOULD DO A VIEW PACKAGE REFACTOR (view = {cli + cmd + common})
		- refactor errors management
		- style interface
		- display date using mergestat/timediff

1. A task must have: ID, Task (name), Created At, Done x
2. Must support commands x
3. Must be a runnable application x
4. Must support CRUD (create, list, update (write done/not done), delete) x
5. Must support:
	- file storage (CSV) x
	- SQLITE
	- (choose the storage through env vars)

------ packages
encoding/csv
text/tabwriter for writing out tab aligned output
github.com/spf13/cobra for the command line interface
github.com/mergestat/timediff for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)
*/
