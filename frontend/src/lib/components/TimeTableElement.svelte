<script lang="ts">
    import { GetTimetableByDayAndHour, GetClassSubjectById } from "../../../wailsjs/go/main/App";
    import type { Day } from "../enums/Day";
    import { Grade } from "../enums/Grade";
    import type { ClassSubjectElement } from "../types/classsubjects";
    import ClassSubjectsCombobox from "./comboboxes/ClassSubjectsCombobox.svelte";
    
    export let Day: Day;
    export let Hour: string;

    let Data: ClassSubjectElement = {
        Id: "0",
        ClassId: "0",
        ClassName: "",
        ClassGrade: Grade.NONE,
        SubjectId: "0",
        SubjectName: "",
        Sa: false,
        Ex: false,
        Ka: false
    };

    getTimetable ();

    function getTimetable () {
        GetTimetableByDayAndHour (Day, Hour).then ((result) => {
            if (result !== undefined) {
                GetClassSubjectById(result.ClassSubjectId).then ((clsub) => {
                if (clsub !== undefined)
                {
                    Data = {
                        Id: clsub.Id,
                        ClassId: clsub.ClassId,
                        ClassName: clsub.ClassName,
                        ClassGrade: clsub.ClassGrade,
                        SubjectId: clsub.SubjectId,
                        SubjectName: clsub.SubjectName,
                        Sa: clsub.Sa,
                        Ex: clsub.Ex,
                        Ka: clsub.Ka
                    }
                }                
                }).catch ((err) => {
                    alert (err);
                })
            }            
        }).catch ((error) => {
            alert (error);
        })
    }

</script>

<ClassSubjectsCombobox bind:Value={Data.Id} />

