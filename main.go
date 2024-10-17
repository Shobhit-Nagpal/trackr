package main

import (
	"flag"
	"log"
	"os"

	"github.com/Shobhit-Nagpal/trackr/internal/db"
	"github.com/Shobhit-Nagpal/trackr/internal/trackr/add"
	"github.com/Shobhit-Nagpal/trackr/internal/trackr/cmd"
	"github.com/Shobhit-Nagpal/trackr/internal/trackr/list"
)

func main() {

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
  addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	if len(os.Args) < 2 {
		cmd.Initialize()
	}

	err := db.InitDB()
	if err != nil {
		log.Fatalf("DB INIT ERROR: ", err.Error())
	}

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		list.Initialize()
	case "add":
    addCmd.Parse(os.Args[2:])
    add.Initialize()
	case "remove":
		log.Println("Add remove subcommand")
	case "view":
		log.Println("Add view subcommand")
  default:
    log.Println("Command not recognized")
    os.Exit(1)
	}
}
