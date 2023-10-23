package classes

import (
	"context"
	"fmt"

	"github.com/MasterEda92/lessons-planner/DTOs"
	"github.com/MasterEda92/lessons-planner/db"

	"github.com/google/uuid"
)

type ClassesRepo struct {
	client *db.PrismaClient
}

// Create new Repository
func NewClassesRepo(client *db.PrismaClient) *ClassesRepo {
	return &ClassesRepo{client}
}

// Get all classes from the database
func (r *ClassesRepo) GetAll() ([]DTOs.ClassDto, error) {

	classes, err := r.client.Class.FindMany().Exec(context.Background())

	if err != nil {
		return nil, err
	}

	var _classes []DTOs.ClassDto

	for _, class := range classes {
		cl := DTOs.ClassDto{
			Id:           class.ID,
			Name:         class.Name,
			Grade:        class.Grade,
			PupilCount:   class.PupilCount,
			PupilCountTI: class.PupilCountTimeIncrease,
		}

		_classes = append(_classes, cl)
	}

	return _classes, nil
}

// Get class with given Id
func (r *ClassesRepo) GetById(id string) (*DTOs.ClassDto, error) {
	_class, err := r.client.Class.FindUnique(db.Class.ID.Equals(id)).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	fmt.Printf("Get class with id %s from db: %v", id, _class)

	return &DTOs.ClassDto{
		Id:           _class.ID,
		Name:         _class.Name,
		Grade:        _class.Grade,
		PupilCount:   _class.PupilCount,
		PupilCountTI: _class.PupilCountTimeIncrease,
	}, nil
}

// Create a new class
func (r *ClassesRepo) CreateClass(class DTOs.CreateClassDto) (*DTOs.CreateClassResponseDto, error) {
	createdClass, err := r.client.Class.CreateOne(
		db.Class.ID.Set(uuid.NewString()),
		db.Class.Name.Set(class.Name),
		db.Class.Grade.Set(class.Grade),
		db.Class.PupilCount.Set(class.PupilCount),
		db.Class.PupilCountTimeIncrease.Set(class.PupilCountTI),
	).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return &DTOs.CreateClassResponseDto{
		Id:           createdClass.ID,
		Name:         createdClass.Name,
		Grade:        createdClass.Grade,
		PupilCount:   createdClass.PupilCount,
		PupilCountTI: createdClass.PupilCountTimeIncrease,
	}, nil
}

// Update class
func (r *ClassesRepo) UpdateClassWithId(_class DTOs.UpdateClassDto) (*DTOs.UpdateClassResponseDto, error) {

	updatedClass, err := r.client.Class.FindUnique(
		db.Class.ID.Equals(_class.Id)).Update(
		db.Class.PupilCount.Set(_class.PupilCount),
		db.Class.PupilCountTimeIncrease.Set(_class.PupilCountTI),
	).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return &DTOs.UpdateClassResponseDto{
		Id:           updatedClass.ID,
		Name:         updatedClass.Name,
		Grade:        updatedClass.Grade,
		PupilCount:   updatedClass.PupilCount,
		PupilCountTI: updatedClass.PupilCountTimeIncrease,
	}, nil
}

// delete class with given id
func (r *ClassesRepo) DeleteClassWithId(class DTOs.DeleteClassDto) (*DTOs.DeleteClassResponseDto, error) {
	deletedClass, err := r.client.Class.FindUnique(
		db.Class.ID.Equals(class.Id),
	).Delete().Exec(context.Background())

	if err != nil {
		return nil, err
	}

	fmt.Printf("Class deleted from database: %v", deletedClass)

	return &DTOs.DeleteClassResponseDto{
		Id:           deletedClass.ID,
		Name:         deletedClass.Name,
		Grade:        deletedClass.Grade,
		PupilCount:   deletedClass.PupilCount,
		PupilCountTI: deletedClass.PupilCountTimeIncrease,
	}, nil
}
