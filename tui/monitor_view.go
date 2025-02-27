package tui

import (
	"andybarilla.com/cesc/monitor"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type MonitorViewModel struct {
	monitor  *monitor.Monitor
	textarea textarea.Model
	err      error
}

func InitMonitorView(monitor *monitor.Monitor) tea.Model {
	m := MonitorViewModel{
		monitor:  monitor,
		textarea: textarea.New(),
	}

	return m
}

func (m MonitorViewModel) Init() tea.Cmd {
	return nil
}

func (m MonitorViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.String() == "q":
			return m, tea.Quit
		case msg.String() == "esc":
			return m, showListCmd()
		default:
			m.textarea, cmd = m.textarea.Update(msg)
		}
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m MonitorViewModel) View() string {
	m.textarea.SetValue(m.monitor.Output())
	return m.textarea.View()
}
