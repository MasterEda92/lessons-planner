<script lang="ts">
    import { onMount } from "svelte";
    import { GetClassById, UpdateClass } from '../../../wailsjs/go/main/App';
    import Heading1 from "../components/headings/heading1.svelte";
    import { Grade } from "../enums/Grade";
    import { navigate } from "svelte-routing";
    import type { ClassListElement } from "../types/classes";

    export let id: string

    let _class: ClassListElement = {
        Id: id,
        Name: "",
        Grade: Grade.NONE,
        PupilCount: 0,
        PupilCountTimeIncreas: 0
    }

    function getClass () {
        GetClassById (id).then ((result) => {

        _class = {
            Id: result.Id,
            Name: result.Name,
            Grade: result.Grade,
            PupilCount: result.PupilCount,
            PupilCountTimeIncreas: result.PupilCountTI
        }
        }).catch ((error) => {
            alert (error);
        })
    }

    onMount (() => {
        getClass ();
    })

    function updateClass () {
        UpdateClass ({
            Id: id,
            PupilCount: _class.PupilCount,
            PupilCountTI: _class.PupilCountTimeIncreas
        }).then ((result) => {
            navigate("/classes", { replace: true });
        }).catch ((error) => {
            alert (error);
        })
    }
</script>


<Heading1>Edit Class</Heading1>

<form class=" py-3 space-y-6">
    <div>
        <label for="className" class="block mb-2 text-sm font-medium text-purple-950">Name</label>
        <input disabled bind:value={_class.Name} type="text" name="className" id="className" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-full p-2.5">
    </div>
    <div>
        <label for="Grade" class="block mb-2 text-sm font-medium text-purple-950">Name</label>
        <input disabled bind:value={_class.Grade} type="text" name="Grade" id="Grade" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-full p-2.5">
    </div>
    <div>
        <label for="pupilCount" class="block mb-2 text-sm font-medium text-purple-950">Pupil Count</label>
        <input bind:value={_class.PupilCount} type="number" name="pupilCount" id="pupilCount" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-full p-2.5" placeholder="Pupil Count" required>
    </div>
    <div>
        <label for="pupilCountTI" class="block mb-2 text-sm font-medium text-purple-950">Pupil Count with time increase</label>
        <input bind:value={_class.PupilCountTimeIncreas} type="number" name="pupilCountTI" id="pupilCountTI" class="bg-gray-50 border border-gray-300 text-purplr-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-full p-2.5" placeholder="Pupil Count with time increase" required>
    </div>
    <button 
        on:click={updateClass} 
        type="button" 
        class="btn btn-primary">
        Save
    </button>
</form>
