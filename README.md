# Trackr

A sleek command-line interface tool for tracking your projects with style. Built with Go and enhanced with modern TUI libraries for a beautiful terminal experience.

## Features

- 📋 List all your projects
- ➕ Add new projects
- 🔍 View detailed project information
- 🗑️ Remove projects
- 💅 Beautiful terminal UI powered by Bubbletea and Lipgloss

## Installation

```bash
go install github.com/Shobhit-Nagpal/trackr@latest
```

Or clone the repository and build from source:

```bash
git clone https://github.com/Shobhit-Nagpal/trackr.git
cd trackr
go build
```

## Usage

### Basic Commands

```bash
trackr list    # List all projects
trackr add     # Add a new project
trackr remove  # Remove a project
trackr view    # View project details
```

### Command Details

#### List Projects
```bash
trackr list
```
Displays a formatted list of all your tracked projects.

#### Add Project
```bash
trackr add
```
Launches an interactive prompt to add a new project with relevant details.

#### Remove Project
```bash
trackr remove [project-name]
```
Removes a project from tracking. If no project name is provided, shows an interactive list.

#### View Project
```bash
trackr view [project-name]
```
Shows detailed information about a specific project. If no project name is provided, shows an interactive list.

## Dependencies

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Glamour](https://github.com/charmbracelet/glamour) - Markdown rendering
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing-feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

[MIT License](LICENSE)

## Author

Shobhit Nagpal

## Acknowledgments

- The [Charm](https://charm.sh/) team for their amazing terminal UI libraries
- The Go community for inspiration and support
