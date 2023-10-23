package DTOs

type ClassSubjectsDTO struct {
	Id          string
	ClassId     string
	ClassName   string
	ClassGrade  int
	SubjectId   string
	SubjectName string
	Sa          bool
	Ex          bool
	Ka          bool
}

type CreateClassSubjectDTO struct {
	ClassId   string
	SubjectId string
	Sa        bool
	Ex        bool
	Ka        bool
}

type CreateClassSubjectResponseDTO struct {
	Id          string
	ClassId     string
	ClassName   string
	ClassGrade  int
	SubjectId   string
	SubjectName string
	Sa          bool
	Ex          bool
	Ka          bool
}

type UpdateClassSubjectDTO struct {
	Id string
	Sa bool
	Ex bool
	Ka bool
}

type UpdateClassSubjectResponseDTO struct {
	Id          string
	ClassId     string
	ClassName   string
	ClassGrade  int
	SubjectId   string
	SubjectName string
	Sa          bool
	Ex          bool
	Ka          bool
}

type DeleteClassSubjectDTO struct {
	Id string
}

type DeleteClassSubjectResponseDTO struct {
	Id        string
	ClassId   string
	SubjectId string
	Sa        bool
	Ex        bool
	Ka        bool
}
