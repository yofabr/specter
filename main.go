package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type step int

const (
	stepLoading step = iota
	stepEnvironment
	stepUsername
	stepConfirm
	stepDone
)

type model struct {
	step          step
	loadingTick   int
	loadingTitles []string
	envSelected   int
	environments  []string
	username      textinput.Model
}

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FFAA"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFAA")).
			Bold(true)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))

	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#555555"))
)

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter your username"
	ti.Focus()

	return model{
		step:        stepLoading,
		loadingTick: 0,
		loadingTitles: []string{
			"Initializing Specter...",
			"Loading configuration...",
			"Preparing environment...",
			"Connecting to services...",
			"Almost ready...",
		},
		envSelected:  0,
		environments: []string{"Development", "Staging", "Production"},
		username:     ti,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, tick())
}

func tick() tea.Cmd {
	return tea.Tick(800*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type tickMsg time.Time

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tickMsg:
		if m.step == stepLoading {
			m.loadingTick++
			if m.loadingTick >= len(m.loadingTitles) {
				m.loadingTick = len(m.loadingTitles) - 1
			}
			return m, tick()
		}

	case tea.KeyMsg:

		// 🌍 Global commands (work at ANY step)
		switch msg.String() {
		case "ctrl+c", "q", "esc", "exit":
			return m, tea.Quit
		}

		// Step-specific logic
		switch m.step {

		case stepLoading:
			if msg.String() == "enter" || msg.String() == " " {
				m.step = stepEnvironment
			}

		case stepEnvironment:
			switch msg.String() {
			case "up":
				if m.envSelected > 0 {
					m.envSelected--
				}
			case "down":
				if m.envSelected < len(m.environments)-1 {
					m.envSelected++
				}
			case "enter":
				m.step = stepUsername
			}

		case stepUsername:
			var cmd tea.Cmd
			m.username, cmd = m.username.Update(msg)

			if msg.String() == "enter" {
				m.step = stepConfirm
			}
			return m, cmd

		case stepConfirm:
			switch msg.String() {
			case "y":
				m.step = stepDone
			case "n":
				m.step = stepEnvironment
			}

		case stepDone:
			// waiting for quit
		}
	}

	return m, nil
}

func (m model) View() string {

	// 🔥 Persistent Header (ALWAYS visible)
	header := headerStyle.Render(`
	************************************************************************************************-----

    ███████╗██████╗ ███████╗ ██████╗████████╗███████╗██████╗ 
    ██╔════╝██╔══██╗██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔══██╗
    ███████╗██████╔╝█████╗  ██║        ██║   █████╗  ██████╔╝
    ╚════██║██╔═══╝ ██╔══╝  ██║        ██║   ██╔══╝  ██╔══██╗
    ███████║██║     ███████╗╚██████╗   ██║   ███████╗██║  ██║
    ╚══════╝╚═╝     ╚══════╝ ╚═════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝
                                                         

    ************************************************************************************************-----  
	`) + "\n\n"

	about := normalStyle.Render(`
	Tiny, Fast, and Deployable anywhere — automate the mundane, unleash your creativity
	`)

	var body string

	spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinnerChar := spinner[m.loadingTick%len(spinner)]

	switch m.step {

	case stepLoading:
		body += spinnerChar + " " + m.loadingTitles[m.loadingTick] + "\n\n"
		body += normalStyle.Render("Press Enter or wait...")

	case stepEnvironment:
		body += "Select Environment:\n\n"
		for i, env := range m.environments {
			if i == m.envSelected {
				body += selectedStyle.Render("> "+env) + "\n"
			} else {
				body += normalStyle.Render("  "+env) + "\n"
			}
		}
		body += "\nUse ↑ ↓ and press Enter"

	case stepUsername:
		body += "Set Username:\n\n"
		body += m.username.View()
		body += "\n\nPress Enter to continue"

	case stepConfirm:
		body += fmt.Sprintf(
			"Confirm Setup:\n\nEnvironment: %s\nUsername: %s\n\nProceed? (y/n)",
			m.environments[m.envSelected],
			m.username.Value(),
		)

	case stepDone:
		body += "✅ Setup Complete!\n\nConfiguration saved."
	}

	// 🌍 Persistent Footer (global help)
	footer := "\n\n" + footerStyle.Render(
		"Commands: ↑ ↓ navigate • Enter select • q/esc/ctrl+c/exit quit anytime",
	)

	return header + about + body + footer
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
