
import { Grade } from "../enums/Grade"

export type ClassSubjectElement = {
    Id: string,
    ClassId: string,
    ClassName: string,
    ClassGrade: Grade,
    SubjectId: string,
    SubjectName: string,
    Sa: boolean,
    Ex: boolean,
    Ka: boolean
}