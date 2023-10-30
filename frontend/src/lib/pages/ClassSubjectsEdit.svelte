<script lang="ts">
  import { navigate } from "svelte-routing";
  import { GetClassSubjectById, UpdateClassSubject } from "../../../wailsjs/go/main/App";
  import Heading1 from "../components/headings/heading1.svelte";
  import { Grade } from "../enums/Grade";
  import type { ClassSubjectElement } from "../types/classsubjects";
  

  export let id: string

  let class_sub: ClassSubjectElement = {
    Id: id,
    ClassId: "0",
    ClassName: "",
    ClassGrade: Grade.NONE,
    SubjectId: "0",
    SubjectName: "",
    Sa: false,
    Ex: false,
    Ka: false
  }

  function GetClassSubject () {
    GetClassSubjectById (id).then ((result) => {
      class_sub = {
        Id: result.Id,
        ClassId: result.ClassId,
        ClassName: result.ClassName,
        ClassGrade: result.ClassGrade,
        SubjectId: result.SubjectId,
        SubjectName: result.SubjectName,
        Sa: result.Sa,
        Ex: result.Ex,
        Ka: result.Ka
      }
    }).catch ((error) => {
      alert (error)
    })
  }

  GetClassSubject ();

  function updateClassSubject () {
    UpdateClassSubject ({
      Id: class_sub.Id,
      Sa: class_sub.Sa,
      Ex: class_sub.Ex,
      Ka: class_sub.Ka
    }).then ((result) => {
      navigate("/class_subjects", { replace: true });
    }).catch ((error) => {
        alert (error);
    })
  }

</script>

<Heading1>Edit Class-Subject</Heading1>

<form class="py-3 space-y-6">
  <div>
      <label for="className" class="block mb-2 text-sm font-medium text-purple-950">Class</label>
      <input disabled bind:value={class_sub.ClassName} type="text" name="className" id="className" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-1/3 p-2.5">
  </div>
  <div>
    <label for="subjectName" class="block mb-2 text-sm font-medium text-purple-950">Subject</label>
    <input disabled bind:value={class_sub.SubjectName} type="text" name="subjectName" id="subjectName" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-1/3 p-2.5">
</div>
  <div>
    <label for="Grade" class="block mb-2 text-sm font-medium text-purple-950">Grade</label>
    <input disabled bind:value={class_sub.ClassGrade} type="text" name="Grade" id="Grade" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-1/3 p-2.5">
  </div>
  <div>
    <label for="SA" class="block mb-2 text-sm font-medium text-purple-950">SA</label>
    <input bind:checked={class_sub.Sa} type="checkbox" name="SA" id="SA" class="w-4 h-4 text-purple-950 bg-gray-100 border-gray-300 rounded focus:ring-purple-950 focus:ring-2">      
  </div>
  <div>
    <label for="EX" class="block mb-2 text-sm font-medium text-purple-950">Ex</label>
    <input bind:checked={class_sub.Ex} type="checkbox" name="EX" id="EX" class="w-4 h-4 text-purple-950 bg-gray-100 border-gray-300 rounded focus:ring-purple-950 focus:ring-2">      
  </div>
  <div>
    <label for="KA" class="block mb-2 text-sm font-medium text-purple-950">KA</label>
    <input bind:checked={class_sub.Ka} type="checkbox" name="KA" id="KA" class="w-4 h-4 text-purple-950 bg-gray-100 border-gray-300 rounded focus:ring-purple-950 focus:ring-2">      
  </div>
  <button 
      on:click={updateClassSubject} 
      type="button" 
      class="btn btn-primary">
      Save
  </button>
</form>
