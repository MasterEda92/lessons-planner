package subjects

import (
	"github.com/MasterEda92/lessons-planner/DTOs"
)

type ISubjectsRepo interface {
	GetAll() ([]DTOs.SubjectDto, error)
	GetSubjectById(id string) (*DTOs.SubjectDto, error)
	CreateSubject(subject DTOs.CreateSubjectDto) (*DTOs.CreateSubjectResponseDto, error)
	UpdateSubjectWithId(subject DTOs.UpdateSubjectDto) (*DTOs.UpdateSubjectResponseDto, error)
	DeleteSubjectWithId(subject DTOs.DeleteSubjectDto) (*DTOs.DeleteSubjectResponseDto, error)
}
