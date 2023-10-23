package DTOs

// Get
type ClassDto struct {
	Id           string
	Name         string
	Grade        int
	PupilCount   int
	PupilCountTI int
}

// Create
type CreateClassDto struct {
	Name         string
	Grade        int
	PupilCount   int
	PupilCountTI int
}

type CreateClassResponseDto struct {
	Id           string
	Name         string
	Grade        int
	PupilCount   int
	PupilCountTI int
}

// Update
type UpdateClassDto struct {
	Id           string
	PupilCount   int
	PupilCountTI int
}

type UpdateClassResponseDto struct {
	Id           string
	Name         string
	Grade        int
	PupilCount   int
	PupilCountTI int
}

// Delete
type DeleteClassDto struct {
	Id string
}

type DeleteClassResponseDto struct {
	Id           string
	Name         string
	Grade        int
	PupilCount   int
	PupilCountTI int
}
