<script lang="ts">
    import type { Class } from "../../types/classes";
    import { GetAllClasses } from '../../../../wailsjs/go/main/App';
    
    let classes: Class[] = [];
  
    export let Value: string = "0";
    
    function GetClasses () {
        GetAllClasses ().then ((result) => {
            classes = [];
            result.forEach(element => {
                classes.push ({Id: element.Id, Name: element.Name})
            });
            classes = [...classes];
        }).catch ((error) => {
            alert (error);
        })
    }

    GetClasses ();
  
  </script>
  
  <label for="classes_combobox" class="block mr-2 my-1 flex-nowrap mb-2 text-sm font-medium text-purple-950">Class</label>
  <select id="classes_combobox" bind:value={Value} class="bg-gray-50 border border-gray-300 text-purple-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-full mr-2 my-1">
  {#each classes as _class (_class.Id)}
      <option value={_class.Id}>{_class.Name}</option>
  {/each}
  </select>