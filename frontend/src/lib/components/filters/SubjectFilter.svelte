<script lang="ts">
    import type { Subject } from "../../types/subjects";
    import { GetAllSubjects } from '../../../../wailsjs/go/main/App';
    
    let subjects: Subject[] = [];
  
    export let Value: string = "0";
    export let selChange: (event: any) => void;

    function GetSubjects () {
        GetAllSubjects ().then ((result) => {
            subjects = [];
            subjects.push ({Id: "0", Name: "no filter"});
            result.forEach(element => {
                subjects.push ({Id: element.Id, Name: element.Name})
            });
            subjects = [...subjects];
        }).catch ((error) => {
            alert (error);
        })
    }

    GetSubjects ();
  
  </script>
  
  <label for="subjects_filter" class="block mr-2 my-1 flex-nowrap mb-2 text-sm font-medium text-purple-950">Subject filter</label>
  <select on:change={selChange} id="subjects_filter" bind:value={Value} class="bg-gray-50 border border-gray-300 text-purple-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-auto mr-2 my-1">
  {#each subjects as _subject (_subject.Id)}
      <option value={_subject.Id}>{_subject.Name}</option>
  {/each}
  </select>