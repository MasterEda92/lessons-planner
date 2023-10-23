<script lang="ts">
  import { onMount } from 'svelte';
  import Heading1 from "../components/headings/heading1.svelte";
  import GradeFilter from "../components/filters/GradeFilter.svelte";
  import ClassFilter from "../components/filters/ClassFilter.svelte";
  import SubjectFilter from "../components/filters/SubjectFilter.svelte";
  import { Grade } from "../enums/Grade";
  import type { ClassSubjectElement } from "../types/classsubjects";  
  import { navigate } from "svelte-routing";
  import { GetAllClassSubjects, DeleteClassSubject, CreateClassSubject } from "../../../wailsjs/go/main/App";
  import { initFlowbite } from 'flowbite';
  import ClassesCombobox from '../components/comboboxes/ClassesCombobox.svelte';
  import SubjectsCombobox from '../components/comboboxes/SubjectsCombobox.svelte';

  // component initialisation
  onMount (() => {
        initFlowbite();
    });
  
  // variables for filters
  let _subject: string = "0";
  let _class: string = "0";
  let _grade: Grade = Grade.NONE;
  
  // change handlers for filters
  function selChangeSubject (event: any) {
    applyFilters (_grade, _class, event.target.value);
  }
  function selChangeClass (event: any) {
    applyFilters (_grade, event.target.value, _subject);
  }
  function selChangeGrade (event: any) {
    applyFilters (event.target.value, _class, _subject);
  }

  // filter handling
  function applyFilters (grade: Grade, cl: string, sub: string) {
        cl_subs_filtered = [];
        class_subjects.forEach(element => {
            if ((grade == element.ClassGrade || grade == Grade.NONE) 
              && (cl == element.ClassId || cl == "0") 
              && (sub == element.SubjectId || sub == "0")) {
              cl_subs_filtered.push(element);
            }
        });
        cl_subs_filtered = [...cl_subs_filtered];
    }

  // data
  let class_subjects: ClassSubjectElement[] = []
  let cl_subs_filtered: ClassSubjectElement[] = [...class_subjects];

  // init data
  function getClassSubjects () {
    GetAllClassSubjects ().then ((result) => {
      class_subjects = [];
      result.forEach(element => {
        class_subjects.push ({
          Id: element.Id,
          ClassId: element.ClassId,
          ClassName: element.ClassName,
          ClassGrade: element.ClassGrade,
          SubjectId: element.SubjectId,
          SubjectName: element.SubjectName,
          Sa: element.Sa,
          Ex: element.Ex,
          Ka: element.Ka
        });
        cl_subs_filtered = [...class_subjects];
      });
    }).catch ((error) => {
      alert (error);
    })
  }

  // get data
  getClassSubjects();  

  // delete class subject
  function deleteClassSubject (id: string) {
    DeleteClassSubject ({
      Id: id
    }).then ((result) => {
      getClassSubjects ();
      applyFilters (_grade, _class, _subject);
    }).catch ((error) => {
      alert (error);
    })
  }

  // variables for createClassSubject
  let newClassSubClassId: string = "";
  let newClassSubSubjectId: string = "";
  let newClassSubSA: boolean = false;
  let newClassSubEx: boolean = false;
  let newClassSubKA: boolean = false;

  // create a new class subject
  function createClassSubject () {
    CreateClassSubject ({
      ClassId: newClassSubClassId,
      SubjectId: newClassSubSubjectId,
      Sa: newClassSubSA,
      Ex: newClassSubEx,
      Ka: newClassSubKA
    }).then ((result) => {
      class_subjects.push ({
          Id: result.Id,
          ClassId: result.ClassId,
          ClassName: result.ClassName,
          ClassGrade: result.ClassGrade,
          SubjectId: result.SubjectId,
          SubjectName: result.SubjectName,
          Sa: result.Sa,
          Ex: result.Ex,
          Ka: result.Ka
      })
      applyFilters (_grade, _class, _subject);
    }).catch ((error) => {
      alert (error);
    })

    newClassSubClassId = "";
    newClassSubSubjectId = "";
    newClassSubSA = false;
    newClassSubEx = false;
    newClassSubKA = false;
  }

</script>

<Heading1>Class-Subjects</Heading1>

<!-- Filter Line-->
<div class="py-2 flex items-center justify-between">
  <div class="flex items-center">
    <SubjectFilter selChange={selChangeSubject} bind:Value={_subject} />
    <ClassFilter selChange={selChangeClass} bind:Value={_class} />
    <GradeFilter selChange={selChangeGrade} bind:Value={_grade} />
  </div>
  <button 
      data-modal-target="create_modal" 
      data-modal-toggle="create_modal" 
      class="btn btn-primary"
      type="button">
      New
  </button>
</div>

