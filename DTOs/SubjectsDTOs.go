package DTOs

// Get
type SubjectDto struct {
	Id    string
	Name  string
	Notes string
}

// Create
type CreateSubjectDto struct {
	Name  string
	Notes string
}

type CreateSubjectResponseDto struct {
	Id    string
	Name  string
	Notes string
}

// Update
type UpdateSubjectDto struct {
	Id    string
	Name  string
	Notes string
}

type UpdateSubjectResponseDto struct {
	Id    string
	Name  string
	Notes string
}

// Delete
type DeleteSubjectDto struct {
	Id string
}

type DeleteSubjectResponseDto struct {
	Id    string
	Name  string
	Notes string
}
