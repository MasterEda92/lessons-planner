package classsubjects

import (
	"context"
	"errors"

	"github.com/MasterEda92/lessons-planner/DTOs"
	"github.com/MasterEda92/lessons-planner/db"

	"github.com/google/uuid"
)

var (
	ErrClassNotFound   = errors.New("given class does not exist")
	ErrSubjectNotFound = errors.New("given subject does not exist")
)

type ClassSubjectsRepo struct {
	client *db.PrismaClient
}

// Create new Repository
func NewClassSubjectsRepo(client *db.PrismaClient) *ClassSubjectsRepo {
	return &ClassSubjectsRepo{client}
}

// Get all ClassSubjects
func (r *ClassSubjectsRepo) GetAll() ([]DTOs.ClassSubjectsDTO, error) {
	class_subjects, err := r.client.ClassSubject.FindMany().
		With(
			db.ClassSubject.Class.Fetch(),
			db.ClassSubject.Subject.Fetch()).
		Exec(context.Background())

	if err != nil {
		return nil, err
	}

	var _class_subjects []DTOs.ClassSubjectsDTO

	for _, clsubject := range class_subjects {
		sub := DTOs.ClassSubjectsDTO{
			Id:          clsubject.ID,
			ClassId:     clsubject.ClassID,
			ClassName:   clsubject.Class().InnerClass.Name,
			ClassGrade:  clsubject.Class().InnerClass.Grade,
			SubjectId:   clsubject.SubjectID,
			SubjectName: clsubject.Subject().InnerSubject.Name,
			Ex:          clsubject.Ex,
			Sa:          clsubject.Sa,
			Ka:          clsubject.Ka,
		}

		_class_subjects = append(_class_subjects, sub)
	}

	return _class_subjects, nil
}

// get a single ClassSubject by its id
func (r *ClassSubjectsRepo) GetClassSubjectById(Id string) (*DTOs.ClassSubjectsDTO, error) {
	cl, err := r.client.ClassSubject.FindUnique(db.ClassSubject.ID.Equals(Id)).
		With(
			db.ClassSubject.Class.Fetch(),
			db.ClassSubject.Subject.Fetch()).
		Exec(context.Background())

	if err != nil {
		return nil, err
	}

	_cl := DTOs.ClassSubjectsDTO{
		Id:          cl.ID,
		ClassId:     cl.ClassID,
		ClassName:   cl.Class().InnerClass.Name,
		ClassGrade:  cl.Class().InnerClass.Grade,
		SubjectId:   cl.SubjectID,
		SubjectName: cl.Subject().InnerSubject.Name,
		Ex:          cl.Ex,
		Sa:          cl.Sa,
		Ka:          cl.Ka,
	}

	return &_cl, nil
}

// create a new ClassSubject
func (r *ClassSubjectsRepo) CreateClassSubject(classSub DTOs.CreateClassSubjectDTO) (*DTOs.CreateClassSubjectResponseDTO, error) {
	ctx := context.Background()

	// check if class with given id exists
	_, err := r.client.Class.FindUnique(db.Class.ID.Equals(classSub.ClassId)).Exec(ctx)
	if errors.Is(err, db.ErrNotFound) {
		return nil, ErrClassNotFound
	}
	if err != nil {
		return nil, err
	}

	// check if subject with given id exists
	_, err = r.client.Subject.FindUnique(db.Subject.ID.Equals(classSub.SubjectId)).Exec(ctx)
	if errors.Is(err, db.ErrNotFound) {
		return nil, ErrSubjectNotFound
	}
	if err != nil {
		return nil, err
	}

	// create new ClassSubject
	createdCS, err := r.client.ClassSubject.CreateOne(
		db.ClassSubject.ID.Set(uuid.NewString()),
		db.ClassSubject.Class.Link(
			db.Class.ID.Equals(classSub.ClassId),
		),
		db.ClassSubject.Subject.Link(
			db.Subject.ID.Equals(classSub.SubjectId),
		),
		db.ClassSubject.Sa.Set(classSub.Sa),
		db.ClassSubject.Ex.Set(classSub.Ex),
		db.ClassSubject.Ka.Set(classSub.Ka),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	// Fetch additional data from created class subject
	cl, err := r.client.ClassSubject.FindUnique(db.ClassSubject.ID.Equals(createdCS.ID)).
		With(
			db.ClassSubject.Class.Fetch(),
			db.ClassSubject.Subject.Fetch()).
		Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return &DTOs.CreateClassSubjectResponseDTO{
		Id:          cl.ID,
		ClassId:     cl.ClassID,
		ClassName:   cl.Class().InnerClass.Name,
		ClassGrade:  cl.Class().InnerClass.Grade,
		SubjectId:   cl.SubjectID,
		SubjectName: cl.Subject().InnerSubject.Name,
		Sa:          cl.Sa,
		Ka:          cl.Ka,
		Ex:          cl.Ex,
	}, nil
}

// update existing ClassSubject
func (r *ClassSubjectsRepo) UpdateClassSubject(classSub DTOs.UpdateClassSubjectDTO) (*DTOs.UpdateClassSubjectResponseDTO, error) {
	updated, err := r.client.ClassSubject.FindUnique(
		db.ClassSubject.ID.Equals(classSub.Id)).
		Update(
			db.ClassSubject.Sa.Set(classSub.Sa),
			db.ClassSubject.Ex.Set(classSub.Ex),
			db.ClassSubject.Ka.Set(classSub.Ka),
		).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	// Fetch additional data from updated class subject
	cl, err := r.client.ClassSubject.FindUnique(db.ClassSubject.ID.Equals(updated.ID)).
		With(
			db.ClassSubject.Class.Fetch(),
			db.ClassSubject.Subject.Fetch()).
		Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return &DTOs.UpdateClassSubjectResponseDTO{
		Id:          cl.ID,
		ClassId:     cl.ClassID,
		ClassName:   cl.Class().InnerClass.Name,
		ClassGrade:  cl.Class().InnerClass.Grade,
		SubjectId:   cl.SubjectID,
		SubjectName: cl.Subject().InnerSubject.Name,
		Sa:          cl.Sa,
		Ka:          cl.Ka,
		Ex:          cl.Ex,
	}, nil
}

// delete existing ClassSubject with given id
func (r *ClassSubjectsRepo) DeleteClassSubject(classSub DTOs.DeleteClassSubjectDTO) (*DTOs.DeleteClassSubjectResponseDTO, error) {
	deleted, err := r.client.ClassSubject.FindUnique(
		db.ClassSubject.ID.Equals(classSub.Id),
	).Delete().Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return &DTOs.DeleteClassSubjectResponseDTO{
		Id:        deleted.ID,
		ClassId:   deleted.ClassID,
		SubjectId: deleted.SubjectID,
		Sa:        deleted.Sa,
		Ka:        deleted.Ka,
		Ex:        deleted.Ex,
	}, nil

}
