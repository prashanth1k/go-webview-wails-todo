<template>
    <form @submit.prevent="addTodo">
        <div class="row">
            <div class="col-12 text-left">
                <h5>Add Task</h5>
            </div>
            <div class="col-7-md col-12">
                <label for="desc">Description</label>
                <input type="text" id="desc" v-model="todo.description" />
            </div>
            <div class="col-3-md col-12">
                <label for="due">Due (yyyy-mm-dd)</label>
                <input type="datetime" id="Due" v-model="todo.due" />
            </div>
            <div class="col-2-md col-12 cardaction">
                <a @click="addTodo" class="button primary">Add</a>
            </div>
        </div>
    </form>
</template>

<script setup>
import { reactive } from 'vue';
import { CreateTodo } from '../../wailsjs/go/main/App'

const emit = defineEmits(['addTodo'])
const todo = reactive({ description: "", created: "", due: "", status: "" });

const addTodo = async () => {
    console.log(`Emitting addTodo event..!`)
    todo.status = "In Progress"
    todo.created = new Date().toISOString().substring(0, 16).replace("T", " ")
    const todoResponse = await CreateTodo(todo)
    console.log('todoResponse: ', todoResponse);
    emit('addTodo', { ...todo })
    todo.description = ""
    todo.due = ""
}
</script>

  
    
  