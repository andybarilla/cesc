package tui

import (
	"andybarilla.com/cesc/monitor"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type MonitorListModel struct {
	list list.Model
}

type SelectMsg struct {
	ActiveMonitorID uuid.UUID
}

type ShowListMsg struct {
}

func InitMonitorList(monitors []monitor.Monitor) tea.Model {
	items, err := newMonitorList(monitors)
	if err != nil {
		return nil
	}

	m := MonitorListModel{
		list: list.New(items, list.NewDefaultDelegate(), 8, 8),
	}

	m.list.Title = "Monitors"
	return m
}

func newMonitorList(monitors []monitor.Monitor) ([]list.Item, error) {
	return monitorsToItems(monitors), nil
}

func (m MonitorListModel) View() string {
	return DocStyle.Render(m.list.View())
}

func monitorsToItems(monitors []monitor.Monitor) []list.Item {
	items := make([]list.Item, len(monitors))
	for i, mon := range monitors {
		items[i] = list.Item(mon)
	}
	return items
}

func (m MonitorListModel) Init() tea.Cmd {
	return nil
}

func (m MonitorListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		top, right, bottom, left := DocStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case tea.KeyMsg:
		switch {
		case msg.String() == "q":
			return m, tea.Quit
		case msg.String() == "enter":
			cmd = selectMonitorCmd(m.getActiveMonitorID())
		default:
			m.list, cmd = m.list.Update(msg)
		}
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m MonitorListModel) getActiveMonitorID() uuid.UUID {
	return m.list.SelectedItem().(monitor.Monitor).ID
}
