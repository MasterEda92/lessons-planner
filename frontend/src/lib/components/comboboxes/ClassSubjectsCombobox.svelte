<script lang="ts">
    import { GetAllClassSubjects } from "../../../../wailsjs/go/main/App";
  import { Grade } from "../../enums/Grade";
    import type { ClassSubjectElement } from "../../types/classsubjects";
    


    let clsubs: ClassSubjectElement[] = [];
  
    export let Value: string = "0";
    
    function GetClassSubjects () {
        GetAllClassSubjects ().then ((result) => {
            clsubs = [];
            clsubs.push ({
                    Id: "0",
                    ClassId: "0",
                    ClassName: "",
                    ClassGrade: Grade.NONE,
                    SubjectId: "0",
                    SubjectName: "",
                    Sa: false,
                    Ex: false,
                    Ka: false
                })
            result.forEach(element => {
                clsubs.push ({
                    Id: element.Id,
                    ClassId: element.ClassId,
                    ClassName: element.ClassName,
                    ClassGrade: element.ClassGrade,
                    SubjectId: element.SubjectId,
                    SubjectName: element.SubjectName,
                    Sa: element.Sa,
                    Ex: element.Ex,
                    Ka: element.Ka
                })
            });
            clsubs = [...clsubs];
        }).catch ((error) => {
            alert (error);
        })
    }

    GetClassSubjects ();
  
</script>

<select id="class_subjects_combobox" bind:value={Value} class="bg-gray-50 border border-gray-300 text-purple-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-auto mr-2 my-1">
{#each clsubs as _clsub (_clsub.Id)}
    <option value={_clsub.Id}>{_clsub.ClassName} - {_clsub.SubjectName}</option>
{/each}
</select>