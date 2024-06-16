package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// New program w /initial model and program option
	m := NewModel()
	p := tea.NewProgram(m)
	// Run

	_, err := p.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

// Model : app state
type Model struct {
	title string

	textinput textinput.Model
}

// NewModel : initial model
func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Type something...	"
	ti.Focus()
	return Model{
		title:     "Hello, Bubble Tea!",
		textinput: ti,
	}
}

// Init: kick off the event loop
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update: handle messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			v := m.textinput.Value()
			return m, handleQuerySearch(v)
		}
	}
	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

// View: return a string based on the state of our model
func (m Model) View() string {
	s := m.textinput.View()

	return s
}

// Cmd
func handleQuerySearch(q string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("https://api.urbandictionary.com/v0/define?term=%s", url2.QueryEscape(q))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return ErrorMsg(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return ErrorMsg(err)
		}

		defer res.Body.Close()

		b, _ := io.ReadAll()

	}
}

// Msg
type ErrorMsg error
