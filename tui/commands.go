package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func selectMonitorCmd(id uuid.UUID) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg{ActiveMonitorID: id}
	}
}

func showListCmd() tea.Cmd {
	return func() tea.Msg {
		return ShowListMsg{}
	}
}
