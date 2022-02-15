package storage

import (
	"cmd/process"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Storage struct {
	list     []*process.Process
	capacity int
}

var (
	pid int
)

func (s *Storage) Init(capacity int) {
	s.capacity = capacity
	pid = 1
}

func (s *Storage) Add(p *process.Process) *process.Process {
	if pid > s.capacity {
		return nil
	}

	p.PID = int32(pid)
	pid++

	s.list = append(s.list, p)

	return p
}

func (s *Storage) Kill(ID int32) {
	for index := 0; index < len(s.list); index++ {
		if proc := s.list[index]; ID == proc.PID {
			proc.Terminate = true
			s.list = append(s.list[:index], s.list[index+1:]...)
			break
		}
	}
}

func (s Storage) Pause(ID int32, flag bool) {
	for _, p := range s.list {
		if p.PID == ID {
			p.Pause = flag
			return
		}
	}
}

func (s Storage) View() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBlackOnGreenWhite)
	t.SetTitle("Monitoring Processes")
	t.AppendHeader(table.Row{"#", "PID", "Delay", "Status", "Created At", "Last Update", "Task", "Executed"})

	for i, p := range s.list {
		var status string
		if status = "RUNNING"; p.Pause {
			status = "WAITING"
		}

		t.AppendRow([]interface{}{i + 1, p.PID, p.Delay, status, p.CreatedAt.Format(time.StampMilli), p.UpdatedAt.Format(time.StampMilli), p.Task, p.Called})
	}
	t.Render()
}
