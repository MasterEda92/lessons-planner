package classes

import (
	"github.com/MasterEda92/lessons-planner/DTOs"
)

type IClassesRepo interface {
	GetAll() ([]DTOs.ClassDto, error)
	GetById(id string) (*DTOs.ClassDto, error)
	CreateClass(class DTOs.CreateClassDto) (*DTOs.CreateClassResponseDto, error)
	UpdateClassWithId(class DTOs.UpdateClassDto) (*DTOs.UpdateClassResponseDto, error)
	DeleteClassWithId(class DTOs.DeleteClassDto) (*DTOs.DeleteClassResponseDto, error)
}
