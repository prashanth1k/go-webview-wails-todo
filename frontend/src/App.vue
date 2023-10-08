<script setup>
import TodoList from './components/TodoList.vue'
import TodoNew from './components/TodoNew.vue'
import { onMounted, reactive } from 'vue';
import { ListTodos } from '../wailsjs/go/main/App'

const todos = reactive({ data: [] });

const getTodoList = async () => {
  todos.data = await ListTodos()
  console.log('todos.data.length: ', todos.data.length);
  console.log('todos.data: ', ...todos.data);
}

onMounted(() => {
  getTodoList()
  console.log(`Fetched todos..!`, todos.data)
})

</script>

<template>
  <div class="container">
    <nav class="nav">
      <div class="nav-left">
        <a class="brand" href="#">Wails Todos!</a>
      </div>
    </nav>

    <div class="card todocard">

      <div class="row">
        <div class="col-12 headersummary">{{ todos.data.length }} tasks | {{ todos.data.filter(el => el.status !=
          "complete"
        ).length }} open </div>
        <div class="col-12">
          <TodoNew :todos="todos" @add-todo="console.log('Received emit..', $event); todos.data.push($event)" />
        </div>
        <div class="col-12">
          <TodoList :todos="todos"
            @update-todo="console.log('Received emit..', $event); todos.data.splice(todos.data.findIndex(el => el.id === $event.id), 1, $event)"
            @delete-todo="console.log('Received emit..', $event); todos.data.splice(todos.data.findIndex(el => el.id === $event.id), 1)" />
        </div>

      </div>
    </div>
  </div>
</template>

<style></style>
