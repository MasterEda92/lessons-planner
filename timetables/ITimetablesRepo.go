package timetables

import (
	"github.com/MasterEda92/lessons-planner/DTOs"
)

type ITimetablesRepo interface {
	GetByDayAndHour(day int, hour string) (*DTOs.TimetableDto, error)
}
