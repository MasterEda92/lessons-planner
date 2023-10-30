package subjects

import (
	"context"
	"fmt"

	"github.com/MasterEda92/lessons-planner/DTOs"
	"github.com/MasterEda92/lessons-planner/db"

	"github.com/google/uuid"
)

type SubjectsRepo struct {
	client *db.PrismaClient
}

// Create new Repository
func NewSubjectsRepo(client *db.PrismaClient) *SubjectsRepo {
	return &SubjectsRepo{client}
}

// Get all classes from the database
func (r *SubjectsRepo) GetAll() ([]DTOs.SubjectDto, error) {

	subjects, err := r.client.Subject.FindMany().Exec(context.Background())

	if err != nil {
		return nil, err
	}

	var _subjects []DTOs.SubjectDto

	for _, subject := range subjects {
		n, ok := subject.Notes()
		if !ok {
			n = ""
		}
		sub := DTOs.SubjectDto{
			Id:    subject.ID,
			Name:  subject.Name,
			Notes: n,
		}

		_subjects = append(_subjects, sub)
	}

	return _subjects, nil
}

// Get subject with given id
func (r *SubjectsRepo) GetSubjectById(id string) (*DTOs.SubjectDto, error) {
	_subject, err := r.client.Subject.FindUnique(db.Subject.ID.Equals(id)).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	fmt.Printf("Get subject with id %s from db: %v", id, _subject)

	n, ok := _subject.Notes()
	if !ok {
		n = ""
	}
	return &DTOs.SubjectDto{
		Id:    _subject.ID,
		Name:  _subject.Name,
		Notes: n,
	}, nil
}

// create a new subject
func (r *SubjectsRepo) CreateSubject(subject DTOs.CreateSubjectDto) (*DTOs.CreateSubjectResponseDto, error) {
	createdSubject, err := r.client.Subject.CreateOne(
		db.Subject.ID.Set(uuid.NewString()),
		db.Subject.Name.Set(subject.Name),
		db.Subject.Notes.Set(subject.Notes),
	).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	n, ok := createdSubject.Notes()
	if !ok {
		n = ""
	}
	return &DTOs.CreateSubjectResponseDto{
		Id:    createdSubject.ID,
		Name:  createdSubject.Name,
		Notes: n,
	}, nil
}

// update subject with given id
func (r *SubjectsRepo) UpdateSubjectWithId(subject DTOs.UpdateSubjectDto) (*DTOs.UpdateSubjectResponseDto, error) {
	updatedSubject, err := r.client.Subject.FindUnique(
		db.Subject.ID.Equals(subject.Id)).
		Update(
			db.Subject.Name.Set(subject.Name),
			db.Subject.Notes.Set(subject.Notes),
		).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	n, ok := updatedSubject.Notes()
	if !ok {
		n = ""
	}
	return &DTOs.UpdateSubjectResponseDto{
		Id:    updatedSubject.ID,
		Name:  updatedSubject.Name,
		Notes: n,
	}, nil
}

// delete subject with given id
func (r *SubjectsRepo) DeleteSubjectWithId(subject DTOs.DeleteSubjectDto) (*DTOs.DeleteSubjectResponseDto, error) {
	deletedSubject, err := r.client.Subject.FindUnique(
		db.Subject.ID.Equals(subject.Id),
	).Delete().Exec(context.Background())

	if err != nil {
		return nil, err
	}

	fmt.Printf("Subject deleted from database: %v", deletedSubject)

	n, ok := deletedSubject.Notes()
	if !ok {
		n = ""
	}
	return &DTOs.DeleteSubjectResponseDto{
		Id:    deletedSubject.ID,
		Name:  deletedSubject.Name,
		Notes: n,
	}, nil
}
