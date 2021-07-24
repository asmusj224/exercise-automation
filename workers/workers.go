package workers

import (
	"github.com/robfig/cron/v3"
)

type Workers interface {
	Start()
}

type workers struct{}

func NewWorkers() Workers {
	return &workers{}
}

func (*workers) Start() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0 4 ? * MON-FRI", func() {
		// fetch random exercise
		// format string
		// send email
	})
	c.Start()
}
