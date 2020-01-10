package schedule

import "github.com/GoGroup/Movie-and-events/model"

type ScheduleRepository interface {
	Schedules() ([]model.Schedule, []error)
	StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error)
}