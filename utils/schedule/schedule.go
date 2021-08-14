package schedule

import (
	cron "github.com/robfig/cron/v3"
	"log"
)

var (
	_cron *cron.Cron
)

func init() {
	_cron = cron.New(cron.WithSeconds())
	_cron.Start()
}

type TaskHandler func()
func AddTask(time string, handler TaskHandler) cron.EntryID {
	id, err := _cron.AddFunc(time, handler)
	if err != nil {
		log.Println(err)
	}
	return id
}

func GetScheduler() *cron.Cron {
	return _cron
}