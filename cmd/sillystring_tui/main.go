package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	ss "github.com/notreal/sillystring"
)

type model struct {
	input  textinput.Model
	output string
}

func initialModel() model {
	in := textinput.New()
	in.Placeholder = "Just start typing"
	in.Focus()
	return model{
		input:  in,
		output: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	m.input, _ = m.input.Update(msg)
	m.output = ""
	translations := ss.GetAllTranslations()
	// random order of ranged map is distracting
	sorted := make([]string, 0)
	for name := range translations {
		sorted = append(sorted, name)
	}
	sort.Strings(sorted)
	for _, t := range sorted {
		m.output += fmt.Sprintf("\n%s", ss.Translate(m.input.Value(), translations[t]))
	}
	return m, nil
}

func (m model) View() string {
	s := m.input.View()
	s += m.output
	s += "\n\nAfter copying text, press Ctrl-C or Esc to quit\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
