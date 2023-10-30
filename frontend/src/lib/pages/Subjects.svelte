<script lang="ts">
  import { navigate } from "svelte-routing";
  import Heading1 from "../components/headings/heading1.svelte";
  import { GetAllSubjects, CreateSubject, DeleteSubject } from '../../../wailsjs/go/main/App';
  import { onMount } from 'svelte';
  import { initFlowbite } from 'flowbite';

  let subjects: {
    Id: string,
    Name: string,
    Notes: string
  }[] = []

  // retriev all classes and add to classes array
  function getAllSubjects () {
    console.log ("outer getAllSubjects")
    GetAllSubjects().then ((result) => {
        let _subjects: { Id: string; Name: string; Notes: string; }[] = [];
        result.forEach ( (value) => {
            _subjects.push ({
              Id: value.Id,
              Name: value.Name,
              Notes: value.Notes
            })
        })
        subjects = [..._subjects];
        console.log ("getAllSubjects", subjects);
        }).catch ((error) => {
            alert (error);
        })
    }

    // delete class with given id and refresh table
    function deleteSubject (_id: string) {
        DeleteSubject({Id: _id}).then ((result) => {
            getAllSubjects ();

        }).catch((error) => {
            alert (error);
        })
    }

    // Create a new class and add to classes array
    function createSubject() {

        CreateSubject ({
            Name: newSubjectName,
            Notes: newSubjectNotes
        }).then ((result) => {
          getAllSubjects();
        }).catch ((error) => {
            alert (error);
        })

        newSubjectName = "";
        newSubjectNotes = "";
    }

    // component initialisation
    onMount (() => {
        initFlowbite();

        getAllSubjects ();
    });

    let newSubjectName: string = ""
    let newSubjectNotes: string = ""
</script>


<Heading1>Subjects</Heading1>

<div class="flex items-center justify-between">
  <div class="flex items-center">
      
  </div>
  <button 
      data-modal-target="create_modal" 
      data-modal-toggle="create_modal" 
      class="btn btn-primary"
      type="button">
      New
  </button>
</div>

<div class="py-2">
  <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
      <table class="w-full text-sm text-left text-gray-500">
          <thead class="text-xs text-purple-950 uppercase bg-gray-200">
              <tr>
                  <th scope="col" class="px-6 py-3">
                      Subject
                  </th>
                  <th scope="col" class="px-6 py-3">
                      Notes
                  </th>
                  <th scope="col" class="px-6 py-3">
                    <span class="sr-only">Edit</span>
                </th>
                <th scope="col" class="px-6 py-3">
                    <span class="sr-only">Delete</span>
              </tr>
          </thead>
          <tbody>
              {#each subjects as subject, i (subject.Id)}
                  <tr class="bg-white hover:bg-gray-50">
                      <th scope="row" class="{(i%2==0) ? "bg-white" : "bg-gray-100"} px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                          {subject.Name}
                      </th>
                      <td class="{(i%2==0) ? "bg-white" : "bg-gray-100"} px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                          {subject.Notes}
                      </td>
                      <td class="{(i%2==0) ? "bg-white" : "bg-gray-100"} px-6 py-4 text-right">
                          <button 
                              on:click={() => {
                                  const url = "/subjects/" + subject.Id;
                                  navigate(url, { replace: true });
                              }}>
                              <svg class="fill-purple-950" width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M13.94 5 19 10.06 9.062 20a2.25 2.25 0 0 1-.999.58l-5.116 1.395a.75.75 0 0 1-.92-.921l1.395-5.116a2.25 2.25 0 0 1 .58-.999L13.938 5Zm7.09-2.03a3.578 3.578 0 0 1 0 5.06l-.97.97L15 3.94l.97-.97a3.578 3.578 0 0 1 5.06 0Z"/></svg>
                          </button>
                      </td>
                      <td class="{(i%2==0) ? "bg-white" : "bg-gray-100"} px-6 py-4 text-right">
                          <button on:click={ () => deleteSubject(subject.Id)}>
                              <svg class="fill-purple-950" width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M21.5 6a1 1 0 0 1-.883.993L20.5 7h-.845l-1.231 12.52A2.75 2.75 0 0 1 15.687 22H8.313a2.75 2.75 0 0 1-2.737-2.48L4.345 7H3.5a1 1 0 0 1 0-2h5a3.5 3.5 0 1 1 7 0h5a1 1 0 0 1 1 1Zm-7.25 3.25a.75.75 0 0 0-.743.648L13.5 10v7l.007.102a.75.75 0 0 0 1.486 0L15 17v-7l-.007-.102a.75.75 0 0 0-.743-.648Zm-4.5 0a.75.75 0 0 0-.743.648L9 10v7l.007.102a.75.75 0 0 0 1.486 0L10.5 17v-7l-.007-.102a.75.75 0 0 0-.743-.648ZM12 3.5A1.5 1.5 0 0 0 10.5 5h3A1.5 1.5 0 0 0 12 3.5Z"/></svg>
                          </button>
                      </td>
                  </tr>
              {/each}
          </tbody>
      </table>
  </div>    
</div>

<!-- Create Subject Modal -->
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
              <h3 class="mb-4 text-xl font-medium text-purple-950">Create Class</h3>
              <form class="space-y-6">
                  <div>
                      <label for="subjectName" class="block mb-2 text-sm font-medium text-purple-950">Subject</label>
                      <input bind:value={newSubjectName} type="text" name="subjectName" id="subjectName" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-full p-2.5" placeholder="Subject Name" required>
                  </div>
                  <div>
                      <label for="Notes" class="block mb-2 text-sm font-medium text-purple-950">Notes</label>
                      <input bind:value={newSubjectNotes} type="text" name="Notes" id="Notes" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-full p-2.5" placeholder="Notes" required>
                  </div>
                  <button 
                      on:click={() => createSubject ()}
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
