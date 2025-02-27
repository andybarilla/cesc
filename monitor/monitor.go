package monitor

import (
	"log"
	"os/exec"

	"github.com/google/uuid"
)

type Monitor struct {
	ID         uuid.UUID
	Name       string
	Path       string
	Args       []string
	Env        []string
	Dir        string
	wasStarted bool
	cmd        *exec.Cmd
}

func GetMonitors() ([]Monitor, error) {
	return []Monitor{
		{
			ID:         uuid.New(),
			Name:       "Hello",
			Path:       "echo",
			Args:       []string{"hello"},
			wasStarted: false,
		},
		{
			ID:         uuid.New(),
			Name:       "List",
			Path:       "ls",
			wasStarted: false,
		},
		{
			ID:         uuid.New(),
			Name:       "Longer Message",
			Path:       "./testslow.sh",
			wasStarted: false,
		},
	}, nil
}

func (m Monitor) Title() string {
	return m.Name
}

func (m Monitor) Description() string {
	return m.ID.String()
}

func (m Monitor) FilterValue() string {
	return m.Name
}

func (m *Monitor) Run() error {
	log.Printf("Running monitor: %s", m.Name)
	m.wasStarted = true
	cmd := exec.Command(m.Path, m.Args...)
	cmd.Env = m.Env
	cmd.Dir = m.Dir
	m.cmd = cmd
	return m.cmd.Start()
}

func (m *Monitor) Output() string {
	if !m.wasStarted {
		m.Run()
	}

	out, err := m.cmd.
	if err != nil {
		return err.Error()
	}
	return string(out)
}
