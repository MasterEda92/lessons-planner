<script lang="ts">
  
  import {Link, Route, Router} from 'svelte-routing';
  import Home from './lib/Home.svelte';
  import Classes from './lib/pages/Classes.svelte';
  import ClassesEdit from './lib/pages/ClassesEdit.svelte';
  import Subjects from './lib/pages/Subjects.svelte';
  import SubjectsEdit from './lib/pages/SubjectsEdit.svelte';
  import ClassSubjects from './lib/pages/ClassSubjects.svelte';
  import Timetable from './lib/pages/Timetable.svelte';
  import TeachingUnits from './lib/pages/TeachingUnits.svelte';
  import ClassSubjectsEdit from './lib/pages/ClassSubjectsEdit.svelte';
</script>

<Router>
  <div class="flex flex-col min-h-screen py-3">
    <nav>
      <ul class="flex justify-center gap-3">
        <Link class="nav-link" to="/">Home</Link>
        <Link class="nav-link" to="/classes">Classes</Link>
        <Link class="nav-link" to="/subjects">Subjects</Link>
        <Link class="nav-link" to="/class_subjects">Class-Subjects</Link>
        <Link class="nav-link" to="/timetable">Timetable</Link>
        <Link class="nav-link" to="/teachingunits">Teaching-Units</Link>
      </ul>      
    </nav>
  
    <div class="text-center flex-1 p-5 border border-purple-950 rounded my-4 mx-auto w-4/5">
      <Route path="/" component={Home} />
      <Route path="/classes" component={Classes} />
      <Route path="/classes/:id" component={ClassesEdit} />
      <Route path="/subjects" component={Subjects} />
      <Route path="/subjects/:id" component={SubjectsEdit} />
      <Route path="/class_subjects" component={ClassSubjects} />
      <Route path="/class_subjects/:id" component={ClassSubjectsEdit} />
      <Route path="/timetable" component={Timetable} />
      <Route path="/teachingunits" component={TeachingUnits} />
    </div>

    <footer class="bg-purple-950/90 text-white py-4 text-center w-full">Lessons-Planner © 2023 Stefan Eder</footer>
  </div>
</Router>
