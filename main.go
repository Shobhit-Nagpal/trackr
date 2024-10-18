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

	err := db.InitDB()
	if err != nil {
		log.Fatalf("DB INIT ERROR: ", err.Error())
	}


	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
  addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	if len(os.Args) < 2 {
		cmd.Render()
	}

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		list.Render()
	case "add":
    addCmd.Parse(os.Args[2:])
    add.Render()
	case "remove":
		log.Println("Add remove subcommand")
	case "view":
		log.Println("Add view subcommand")
  default:
    log.Println("Command not recognized")
    os.Exit(1)
	}
}
