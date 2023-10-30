<script lang="ts">
  import Heading1 from "../components/headings/heading1.svelte";
  import { onMount } from "svelte";
  import { GetSubjectById, UpdateSubject } from '../../../wailsjs/go/main/App';
  import { navigate } from "svelte-routing";
    
    export let id: string

    let subject: {
        Id: string,
        Name: string,
        Notes: string
    } = {Id: id, Name:"", Notes: ""}


    function getClass () {
        GetSubjectById (id).then ((result) => {

        subject = {
            Id: result.Id,
            Name: result.Name,
            Notes: result.Notes
        }
        }).catch ((error) => {
            alert (error);
        })
    }

    onMount (() => {
        getClass ();
    })

    function updateSubject () {
        UpdateSubject ({
            Id: id,
            Name: subject.Name,
            Notes: subject.Notes
        }).then ((result) => {
            navigate("/subjects", { replace: true });
        }).catch ((error) => {
            alert (error);
        })
    }
</script>

<Heading1>Edit Subject</Heading1>

<form class=" py-3 space-y-6">
  <div>
      <label for="SubjectName" class="block mb-2 text-sm font-medium text-purple-950">Subject</label>
      <input disabled bind:value={subject.Name} type="text" name="SubjectName" id="SubjectName" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-full p-2.5">
  </div>
  <div>
      <label for="Notes" class="block mb-2 text-sm font-medium text-purple-950">Notes</label>
      <input bind:value={subject.Notes} type="text" name="Notes" id="Notes" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-full p-2.5">
  </div>
  <button 
      on:click={updateSubject} 
      type="button" 
      class="btn btn-primary">
      Save
  </button>
</form>