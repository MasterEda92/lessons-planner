package classsubjects

import "github.com/MasterEda92/lessons-planner/DTOs"

type IClassSubjectsRepo interface {
	GetAll() ([]DTOs.ClassSubjectsDTO, error)
	GetClassSubjectById(Id string) (*DTOs.ClassSubjectsDTO, error)
	CreateClassSubject(classSub DTOs.CreateClassSubjectDTO) (*DTOs.CreateClassSubjectResponseDTO, error)
	UpdateClassSubject(classSub DTOs.UpdateClassSubjectDTO) (*DTOs.UpdateClassSubjectResponseDTO, error)
	DeleteClassSubject(classSub DTOs.DeleteClassSubjectDTO) (*DTOs.DeleteClassSubjectResponseDTO, error)
}
