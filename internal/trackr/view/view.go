package view

import (
	"fmt"
	"log"

	"github.com/Shobhit-Nagpal/trackr/internal/db"
	"github.com/charmbracelet/glamour"
)

func Render(name string) {
	project := db.GetProject(name)
	out, err := glamour.Render(project, "dark")
	if err != nil {
		log.Fatalf("Error rendering project: %s", err.Error())
	}
	fmt.Print(out)
}