<!-- Class-Table -->
<div class="py-2">
  <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
      <table class="w-full text-sm text-left text-gray-500">
          <thead class="text-xs text-purple-950 uppercase bg-gray-200">
              <tr>
                  <th scope="col" class="px-6 py-3">
                      Class
                  </th>
                  <th scope="col" class="px-6 py-3">
                      Subject
                  </th>
                  <th scope="col" class="px-6 py-3">
                      SA
                  </th>
                  <th scope="col" class="px-6 py-3">
                      Ex
                  </th>
                  <th scope="col" class="px-6 py-3">
                    KA
                  </th>
                  <th scope="col" class="px-6 py-3">
                      <span class="sr-only">Edit</span>
                  </th>
                  <th scope="col" class="px-6 py-3">
                      <span class="sr-only">Delete</span>
                  </th>
              </tr>
          </thead>
          <tbody>
              {#each cl_subs_filtered as _clsub (_clsub.Id)}
                  <tr class="bg-white hover:bg-gray-50">
                      <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                          {_clsub.ClassName}
                      </th>
                      <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                          {_clsub.SubjectName}
                      </td>
                      <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                        <input disabled bind:checked={_clsub.Sa} type="checkbox" class="w-4 h-4 text-purple-950 bg-gray-100 border-gray-300 rounded focus:ring-purple-950 focus:ring-2">
                      </td>
                      <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                        <input disabled bind:checked={_clsub.Ex} type="checkbox" class="w-4 h-4 text-purple-950 bg-gray-100 border-gray-300 rounded focus:ring-purple-950 focus:ring-2">
                      </td>
                      <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                        <input disabled bind:checked={_clsub.Ka} type="checkbox" class="w-4 h-4 text-purple-950 bg-gray-100 border-gray-300 rounded focus:ring-purple-950 focus:ring-2">
                      </td>
                      <td class="px-6 py-4 text-right">
                          <button 
                              on:click={() => {
                                  const url = "/class_subjects/" + _clsub.Id;
                                  navigate(url, { replace: true });
                              }}>
                              <svg class="fill-purple-950" width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M13.94 5 19 10.06 9.062 20a2.25 2.25 0 0 1-.999.58l-5.116 1.395a.75.75 0 0 1-.92-.921l1.395-5.116a2.25 2.25 0 0 1 .58-.999L13.938 5Zm7.09-2.03a3.578 3.578 0 0 1 0 5.06l-.97.97L15 3.94l.97-.97a3.578 3.578 0 0 1 5.06 0Z"/></svg>
                          </button>
                      </td>
                      <td class="px-6 py-4 text-right">
                          <button on:click={ () => deleteClassSubject(_clsub.Id)}>
                              <svg class="fill-purple-950" width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M21.5 6a1 1 0 0 1-.883.993L20.5 7h-.845l-1.231 12.52A2.75 2.75 0 0 1 15.687 22H8.313a2.75 2.75 0 0 1-2.737-2.48L4.345 7H3.5a1 1 0 0 1 0-2h5a3.5 3.5 0 1 1 7 0h5a1 1 0 0 1 1 1Zm-7.25 3.25a.75.75 0 0 0-.743.648L13.5 10v7l.007.102a.75.75 0 0 0 1.486 0L15 17v-7l-.007-.102a.75.75 0 0 0-.743-.648Zm-4.5 0a.75.75 0 0 0-.743.648L9 10v7l.007.102a.75.75 0 0 0 1.486 0L10.5 17v-7l-.007-.102a.75.75 0 0 0-.743-.648ZM12 3.5A1.5 1.5 0 0 0 10.5 5h3A1.5 1.5 0 0 0 12 3.5Z"/></svg>
                          </button>
                      </td>
                  </tr>
              {/each}
          </tbody>
      </table>
  </div>    
</div>

<!-- Create Class Modal -->
<div id="create_modal" tabindex="-1" aria-hidden="true" class="fixed top-0 left-0 right-0 z-50 hidden w-full p-4 overflow-x-hidden overflow-y-auto md:inset-0 h-[calc(100%-1rem)] max-h-full">
  <div class="relative w-full max-w-md max-h-full">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow">
          <button type="button" class="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-purple-950 rounded-lg text-sm w-8 h-8 ml-auto inline-flex justify-center items-center" data-modal-hide="create_modal">
              <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
                  <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
              </svg>
              <span class="sr-only">Close modal</span>
          </button>
          <div class="px-6 py-6 lg:px-8">
              <h3 class="mb-4 text-xl font-medium text-purple-950">Create Class Subject</h3>
              <form class="space-y-6">
                  <div>
                    <ClassesCombobox bind:Value={newClassSubClassId} />
                  </div>
                  <div>
                    <SubjectsCombobox bind:Value={newClassSubSubjectId} />
                  </div>
                  <div>
                    <label for="SA" class="block mb-2 text-sm font-medium text-purple-950">SA</label>
                    <input bind:checked={newClassSubSA} type="checkbox" name="SA" id="SA" class="w-4 h-4 text-purple-950 bg-gray-100 border-gray-300 rounded focus:ring-purple-950 focus:ring-2">      
                  </div>
                  <div>
                    <label for="EX" class="block mb-2 text-sm font-medium text-purple-950">Ex</label>
                    <input bind:checked={newClassSubEx} type="checkbox" name="EX" id="EX" class="w-4 h-4 text-purple-950 bg-gray-100 border-gray-300 rounded focus:ring-purple-950 focus:ring-2">      
                  </div>
                  <div>
                    <label for="KA" class="block mb-2 text-sm font-medium text-purple-950">KA</label>
                    <input bind:checked={newClassSubKA} type="checkbox" name="KA" id="KA" class="w-4 h-4 text-purple-950 bg-gray-100 border-gray-300 rounded focus:ring-purple-950 focus:ring-2">      
                  </div>
                  <button 
                      on:click={createClassSubject}
                      type="button" 
                      class="btn btn-primary" 
                      data-modal-hide="create_modal">
                      Create
                  </button>
              </form>
          </div>
      </div>
  </div>
</div> 
