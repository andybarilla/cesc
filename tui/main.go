package tui

import (
	"fmt"
	"log"
	"os"

	"andybarilla.com/cesc/monitor"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	monitorListView sessionState = iota
	monitorView
)

type MainModel struct {
	state         sessionState
	monitorList   tea.Model
	monitorView   tea.Model
	monitors      []monitor.Monitor
	activeMonitor *monitor.Monitor
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case SelectMsg:
		m.activeMonitor = nil
		for i := range m.monitors {
			if m.monitors[i].ID == msg.ActiveMonitorID {
				m.activeMonitor = &m.monitors[i]
			}
		}
		if m.activeMonitor == nil {
			panic(fmt.Errorf("could not find active monitor: %s", msg.ActiveMonitorID))
		}
		m.state = monitorView
	case ShowListMsg:
		m.state = monitorListView
	}

	switch m.state {
	case monitorListView:
		newList, newCmd := m.monitorList.Update(msg)
		listModel, ok := newList.(MonitorListModel)
		if !ok {
			panic("could not perform type assertion on monitor list model")
		}
		m.monitorList = listModel
		cmd = newCmd
	case monitorView:
		m.monitorView = InitMonitorView(m.activeMonitor)
		newEntry, newCmd := m.monitorView.Update(msg)
		viewModel, ok := newEntry.(MonitorViewModel)
		if !ok {
			panic("could not perform type assertion on monitor view model")
		}
		m.monitorView = viewModel
		cmd = newCmd
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	switch m.state {
	case monitorListView:
		return m.monitorList.View()
	case monitorView:
		return m.monitorView.View()
	}
	return ""
}

func StartTea() {
	if f, err := tea.LogToFile("debug.log", "[cesc]"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	monitors, err := monitor.GetMonitors()
	if err != nil {
		panic(fmt.Errorf("cannot get all monitors: %w", err))
	}

	m := MainModel{
		state:       monitorListView,
		monitorList: InitMonitorList(monitors),
		monitors:    monitors,
	}

	log.Println("Starting cesc...")

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Great googly-moogly, an error has occurred: %v", err)
		os.Exit(1)
	}
}
