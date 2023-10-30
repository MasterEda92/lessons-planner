<script lang="ts">
    import Heading1 from "../components/headings/heading1.svelte";
    import { Hours } from "../consts/hours";
    import { Days } from "../consts/days";
    import TimeTableElement from "../components/TimeTableElement.svelte";
    
  </script>
  
  <Heading1>Timetable</Heading1>

  <!-- Class-Table -->
<div class="py-2">
  <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
      <table class="w-full text-sm text-left text-gray-500">
          <thead class="text-xs text-purple-950 uppercase bg-gray-200">
              <tr>
                  <th scope="col" class="px-6 py-3">
                      Hour
                  </th>
                  {#each Days as day (day.Id) }
                    <th scope="col" class="px-6 py-3">
                      {day.Name}
                    </th>  
                  {/each}
              </tr>
          </thead>
          <tbody>
              {#each Hours as _hour, i (_hour.Id)}
                  <tr class="bg-white hover:bg-gray-50">
                      <th scope="row" class="{(i%2==0) ? "bg-white" : "bg-gray-100"} px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                          {_hour.Name}
                      </th>
                      {#each Days as _day (_day.Id) }
                        <td class="{(i%2==0) ? "bg-white" : "bg-gray-100"} px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                          <TimeTableElement Day={_day.Day} Hour={_hour.Id} />
                        </td>
                      {/each}
                  </tr>
              {/each}
          </tbody>
      </table>
  </div>    
</div>