package main

import (
	"context"
	"fmt"

	"github.com/MasterEda92/lessons-planner/DTOs"
	classsubjects "github.com/MasterEda92/lessons-planner/class_subjects"
	"github.com/MasterEda92/lessons-planner/classes"
	"github.com/MasterEda92/lessons-planner/subjects"
)

// App struct
type App struct {
	ctx               context.Context
	ClassRepo         classes.IClassesRepo
	SubjectRepo       subjects.ISubjectsRepo
	ClassSubjectsRepo classsubjects.IClassSubjectsRepo
}

// NewApp creates a new App application struct
func NewApp(repo classes.IClassesRepo,
	subRepo subjects.ISubjectsRepo,
	clSubRepo classsubjects.IClassSubjectsRepo,
) *App {

	return &App{
		ClassRepo:         repo,
		SubjectRepo:       subRepo,
		ClassSubjectsRepo: clSubRepo,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// classes

// Get all classes
func (a *App) GetAllClasses() ([]DTOs.ClassDto, error) {
	classes, err := a.ClassRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return classes, nil
}

// Get class by its id
func (a *App) GetClassById(id string) (*DTOs.ClassDto, error) {
	_class, err := a.ClassRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Get class with id %s: %v", id, _class)

	return _class, nil
}

// Create a new class
func (a *App) CreateClass(class DTOs.CreateClassDto) (*DTOs.CreateClassResponseDto, error) {
	createdClass, err := a.ClassRepo.CreateClass(class)
	if err != nil {
		return nil, err
	}

	return createdClass, nil
}

// update class
func (a *App) UpdateClass(class DTOs.UpdateClassDto) (*DTOs.UpdateClassResponseDto, error) {
	updatedClass, err := a.ClassRepo.UpdateClassWithId(class)
	if err != nil {
		return nil, err
	}

	return updatedClass, nil
}

// Delete class with given id
func (a *App) DeleteClass(class DTOs.DeleteClassDto) (*DTOs.DeleteClassResponseDto, error) {
	deletedClass, err := a.ClassRepo.DeleteClassWithId(class)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Class deleted: %v", deletedClass)

	return deletedClass, nil
}

// subjects

// Get all subjects
func (a *App) GetAllSubjects() ([]DTOs.SubjectDto, error) {
	subjects, err := a.SubjectRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

// Get subject by its id
func (a *App) GetSubjectById(id string) (*DTOs.SubjectDto, error) {
	subject, err := a.SubjectRepo.GetSubjectById(id)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Get class with id %s: %v", id, subject)

	return subject, nil
}

// Create a new subject
func (a *App) CreateSubject(subject DTOs.CreateSubjectDto) (*DTOs.CreateSubjectResponseDto, error) {
	createdSubject, err := a.SubjectRepo.CreateSubject(subject)
	if err != nil {
		return nil, err
	}

	return createdSubject, nil
}

// update subject
func (a *App) UpdateSubject(subject DTOs.UpdateSubjectDto) (*DTOs.UpdateSubjectResponseDto, error) {
	updatedSubject, err := a.SubjectRepo.UpdateSubjectWithId(subject)
	if err != nil {
		return nil, err
	}

	return updatedSubject, nil
}

// Delete subject with given id
func (a *App) DeleteSubject(subject DTOs.DeleteSubjectDto) (*DTOs.DeleteSubjectResponseDto, error) {
	deletedSubject, err := a.SubjectRepo.DeleteSubjectWithId(subject)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Subject deleted: %v", deletedSubject)

	return deletedSubject, nil
}

// class subjects

// Get all class subjects
func (a *App) GetAllClassSubjects() ([]DTOs.ClassSubjectsDTO, error) {
	cl_subjects, err := a.ClassSubjectsRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return cl_subjects, nil
}

// Get class subject by its id
func (a *App) GetClassSubjectById(id string) (*DTOs.ClassSubjectsDTO, error) {
	cl_subject, err := a.ClassSubjectsRepo.GetClassSubjectById(id)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Get class subject with id %s: %v", id, cl_subject)

	return cl_subject, nil
}

// Create a new class subject
func (a *App) CreateClassSubject(cl_subject DTOs.CreateClassSubjectDTO) (*DTOs.CreateClassSubjectResponseDTO, error) {
	created, err := a.ClassSubjectsRepo.CreateClassSubject(cl_subject)
	if err != nil {
		return nil, err
	}

	return created, nil
}

// update class subject
func (a *App) UpdateClassSubject(cl_subject DTOs.UpdateClassSubjectDTO) (*DTOs.UpdateClassSubjectResponseDTO, error) {
	updated, err := a.ClassSubjectsRepo.UpdateClassSubject(cl_subject)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// delete class subject with given id
func (a *App) DeleteClassSubject(cl_subject DTOs.DeleteClassSubjectDTO) (*DTOs.DeleteClassSubjectResponseDTO, error) {
	deleted, err := a.ClassSubjectsRepo.DeleteClassSubject(cl_subject)
	if err != nil {
		return nil, err
	}

	fmt.Printf("class-subject deleted: %v", deleted)

	return deleted, nil
}
