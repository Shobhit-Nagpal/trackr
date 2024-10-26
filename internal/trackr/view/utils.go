package view

import (
	"log"

	"github.com/Shobhit-Nagpal/trackr/internal/db"
	"github.com/charmbracelet/glamour"
)

func GetRenderedMarkdown(name string) string {
	project := db.GetProject(name)
	out, err := glamour.Render(project, "dark")
	if err != nil {
		log.Fatalf("Error rendering project: %s", err.Error())
	}

	return out
}
