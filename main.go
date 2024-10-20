package main

import (
	"flag"
	"log"
	"os"

	"github.com/Shobhit-Nagpal/trackr/internal/db"
	"github.com/Shobhit-Nagpal/trackr/internal/trackr/add"
	"github.com/Shobhit-Nagpal/trackr/internal/trackr/cmd"
	"github.com/Shobhit-Nagpal/trackr/internal/trackr/list"
	"github.com/Shobhit-Nagpal/trackr/internal/trackr/remove"
)

func main() {

	err := db.InitDB()
	if err != nil {
		log.Fatalf("DB INIT ERROR: ", err.Error())
	}

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	viewCmd := flag.NewFlagSet("view", flag.ExitOnError)

	if len(os.Args) < 2 {
		cmd.Render()
		return
	}

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		list.Render()
	case "add":
		addCmd.Parse(os.Args[2:])
		add.Render()
	case "remove":
		removeCmd.Parse(os.Args[2:])
		if len(os.Args) < 3 {
			remove.Render()
		} else {
			remove.Remove(os.Args[2])
		}
	case "view":
		viewCmd.Parse(os.Args[2:])
		log.Println("Add remove subcommand")
	default:
		log.Println("Command not recognized")
	}
}
