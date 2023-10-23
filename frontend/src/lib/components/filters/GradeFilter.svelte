<script lang="ts">
  import { Grade } from "../../enums/Grade";
  import {grades} from "../../consts/grades";
  import type { GradeElement } from "../../types/grades";

  let _grades: GradeElement[] = grades;

  export let Value: Grade = Grade.NONE
  export let selChange: (event: any) => void;

</script>

<label for="grades_filter" class="block mr-2 my-1 flex-nowrap mb-2 text-sm font-medium text-purple-950">Grade filter</label>
<select on:change={selChange} id="grades_filter" bind:value={Value} class="bg-gray-50 border border-gray-300 text-purple-950 text-sm rounded-lg focus:ring-purple-950 focus:border-purple-950 block w-auto mr-2 my-1">
{#each _grades as grade (grade.id)}
    <option value={grade.id}>{grade.name}</option>
{/each}
</select>
