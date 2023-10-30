package timetables

import (
	"context"

	"github.com/MasterEda92/lessons-planner/DTOs"
	"github.com/MasterEda92/lessons-planner/db"
)

type TimetablesRepo struct {
	client *db.PrismaClient
}

// Create new Repository
func NewTimetablesRepo(client *db.PrismaClient) *TimetablesRepo {
	return &TimetablesRepo{client}
}

// get timetable entry for a specific day and hour
func (r *TimetablesRepo) GetByDayAndHour(day int, hour string) (*DTOs.TimetableDto, error) {
	timetable, err := r.client.Timetable.FindFirst(
		db.Timetable.Day.Equals(day),
		db.Timetable.Hours.Some(
			db.TimetableHours.Hour.Where(db.Hour.Hour.Contains(hour)),
		),
	).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return &DTOs.TimetableDto{
		Id:             timetable.ID,
		Day:            timetable.Day,
		ClassSubjectId: timetable.ClassSubjectID,
	}, nil
}
